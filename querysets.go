package gojang

import (
	"database/sql"
	//"reflect"
)

type QuerySet struct {
	model     *Model
	Query     string
	evaluated bool

	//distinct  bool
	selected []string
	//deferred  []string

	from  string
	lookups []lookup

	Ordered bool
	orderBy []sortExpression

	rows sql.Rows //cache
	//resultCache
}

type sortExpression struct {
	field field
	desc bool
}

func (s sortExpression) toSql() string {
	sql := s.field.sqlField()

	if s.desc {
		sql += " DESC"
	} else {
		sql += " ASC"
	}

	return sql
}

//Functions that return QuerySets

func (q QuerySet) Filter(lookups ...lookup) QuerySet {
	for _, lookup := range lookups {
		q.lookups = append(q.lookups, lookup)
	}

	q.Query = q.buildQuery()
	return q
}

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

func (q QuerySet) All() QuerySet {
	q.evaluated = false
	//maybe do other stuff?
	return q
}




//Functions that do not return Querysets

func (q QuerySet) Get() {
	//row := q.queryRow()
	return
}

func (q QuerySet) Count() int {
	return 0
}

func (q QuerySet) Exists() bool {
	return false
}

func (q QuerySet) Delete() int {
	return 0
}


//database/sql wrappers

// func (q QuerySet) Evaluate() []modelInstance {
// 	rows, err := q.query()
//
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
//
// 	columnTypes, err := rows.ColumnTypes()
// 	objects := make([]modelInstance, 0)
// 	dbColumnMap := q.model.dbColumnToAttrMap()
//
// 	for rows.Next() {
// 		pointers := make([]interface{}, len(columnTypes))
//
// 		for i, _ := range columnTypes {
// 			pointers[i] = new(interface{})
// 		}
//
// 		err := rows.Scan(pointers...)
//
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		obj := q.model.NewInstance()
//
// 		for idx, ptr := range pointers {
// 			attr := dbColumnMap[columnTypes[idx].Name()]
// 			val := reflect.ValueOf(ptr).Elem().Interface()
// 			obj.Set(attr, val)
// 		}
//
// 		objects = append(objects, obj)
// 	}
//
// 	return objects
// }

func (q QuerySet) exec() (sql.Result, error) {
	return q.model.db.Exec(q.Query)
}

func (q QuerySet) query() (*sql.Rows, error) {
	return q.model.db.Query(q.Query)
}

func (q QuerySet) queryRow() *sql.Row {
	return q.model.db.QueryRow(q.Query)
}
