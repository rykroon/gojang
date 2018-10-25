package gojang

import ()

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

	sql += ";"
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
	columns := "("
	values := "("

	for _, field := range q.set {
		columns += dbq(field.getDbColumn()) + ", "
		values += field.valueToSql() + ", "
	}

	columns = columns[:len(columns)-2] + ")"
	values = values[:len(values)-2] + ")"

	return columns + " VALUES " + values
}

func (q QuerySet) processUpdate() string {
	sql := " SET "

	for _, field := range q.set {
		sql += dbq(field.getDbColumn()) + " = " + field.valueToSql() + ", "
	}

	sql = sql[:len(sql)-2]
	return sql

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
