package models

import ()

func (q QuerySet) buildQuery() string {
	sql := ""
	sql += q.processSelect()
	//sql += " FROM " + q.from
	sql += q.processFrom()
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
		for _, annotation := range q.annotated {
			sql += annotation + ", "
		}
	}

	sql = sql[0 : len(sql)-2]

	return sql
}

func (q QuerySet) processFrom() string {
	sql := " FROM "

	sql += q.model.dbTable

	// for _, field := range q.model.fieldList() {
	// 	if field.isRelation {
	// 		joinModel := field.relatedModel
	// 		joinTable := joinModel.dbTable
	// 		joinField := joinModel.getPrimaryKey().toSql()
	// 		sql += " JOIN " + joinTable + " ON " + field.toSql() + " = " + joinField
	// 	}
	// }

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
