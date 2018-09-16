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

func (m Manager) Filter(l lookup) QuerySet {
	qs := m.All()
	qs = qs.Filter(l)
	return qs
}

func (m Manager) Exclude(l lookup) QuerySet {
	qs := m.All()
	qs = qs.Exclude(l)
	return qs
}

//Should the parameter be a lookup or do I create an aggregattion struct?
func (m Manager) Annotate() QuerySet {
	qs := m.All()
	return qs
}
