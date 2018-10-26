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

func (m Manager) Create(fields ...field) error {
	qs := newQuerySet(m.model)
	return qs.Create(fields...)
}

func (m Manager) Count() (int, error) {
	qs := newQuerySet(m.model)
	return qs.Count()
}

//func (m Manager) Aggregate(aggregates ...aggregate) (map[string]interface{}, error) {
func (m Manager) Aggregate(aggregates ...function) (map[string]interface{}, error) {
	qs := newQuerySet(m.model)
	return qs.Aggregate(aggregates...)
}

func (m Manager) Exists() (bool, error) {
	qs := newQuerySet(m.model)
	return qs.Exists()
}

func (m Manager) Update(fields ...field) (int, error) {
	qs := newQuerySet(m.model)
	return qs.Update(fields...)
}

func (m Manager) Delete() (int, error) {
	qs := newQuerySet(m.model)
	return qs.Delete()
}
