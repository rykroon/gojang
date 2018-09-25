package models

import (
	"database/sql"
	"fmt"
)

type QuerySet struct {
	model     *Model
	Query     string
	evaluated bool

	distinct  bool
	selected  []string
	deferred  []string
	annotated []string

	insert bool
	update bool
	delete bool

	columns []string
	values []string

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
		if field.primaryKey {
			panic("Cannot defer the primary key")
		}

		q.deferred = append(q.deferred, field.toSql())
	}

	q.Query = q.buildQuery()
	return q
}

//clear current array of select fields and deffered fields
func (q QuerySet) Only(fields ...field) QuerySet {
	q.selected = nil
	q.deferred = nil

	foundPrimaryKey := false

	for _, field := range fields {
		if field.primaryKey {
			foundPrimaryKey = true
		}

		q.selected = append(q.selected, field.toSql())
	}

	if !foundPrimaryKey {
		q.selected = append(q.selected, q.model.getPrimaryKey().toSql())
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

	//q.queryRow()
	return Instance{}
}

func (q QuerySet) Exists() bool {
	return false
}

func (q QuerySet) Update() {
	q.insert = true
	q.Query = q.buildQuery()
	//q.exec()
}

func (q QuerySet) Delete() {
	q.delete = true
	q.Query = q.buildQuery()
	//q.exec()
}

//database/sql wrappers

func (q QuerySet) Evaluate() {
	rows, err := q.query()

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	cols, err := rows.Columns()


	for rows.Next() {
		vals := make([]interface{}, len(cols))

		for i, _ := range cols {
			fmt.Println(cols[i])
			//get type from q.model.Field() // need reverse lookup on dbcolumn
			vals[i] = new(sql.RawBytes)
		}

		err := rows.Scan(vals...)

		if err != nil {
			panic(err)
		}
		
		fmt.Println(vals)
	}


}

func (q QuerySet) exec() (sql.Result, error) {
	return q.model.db.Exec(q.Query)
}

func (q QuerySet) query() (*sql.Rows, error) {
	return q.model.db.Query(q.Query)
}

func (q QuerySet) queryRow() *sql.Row {
	return q.model.db.QueryRow(q.Query)
}
