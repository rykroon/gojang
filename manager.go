package gojang

import ()

type Manager struct {
	model *Model
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

func (m Manager) Get(lookups ...lookup) (interface{}, error) {
	qs := newQuerySet(m.model)
	return qs.Get(lookups...)
}

func (m Manager) Create(assignments ...assignment) error {
	qs := newQuerySet(m.model)
	return qs.Create(assignments...)
}

func (m Manager) Count() (int, error) {
	qs := newQuerySet(m.model)
	return qs.Count()
}

func (m Manager) Aggregate(aggregates ...aggregate) (map[string]interface{}, error) {
	qs := newQuerySet(m.model)
	return qs.Aggregate(aggregates...)
}

func (m Manager) Exists() (bool, error) {
	qs := newQuerySet(m.model)
	return qs.Exists()
}

func (m Manager) Update(assignments ...assignment) (int, error) {
	qs := newQuerySet(m.model)
	return qs.Update(assignments...)
}

func (m Manager) Delete() (int, error) {
	qs := newQuerySet(m.model)
	return qs.Delete()
}
