package models

import ("fmt")

type Manager struct {
	model *Model
	//queryset QuerySet
}

func (m Manager) All() QuerySet {
	fmt.Println(m.model)
	qs := QuerySet{}
	qs.model = m.model
	qs.from = m.model.dbTable
	qs.select_ = "*"
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
