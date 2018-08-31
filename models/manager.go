package models

import ()

type Manager struct {
  model *model
	//queryset QuerySet
}

func (m Manager) All() QuerySet {
	return QuerySet{}
}

func (m Manager) Filter() QuerySet {
	return QuerySet{}
}

func (m Manager) Exclude() QuerySet {
	return QuerySet{}
}
