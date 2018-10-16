package gojang

import (
	//"fmt"
)

func (q QuerySet) buildQuery() string {
	sql := ""
	sql += q.processSelect()
	sql += q.processFrom()
	sql += q.processWhere()
	sql += q.processOrderBy()

	sql += ";"
	return sql
}

func (q QuerySet) processSelect() string {
	sql := "SELECT "

	if len(q.selected) == 0 {
		sql += "*"
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

		for i, filter := range q.lookups {
			if i == 0 {
				sql += filter.toSql()
			} else {
				sql += " AND " + filter.toSql()
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
			sql += expr.toSql() + ", "
		}

		sql = sql[0 : len(sql)-2]
	}

	return sql
}
