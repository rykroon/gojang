package models

import ()

type QuerySet struct {
  model *model
	query string

	//select_ string
	//from string
	where []string
	//groupBy string
	//having string
	//orderBy string
}

func (q QuerySet) buildQuery() string {
  return ""
}



//Functions that return QuerySets

func (q QuerySet) Filter(l Lookup) QuerySet {
  q.where = append(q.where, l.toSql())
  q.query = q.buildQuery()
	return q
}

func (q QuerySet) Exclude(l Lookup) QuerySet {
  sql := " NOT(" + l.toSql() + ")"
  q.where = append(q.where, sql)
  q.query = q.buildQuery()
	return q
}

func (q QuerySet) OrderBy(f []string) QuerySet {
	return QuerySet{}
}



//Functions that do not return Querysets
//Lookup can be empty. Also takes into account previous filters/excludes/etc
func (q QuerySet) Get(l Lookup) Instance {
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
