package models

import ()

type Manager struct {
	model *Model
	//queryset QuerySet
}

func (m Manager) All() QuerySet {
	qs := QuerySet{}
	qs.model = m.model
	qs.from = m.model.dbTable
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

// func (m Manager) Filter(l lookup) QuerySet {
// 	qs := m.All()
// 	qs = qs.Filter(l)
// 	return qs
// }
//
// func (m Manager) Exclude(l lookup) QuerySet {
// 	qs := m.All()
// 	qs = qs.Exclude(l)
// 	return qs
// }
