package gojang

import ()

type Manager struct {
	model *Model
	//queryset QuerySet
}

func newManager(model *Model) Manager {
	manager := Manager{model: model}
	return manager
}

func (m Manager) All() QuerySet {
	qs := newQuerySet(m.model)
	qs = qs.All()
	return qs
}

func (m Manager) Filter(lookups ...lookup) QuerySet {
	qs := newQuerySet(m.model)
	qs = qs.Filter(lookups...)
	return qs
}

func (m Manager) Exclude(lookups ...lookup) QuerySet {
	qs := newQuerySet(m.model)
	qs = qs.Exclude(lookups...)
	return qs
}

func (m Manager) OrderBy(orderBys ...sortExpression) QuerySet {
	qs := newQuerySet(m.model)
	qs = qs.OrderBy(orderBys...)
	return qs
}

// func (m Manager) Get(lookups ...lookup) {
// 	qs := newQuerySet(m.model)
// 	return qs.Get(lookups...)
// }
