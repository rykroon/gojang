package models

import (
)

type QuerySet struct {
	//Db ...some sort of db connection

	model *Model
	Query string

	distinct bool
	selected []string
	deferred []string
	annotated []string //add some logic

	from     string
	where    []string
	//groupBy      string
	//having       string
	Ordered bool
	orderBy []sortExpression
}

type sortExpression struct {
	field Field
	orderOption string
}


func (q QuerySet) buildQuery() string {
	sql := ""
	sql += q.processSelect()
	sql += " FROM " + q.from
	sql += q.processWhere()
	sql += q.processOrderBy()

	sql += ";"
	return sql
}


func (q QuerySet) processSelect() string {
	sql := "SELECT "

	if q.distinct {
		sql += "DISTINCT "
	}

	if len(q.selected) == 0 && len(q.deferred) == 0 && len(q.annotated) == 0 {
		sql += "*"
		return sql
	}

	selected := []string{}

	if len(q.selected) != 0 {
		selected = q.selected
	} else {
		selected = q.model.sqlFieldList()
	}

	for _, field := range selected {
		foundDefer := false

		if len(q.deferred) != 0 {
			for _, deferredField := range q.deferred {
				if field == deferredField {
					foundDefer = true
					break
				}
			}

			if foundDefer {
				continue
			}
		}

		sql += field + ", "
	}

	if len(q.annotated) != 0 {
		for _,annotation := range q.annotated {
			sql += annotation + ", "
		}
	}

	sql = sql[0 : len(sql)-2]
	return sql
}



func (q QuerySet) processWhere() string {
	sql := ""

	if len(q.where) != 0 {
		sql += " WHERE "

		for i, filter := range q.where {
			if i == 0 {
				sql += filter
			} else {
				sql += " AND " + filter
			}
		}
	}

	return sql
}



func (q QuerySet) processOrderBy() string {
	sql := ""

	if len(q.orderBy) != 0 {
		sql += " ORDER BY "

		for _, expr := range q.orderBy {
			sql += expr.field.toSql() + " " + expr.orderOption + ", "
		}

		sql = sql[0 : len(sql)-2]
	}

	return sql
}


func (f Field) Asc() sortExpression {
	return sortExpression{field: f, orderOption: "ASC"}
}

func (f Field) Desc() sortExpression {
	return sortExpression{field: f, orderOption: "DESC"}
}


//Functions that return QuerySets

func (q QuerySet) Filter(l lookup) QuerySet {
	q.where = append(q.where, l.asSql())
	q.Query = q.buildQuery()
	return q
}

func (q QuerySet) Exclude(l lookup) QuerySet {
	sql := "NOT(" + l.asSql() + ")"
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


func (q QuerySet) Distinct(fields ...Field) QuerySet {
	q.distinct = true

	for _, field := range fields {
		q.selected = append(q.selected, field.toSql())
	}

	q.Query = q.buildQuery()
	return q
}


//add fields to the deferred list
func (q QuerySet) Defer(fields ...Field) QuerySet {
	for _, field := range fields {
		q.deferred = append(q.deferred, field.toSql())
	}

	q.Query = q.buildQuery()
	return q
}

//clear current array of select fields and deffered fields
func (q QuerySet) Only(fields ...Field) QuerySet {
	q.selected = nil
	q.deferred = nil

	for _, field := range fields {
		q.selected = append(q.selected, field.toSql())
	}

	q.Query = q.buildQuery()
	return q
}




//Functions that do not return Querysets

//Lookup can be empty. Also takes into account previous filters/excludes/etc
func (q QuerySet) Get() Instance {
	return Instance{}
}

func (q QuerySet) Exists() bool {
	return false
}

func (q QuerySet) Delete() {

}




//aggregate functions

//maybe create a seperate aggregate struct
func (q QuerySet) Count() int {
	q.selected = nil
	q.deferred = nil
	//q.selected = append(q.selected, "COUNT(*)")
	q.Query = q.buildQuery()
	//execute
	return 0
}
