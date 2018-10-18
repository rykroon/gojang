package gojang

import (
	"database/sql"
	//"errors"
)

type QuerySet struct {
	model *Model
	Query string
	db    *sql.DB

	//distinct  bool
	selected []string
	//deferred
	delete bool

	joins   []relatedField
	lookups []lookup

	Ordered bool
	orderBy []sortExpression

	rows sql.Rows //cache
	//resultCache
}

type sortExpression struct {
	field field
	desc  bool
}

func (s sortExpression) toSql() string {
	sql := s.field.toSql()

	if s.desc {
		sql += " DESC"
	} else {
		sql += " ASC"
	}

	return sql
}

func newQuerySet(model *Model) QuerySet {
	q := QuerySet{model: model, db: model.db}
	for _, field := range model.fields {
		q.selected = append(q.selected, field.toSql())
	}

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
func (q QuerySet) Get(lookups ...lookup) error {
	if len(lookups) > 0 {
		q = q.Filter(lookups...)
	}

	rows, err := q.query()
	if err != nil {
		return err
	}

	numOfRows := 0

	for rows.Next() {
		numOfRows += 1

		if numOfRows > 2 {
			return NewMultipleObjectsReturned()
		}

		err := q.model.setFromRows(rows)
		if err != nil {
			return err
		}
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	if numOfRows == 0 {
		return NewObjectDoesNotExist()
	}

	return nil
}

//Returns an integer representing the number of objects in the database matching the QuerySet.
func (q QuerySet) Count() (int, error) {
	q.selected = nil
	q.selected = append(q.selected, "COUNT(*)")
	q.Query = q.buildQuery()
	row := q.queryRow()
	result := 0

	err := row.Scan(&result)
	if err != nil {
		return 0, err
	}

	return result, nil
}

//Returns True if the QuerySet contains any results, and False if not.
func (q QuerySet) Exists() (bool, error) {
	count, err := q.Count()
	if err != nil {
		return false, err
	}

	result := count != 0
	return result, nil
}

// func (q QuerySet) Update() int {
// 	return 0
// }

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

//database/sql wrappers
func (q QuerySet) exec() (sql.Result, error) {
	return q.db.Exec(q.Query)
}

func (q QuerySet) query() (*sql.Rows, error) {
	return q.db.Query(q.Query)
}

func (q QuerySet) queryRow() *sql.Row {
	return q.db.QueryRow(q.Query)
}
