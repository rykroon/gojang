package models

import ()

type QuerySet struct {
	model *Model
	Query string

	selected []string
	deferred []string
	distinct bool
	from     string
	where    []string
	//groupBy      string
	//having       string
	orderBy []string
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
		sql += " DISTINCT "
	}

	if len(q.selected) == 0 && len(q.deferred) == 0 {
		return "*"
	}

	selected := []string{}

	if len(q.selected) != 0 {
		selected = q.selected
	} else {
		selected = q.model.fieldList()
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

		for _, field := range q.orderBy {
			if string(field[0]) == "-" {
				field = field[1:]
				sql += field + " DESC, "
			} else {
				sql += field + ", "
			}
		}

		sql = sql[0 : len(sql)-2]
	}

	return sql
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

func (q QuerySet) OrderBy(fields ...string) QuerySet {
	for _, field := range fields {
		q.orderBy = append(q.orderBy, field)
	}

	q.Query = q.buildQuery()
	return q
}

func (q QuerySet) Distinct(fields ...string) QuerySet {
	q.distinct = true

	for _, field := range fields {
		q.selected = append(q.selected, field)
	}

	q.Query = q.buildQuery()
	return q
}

//add fields to the deferred list
func (q QuerySet) Defer(fields ...string) QuerySet {
	for _, field := range fields {
		q.deferred = append(q.deferred, field)
	}

	q.Query = q.buildQuery()
	return q
}

//clear current array of select fields and deffered fields
func (q QuerySet) Only(fields ...string) QuerySet {
	q.selected = nil
	q.deferred = nil

	for _, field := range fields {
		q.selected = append(q.selected, field)
	}

	q.Query = q.buildQuery()
	return q
}

//Functions that do not return Querysets

//Lookup can be empty. Also takes into account previous filters/excludes/etc
func (q QuerySet) Get() Instance {
	return Instance{}
}

func (q QuerySet) Count() int {
	q.selected = nil
	q.deferred = nil
	q.selected = append(q.selected, "COUNT(*)")
	q.Query = q.buildQuery()
	//execute
	return 0
}

func (q QuerySet) Exists() bool {
	return false
}

func (q QuerySet) Delete() {

}
