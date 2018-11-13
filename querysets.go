package gojang

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	sq "github.com/masterminds/squirrel"
	"strings"
)

type QuerySet struct {
	model *Model
	Query string
	db    *sql.DB

	//distinct  bool
	selected []selecter //selectExpression
	//deferred
	insert bool
	update bool
	delete bool

	selectBuilder sq.SelectBuilder
	insertBuilder sq.InsertBuilder
	updateBuilder sq.UpdateBuilder
	deleteBuilder sq.DeleteBuilder

	set []assignment

	joins   []relatedField
	lookups []lookup

	Ordered bool
	orderBy []orderByExpression

	ResultCache []object
}

type selecter interface {
	sql.Scanner
	driver.Valuer
	Alias() string
	As(string)
	asSql() string
	getValue() interface{}
}

func newQuerySet(model *Model) QuerySet {
	qs := QuerySet{model: model, db: model.db}
	for _, field := range model.fields {
		qs.selected = append(qs.selected, field.copyField().(selecter))
	}

	qs.selectBuilder = qs.newSelect()
	qs.insertBuilder = qs.newInsert()
	qs.updateBuilder = qs.newUpdate()
	qs.deleteBuilder = qs.newDelete()

	qs.Query = qs.buildQuery()
	return qs
}

func (q *QuerySet) Evaluate() ([]object, error) {
	objects, err := q.queryAndScan()
	if err != nil {
		return nil, err
	}

	q.ResultCache = objects
	return q.ResultCache, nil
}

//
// Squirrel related methods
//

func (qs *QuerySet) newSelect() sq.SelectBuilder {
	return sq.Select().From(qs.model.dbTable).RunWith(qs.model.db)
}

func (qs *QuerySet) newInsert() sq.InsertBuilder {
	return sq.Insert(qs.model.dbTable).RunWith(qs.model.db)
}

func (qs *QuerySet) newUpdate() sq.UpdateBuilder {
	return sq.Update(qs.model.dbTable).RunWith(qs.model.db)
}

func (qs *QuerySet) newDelete() sq.DeleteBuilder {
	return sq.Delete(qs.model.dbTable).RunWith(qs.model.db)
}

//
//Functions that return QuerySets
//

//Returns a new QuerySet containing objects that match the given lookup parameters.
func (qs QuerySet) Filter(lookups ...lookup) QuerySet {
	for _, lookup := range lookups {
		qs.lookups = append(qs.lookups, lookup)
		qs.selectBuilder = qs.selectBuilder.Where(string(lookup))
		qs.updateBuilder = qs.updateBuilder.Where(string(lookup))
		qs.deleteBuilder = qs.deleteBuilder.Where(string(lookup))
	}

	qs.Query = qs.buildQuery()
	return qs
}

//returns a new QuerySet containing objects that do not match the given lookup parameters.
func (qs QuerySet) Exclude(lookups ...lookup) QuerySet {
	var lookupStrings []string
	for _, lookup := range lookups {
		lookupStrings = append(lookupStrings, string(lookup))
	}

	lookupList := strings.Join(lookupStrings, ",")
	notString := fmt.Sprintf("NOT(%v)", lookupList)
	exclude := lookup(notString)
	qs.lookups = append(qs.lookups, exclude)

	qs.selectBuilder = qs.selectBuilder.Where(notString)
	qs.updateBuilder = qs.updateBuilder.Where(notString)
	qs.deleteBuilder = qs.deleteBuilder.Where(notString)

	qs.Query = qs.buildQuery()
	return qs
}

func (qs QuerySet) OrderBy(orderBys ...orderByExpression) QuerySet {
	for _, orderBy := range orderBys {
		qs.orderBy = append(qs.orderBy, orderBy)
		qs.selectBuilder = qs.selectBuilder.OrderBy(string(orderBy))
	}

	qs.Query = qs.buildQuery()
	return qs
}

//Returns a QuerySet that returns an array of maps.
// func (q Queryset) Values() Queryset {
// 	return q
// }

func (q QuerySet) All() QuerySet {
	//maybe do other stuff?
	q.Query = q.buildQuery()
	return q
}

// func (q Queryset) SelectRelated() Queryset {
// 	return q
// }

//
//Functions that do not return Querysets
//

