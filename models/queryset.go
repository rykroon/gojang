package models

import ()

type QuerySet struct {
	model *Model
	Query string

	select_      string
	insrtUpdtDlt string
	from         string
	where        []string
	//groupBy      string
	//having       string
	orderBy []string
}

func (q QuerySet) buildQuery() string {
	s := "SELECT " + q.select_
	s += " FROM " + q.from

	if len(q.where) != 0 {
		for i, filter := range q.where {
			if i == 0 {
				s += " WHERE " + filter
			} else {
				s += " AND " + filter
			}
		}
	}

	if len(q.orderBy) != 0 {
		s += " ORDER BY "

		for idx := 0; idx < len(q.orderBy); idx++ {
			field := q.orderBy[idx]

			if string(field[0]) == "-" {
				field = field[1:]
				s += field + " DESC, "
			} else {
				s += field + ", "
			}
		}
		s = s[0 : len(s)-2]
	}

	s += ";"
	return s
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
	for idx := range fields {
		q.orderBy = append(q.orderBy, fields[idx])
	}

	q.Query = q.buildQuery()
	return q
}

func (q QuerySet) Distinct(fields ...string) QuerySet {
	return q
}

//Functions that do not return Querysets

//Lookup can be empty. Also takes into account previous filters/excludes/etc
func (q QuerySet) Get() Instance {
	return Instance{}
}

func (q QuerySet) Count() int {
	return 0
}

func (q QuerySet) Exists() bool {
	return false
}

func (q QuerySet) Delete() {

}
