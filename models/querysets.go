package models

import (
	"database/sql"
)

type QuerySet struct {
	//Db ...some sort of db connection
	db        sql.DB
	model     *Model
	Query     string
	evaluated bool

	distinct  bool
	selected  []string
	deferred  []string
	annotated []string

	from  string
	where []string
	//groupBy      string
	//having       string
	Ordered bool
	orderBy []sortExpression

	rows sql.Rows
}

type sortExpression struct {
	field       field
	orderOption string
}

func (f field) Asc() sortExpression {
	return sortExpression{field: f, orderOption: "ASC"}
}

func (f field) Desc() sortExpression {
	return sortExpression{field: f, orderOption: "DESC"}
}

//Functions that return QuerySets

func (q QuerySet) Filter(l lookup) QuerySet {
	q.where = append(q.where, l.toSql())
	q.Query = q.buildQuery()
	return q
}

func (q QuerySet) Exclude(l lookup) QuerySet {
	sql := "NOT(" + l.toSql() + ")"
	q.where = append(q.where, sql)
	q.Query = q.buildQuery()
	return q
}

func (q QuerySet) Annotate(a aggregate) QuerySet {
	q.annotated = append(q.annotated, a.asSql())
	q.Query = q.buildQuery()
	return q
}

func (q QuerySet) OrderBy(sortExpressions ...sortExpression) QuerySet {
	for _, sortExpression := range sortExpressions {
		q.orderBy = append(q.orderBy, sortExpression)
	}

	q.Query = q.buildQuery()
	return q
}

func (q QuerySet) Distinct(fields ...field) QuerySet {
	q.distinct = true

	for _, field := range fields {
		q.selected = append(q.selected, field.toSql())
	}

	q.Query = q.buildQuery()
	return q
}

func (q QuerySet) All() QuerySet {
	q.evaluated = false
	//maybe do other stuff?
	return q
}

//add fields to the deferred list
func (q QuerySet) Defer(fields ...field) QuerySet {
	for _, field := range fields {
		q.deferred = append(q.deferred, field.toSql())
	}

	q.Query = q.buildQuery()
	return q
}

//clear current array of select fields and deffered fields
func (q QuerySet) Only(fields ...field) QuerySet {
	q.selected = nil
	q.deferred = nil

	for _, field := range fields {
		q.selected = append(q.selected, field.toSql())
	}

	q.Query = q.buildQuery()
	return q
}

//Functions that do not return Querysets

func (q QuerySet) Get() Instance {
	//row := q.queryRow()
	return Instance{}
}

func (q QuerySet) Count() int {
	q.selected = nil
	q.deferred = nil
	q.annotated = nil

	q.selected = append(q.selected, "COUNT(*)")
	q.Query = q.buildQuery()

	var count int
	err := q.queryRow().Scan(&count)

	if err != nil {
		return count
	}

	return -1
}

func (q QuerySet) First() Instance {
	return Instance{}
}

func (q QuerySet) Last() Instance {
	return Instance{}
}

func (q QuerySet) Aggregate(a aggregate) Instance {
	q.selected = nil
	q.deferred = nil
	q.annotated = nil

	q.annotated = append(q.annotated, a.asSql())
	q.Query = q.buildQuery()
	return Instance{}
}

func (q QuerySet) Exists() bool {
	return false
}

func (q QuerySet) Delete() {

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
