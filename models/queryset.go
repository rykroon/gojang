package models

import ()

type QuerySet struct {
  model *model
	query string

	//select_ string
	//from string
	//where []string
	//groupBy string
	//having string
	//orderBy string
}


//Functions that return QuerySets

func (q QuerySet) Filter(l Lookup) QuerySet {
	return QuerySet{}
}

func (q QuerySet) Exclude(l Lookup) QuerySet {
	return QuerySet{}
}

func (q QuerySet) OrderBy(f []string) QuerySet {
	return QuerySet{}
}


//Functions that do not return Queryset
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
