package models

import ()

type Manager struct {
	model *Model
}

func (m Manager) All() QuerySet {
	qs := QuerySet{}
	qs.from = m.model.dbTable
	qs.select_ = "*"
	qs.Query = qs.buildQuery()
	return qs
}

func (m Manager) Filter(fieldLookup string, value interface{}) QuerySet {
	qs := m.All()
	qs = qs.Filter(fieldLookup, value)
	return qs
}

func (m Manager) Exclude(fieldLookup string, value interface{}) QuerySet {
	qs := m.All()
	qs = qs.Exclude(fieldLookup, value)
	return qs
}
