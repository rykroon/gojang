package models

import ()

type QuerySet struct {
	//Db ...some sort of db connection

	model *Model
	Query string

	distinct bool
	select_  string
	//selected  []string
	//deferred  []string
	//annotated []string

	from  string
	where []string
	//groupBy      string
	//having       string
	Ordered bool
	orderBy []sortExpression
}

type sortExpression struct {
	field       Field
	orderOption string
}

func (f Field) Asc() sortExpression {
	return sortExpression{field: f, orderOption: "ASC"}
}

func (f Field) Desc() sortExpression {
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

// func (q QuerySet) Annotate(a aggregate) QuerySet {
// 	q.annotated = append(q.annotated, a.asSql())
// 	q.Query = q.buildQuery()
// 	return q
// }

func (q QuerySet) OrderBy(sortExpressions ...sortExpression) QuerySet {
	for _, sortExpression := range sortExpressions {
		q.orderBy = append(q.orderBy, sortExpression)
	}

	q.Query = q.buildQuery()
	return q
}

func (q QuerySet) Distinct() QuerySet {
	q.distinct = true
	q.Query = q.buildQuery()
	return q
}

// func (q QuerySet) Distinct(fields ...Field) QuerySet {
// 	q.distinct = true
//
// 	for _, field := range fields {
// 		q.selected = append(q.selected, field.toSql())
// 	}
//
// 	q.Query = q.buildQuery()
// 	return q
// }

//add fields to the deferred list
// func (q QuerySet) Defer(fields ...Field) QuerySet {
// 	for _, field := range fields {
// 		q.deferred = append(q.deferred, field.toSql())
// 	}
//
// 	q.Query = q.buildQuery()
// 	return q
// }
//
// //clear current array of select fields and deffered fields
// func (q QuerySet) Only(fields ...Field) QuerySet {
// 	q.selected = nil
// 	q.deferred = nil
//
// 	for _, field := range fields {
// 		q.selected = append(q.selected, field.toSql())
// 	}
//
// 	q.Query = q.buildQuery()
// 	return q
// }

//Functions that do not return Querysets

func (q QuerySet) Get() Instance {
	return Instance{}
}

func (q QuerySet) Count() int {
	return 0
}

func (q QuerySet) First() Instance {
	return Instance{}
}

func (q QuerySet) Last() Instance {
	return Instance{}
}

func (q QuerySet) Aggregate(a aggregate) Instance {
	return Instance{}
}

func (q QuerySet) Exists() bool {
	return false
}

func (q QuerySet) Delete() {

}
