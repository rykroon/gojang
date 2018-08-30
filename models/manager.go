package models

import ()

type Manager struct {
	queryset Queryset
}

func (m Manager) All() Queryset {
	return Queryset{}
}

func (m Manager) Filter() Queryset {
	return Queryset{}
}

func (m Manager) Exclude() Queryset {
	return Queryset{}
}
