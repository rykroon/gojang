package gojang

import (
//"fmt"
)

func (q QuerySet) buildQuery() string {
	sql := ""
	if q.delete {
		sql += "DELETE "
	} else {
		sql += q.processSelect()
	}

	sql += q.processFrom()
	sql += q.processWhere()

	if !q.delete {
		sql += q.processOrderBy()
	}

	sql += ";"
	return sql
}

func (q QuerySet) processSelect() string {
	sql := "SELECT "

	if len(q.selected) == 0 {
		sql += "*"
	} else {

		for _, expr := range q.selected {
			sql += expr.asExpr() + ", "
		}

		sql = sql[:len(sql)-2]
	}

	return sql
}

func (q QuerySet) processFrom() string {
	sql := " FROM " + dbq(q.model.dbTable)

	if len(q.joins) > 0 {
		// add joins
	}

	return sql
}

func (q QuerySet) processWhere() string {
	sql := ""

	if len(q.lookups) != 0 {
		sql += " WHERE "

		for i, lookup := range q.lookups {
			if i == 0 {
				sql += lookup.asExpr()
			} else {
				sql += " AND " + lookup.asExpr()
			}
		}
	}

	return sql
}

func (q QuerySet) processOrderBy() string {
	sql := ""

	if len(q.orderBy) != 0 {
		sql += " ORDER BY "

		for _, sort := range q.orderBy {
			sql += sort.asExpr() + ", "
		}

		sql = sql[0 : len(sql)-2]
	}

	return sql
}
