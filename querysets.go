package gojang

import (
	"database/sql"
	//"fmt"
	//"reflect"
)

type QuerySet struct {
	model *Model
	Query string
	db    *sql.DB

	//distinct  bool
	selected []selectExpression
	//deferred
	insert bool
	update bool
	delete bool

	set []field

	joins   []relatedField
	lookups []lookup

	Ordered bool
	orderBy []sortExpression

	//rows sql.Rows //cache
	//rows []row

	//resultCache
}

func newQuerySet(model *Model) QuerySet {
	q := QuerySet{model: model, db: model.db}
	for _, field := range model.fields {
		q.selected = append(q.selected, field.copy())
	}

	q.Query = q.buildQuery()
	return q
}

//Functions that return QuerySets

//Returns a new QuerySet containing objects that match the given lookup parameters.
func (q QuerySet) Filter(lookups ...lookup) QuerySet {
	for _, lookup := range lookups {
		q.lookups = append(q.lookups, lookup)
	}

	q.Query = q.buildQuery()
	return q
}

//returns a new QuerySet containing objects that do not match the given lookup parameters.
func (q QuerySet) Exclude(lookups ...lookup) QuerySet {
	for _, lookup := range lookups {
		lookup.not = true
		q.lookups = append(q.lookups, lookup)
	}

	q.Query = q.buildQuery()
	return q
}

func (q QuerySet) OrderBy(orderBys ...sortExpression) QuerySet {
	for _, orderBy := range orderBys {
		q.orderBy = append(q.orderBy, orderBy)
	}

	q.Query = q.buildQuery()
	return q
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

//Functions that do not return Querysets

//populates the Model associated with the queryset with the data returned from the query
func (q QuerySet) Get(lookups ...lookup) ([]interface{}, error) {
	if len(lookups) > 0 {
		q = q.Filter(lookups...)
	}

	rows, err := q.query()
	if err != nil {
		return nil, err
	}

	numOfRows := 0
	dest := q.getDest()

	var result []interface{}

	for rows.Next() {
		numOfRows += 1

		if numOfRows > 2 {
			return nil, NewMultipleObjectsReturned()
		}

		err := rows.Scan(dest...)
		if err != nil {
			return nil, err
		}

		result = q.getScannedValues()
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if numOfRows == 0 {
		return nil, NewObjectDoesNotExist()
	}

	return result, nil
}

func (q QuerySet) Create(fields ...field) error {
	q.insert = true

	for _, field := range fields {
		q.set = append(q.set, field)
	}

	q.Query = q.buildQuery()
	_, err := q.exec()

	if err != nil {
		return err
	}

	return nil
}

//Returns an integer representing the number of objects in the database matching the QuerySet.
func (q QuerySet) Count() (int, error) {
	q.selected = nil
	q.orderBy = nil

	var star star
	countExpr := star.Count()
	q.selected = append(q.selected, countExpr)
	q.Query = q.buildQuery()

	err := q.queryRowAndScan()
	if err != nil {
		return 0, err
	}

	result := int(countExpr.getValue().(int32))
	return result, nil
}

//Returns a map of aggregate values (averages, sums, etc.) calculated over
//the QuerySet. Each argument to aggregate() specifies a value that will
//be included in the map that is returned.
func (q *QuerySet) Aggregate(aggregates ...aggregate) (map[string]interface{}, error) {
	q.selected = nil
	q.orderBy = nil
	result := make(map[string]interface{})

	for _, expr := range aggregates {
		q.selected = append(q.selected, expr)
	}

	q.Query = q.buildQuery()
	err := q.queryRowAndScan()

	if err != nil {
		return nil, err
	}

	for _, agg := range aggregates {
		result[agg.outputField.getDbColumn()] = agg.getValue()
	}

	return result, nil
}

//Returns True if the QuerySet contains any results, and False if not.
func (q QuerySet) Exists() (bool, error) {
	rows, err := q.query()
	if err != nil {
		return false, err
	}

	if rows.Next() {
		return true, nil
	}

	return false, rows.Err()
}

func (q QuerySet) Update(fields ...field) (int, error) {
	q.update = true

	for _, field := range fields {
		q.set = append(q.set, field)
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
func (q QuerySet) Delete() (int, error) {
	q.delete = true
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

//database/sql wrappers an auxillary functions

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

//return the values from the previous call to scan
func (q *QuerySet) getScannedValues() []interface{} {
	var result []interface{}

	for _, expr := range q.selected {
		result = append(result, expr.getValue())
	}

	return result
}

func (q QuerySet) queryRowAndScan() error {
	row := q.queryRow()
	dest := q.getDest()
	err := row.Scan(dest...)
	return err
}

func (q QuerySet) queryAndScan() ([]interface{}, error) {
	rows, err := q.query()
	if err != nil {
		return nil, err
	}

	dest := q.getDest()

	var result []interface{}

	for rows.Next() {
		err := rows.Scan(dest...)

		if err != nil {
			return nil, err
		}

		values := q.getScannedValues()
		result = append(result, values)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}
