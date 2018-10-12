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

	insert bool
	update bool
	delete bool

	columns []string
	values  []string

	from  string
	lookups []lookup

	Ordered bool
	orderBy []sortExpression

	rows sql.Rows
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

// func (q QuerySet) Annotate(a aggregate) QuerySet {
// 	q.annotated = append(q.annotated, a.asSql())
// 	q.Query = q.buildQuery()
// 	return q
// }

func (q QuerySet) OrderBy(orderBys ...sortExpression) QuerySet {
	for _, orderBy := range orderBys {
		q.orderBy = append(q.orderBy, orderBy)
	}

	q.Query = q.buildQuery()
	return q
}

// func (q QuerySet) All() QuerySet {
// 	q.evaluated = false
// 	//maybe do other stuff?
// 	return q
// }




//Functions that do not return Querysets

// func (q QuerySet) Get() {
// 	row := q.queryRow()
// }

func (q QuerySet) Count() int {
	q.selected = nil
	//q.deferred = nil
	//q.annotated = nil

	q.selected = append(q.selected, "COUNT(*)")
	q.Query = q.buildQuery()

	var count int
	err := q.queryRow().Scan(&count)

	if err != nil {
		return count
	}

	return -1
}

// func (q QuerySet) Aggregate(a aggregate) modelInstance {
// 	q.selected = nil
// 	//q.deferred = nil
// 	//q.annotated = nil
//
// 	//q.annotated = append(q.annotated, a.asSql())
// 	q.selected = append(q.selected, a.asSql())
// 	q.Query = q.buildQuery()
//
// 	//q.queryRow()
// 	return modelInstance{}
// }

func (q QuerySet) Exists() bool {
	return false
}

func (q QuerySet) Update() (int64, error) {
	q.insert = true
	q.Query = q.buildQuery()
	result, err := q.exec()

	if err != nil {
		panic(err)
	}

	return result.RowsAffected()
}

func (q QuerySet) Delete() (int64, error) {
	q.delete = true
	q.Query = q.buildQuery()
	result, err := q.exec()

	if err != nil {
		panic(err)
	}

	return result.RowsAffected()
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
