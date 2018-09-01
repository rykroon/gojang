package models

import ()

type Manager struct {
  model *Model
	//queryset QuerySet
}

func (m Manager) All() QuerySet {
  q := QuerySet{}
  q.Query = "SELECT * FROM " + m.model.dbTable
	return q
}

func (m Manager) Filter() QuerySet {
	return QuerySet{}
}

func (m Manager) Exclude() QuerySet {
	return QuerySet{}
}