//populates the Model associated with the queryset with the data returned from the query
func (qs QuerySet) Get(lookups ...lookup) (object, error) {
	if len(lookups) > 0 {
		qs = qs.Filter(lookups...)
	}

	rows, err := qs.query()
	if err != nil {
		return nil, err
	}

	numOfRows := 0
	dest := qs.getDest()
	obj := newObj()

	for rows.Next() {
		numOfRows += 1

		if numOfRows > 2 {
			return nil, NewMultipleObjectsReturned()
		}

		err := rows.Scan(dest...)
		if err != nil {
			return nil, err
		}

		obj = qs.getObject()

	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if numOfRows == 0 {
		return nil, NewObjectDoesNotExist()
	}

	return obj, nil
}

func (q QuerySet) Create(assignments ...assignment) (object, error) {
	q.insert = true
	q.selected = nil
	q.selected = append(q.selected, q.model.Pk.copyField().(selecter))
	obj := newObj()

	for _, assign := range assignments {
		q.set = append(q.set, assign)
	}

	q.Query = q.buildQuery()
	pkeyName := q.model.Pk.ColumnName()
	q.Query += " RETURNING " + dbq(pkeyName)

	obj, err := q.queryRowAndScan()
	if err != nil {
		return nil, err
	}

	for _, assign := range assignments {
		attrName := q.model.colToAttr[assign.lhs.ColumnName()]
		obj.SetAttr(attrName, assign.lhs.getValue())
	}

	return obj, nil
}

//Returns an integer representing the number of objects in the database matching the QuerySet.
func (qs QuerySet) Count() (int, error) {
	sb := qs.selectBuilder.Columns("COUNT(*)")
	result := 0

	err := sb.Scan(&result)
	if err != nil {
		return 0, err
	}

	return result, nil
}

//Returns a map of aggregate values (averages, sums, etc.) calculated over
//the QuerySet. Each argument to aggregate() specifies a value that will
//be included in the map that is returned.
func (q *QuerySet) Aggregate(aggregates ...*aggregate) (object, error) {
	q.selected = nil
	q.orderBy = nil

	for _, expr := range aggregates {
		q.selected = append(q.selected, expr)
	}

	q.Query = q.buildQuery()
	obj, err := q.queryRowAndScan()

	if err != nil {
		return nil, err
	}

	return obj, nil
}

//Returns True if the QuerySet contains any results, and False if not.
func (qs QuerySet) Exists() (bool, error) {
	sb := qs.selectBuilder.Columns("*")
	rows, err := sb.Query()

	if err != nil {
		return false, err
	}

	if rows.Next() {
		return true, nil
	}

	return false, rows.Err()
}

func (q QuerySet) Update(assignments ...assignment) (int, error) {
	q.update = true

	for _, assign := range assignments {
		q.set = append(q.set, assign)
	}

	q.Query = q.buildQuery()
	result, err := q.exec()

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsAffected), nil
}

//Performs an SQL delete query on all rows in the QuerySet and returns the number of objects deleted
func (qs QuerySet) Delete() (int, error) {
	result, err := qs.deleteBuilder.Exec()
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsAffected), nil
}

//
//database/sql wrappers
//
func (q QuerySet) exec() (sql.Result, error) {
	return q.db.Exec(q.Query)
}

func (q QuerySet) query() (*sql.Rows, error) {
	return q.db.Query(q.Query)
}

func (q QuerySet) queryRow() *sql.Row {
	return q.db.QueryRow(q.Query)
}

//returns an interface{} slice with values that implement the sql.Scanner interface
func (q *QuerySet) getDest() []interface{} {
	var result []interface{}

	for _, dest := range q.selected {
		result = append(result, dest)
	}

	return result
}

//Returns an object of the values returned from the previous Scan()
func (q *QuerySet) getObject() object {
	obj := newObj()
	for _, expr := range q.selected {
		obj.SetAttr(expr.Alias(), expr.getValue())
	}

	return obj
}

func (q QuerySet) queryRowAndScan() (object, error) {
	row := q.queryRow()
	dest := q.getDest()
	err := row.Scan(dest...)

	if err != nil {
		return nil, err
	}

	return q.getObject(), nil
}

func (q QuerySet) queryAndScan() ([]object, error) {
	rows, err := q.query()

	if err != nil {
		return nil, err
	}

	var objects []object
	dest := q.getDest()

	for rows.Next() {
		err := rows.Scan(dest...)

		if err != nil {
			return nil, err
		}

		objects = append(objects, q.getObject())
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return objects, nil
}
