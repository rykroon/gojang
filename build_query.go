package gojang

import (
	"strings"
)

func (q QuerySet) buildQuery() string {
	sql := ""

	if q.insert {
		sql += "INSERT INTO "
	} else if q.update {
		sql += "UPDATE "
	} else if q.delete {
		sql += "DELETE "
	} else {
		sql += q.processSelect()
	}

	sql += q.processFrom()

	if q.insert {
		sql += q.processInsert()
	} else if q.update {
		sql += q.processUpdate()
	}

	if !q.insert {
		sql += q.processWhere()
	}

	if !(q.insert || q.update || q.delete) {
		sql += q.processOrderBy()
	}

	//sql += ";"
	return sql
}

func (q QuerySet) processSelect() string {
	sql := "SELECT "

	if len(q.selected) == 0 {
		sql += "*"
	} else {

		for _, expr := range q.selected {
			sql += expr.asSql() + ", "
		}

		sql = sql[:len(sql)-2]
	}

	return sql
}

func (q QuerySet) processFrom() string {
	sql := ""
	tableName := dbq(q.model.dbTable)

	if q.insert || q.update {
		sql = tableName
		return sql
	}

	sql = " FROM " + tableName

	if len(q.joins) > 0 {
		// add joins
	}

	return sql
}

func (q QuerySet) processInsert() string {
	var columnList []string
	var valueList []string

	for _, assign := range q.set {
		columnList = append(columnList, dbq(assign.lhs.getDbColumn()))
		//valueList = append(valueList, assign.lhs.valueAsSql())
		valueList = append(valueList, assign.rhs)
	}

	columns := "(" + strings.Join(columnList, ", ") + ")"
	values := "(" + strings.Join(valueList, ", ") + ")"

	return columns + " VALUES " + values
}

func (q QuerySet) processUpdate() string {
	var assignList []string
	for _, assign := range q.set {
		assignList = append(assignList, assign.asSql())
	}

	return " SET " + strings.Join(assignList, ", ")
}

func (q QuerySet) processWhere() string {
	sql := ""

	if len(q.lookups) != 0 {
		sql += " WHERE "

		for i, lookup := range q.lookups {
			if i == 0 {
				sql += lookup.asSql()
			} else {
				sql += " AND " + lookup.asSql()
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
			sql += sort.asSql() + ", "
		}

		sql = sql[0 : len(sql)-2]
	}

	return sql
}
