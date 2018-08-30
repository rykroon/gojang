package models

import ()

type Manager struct {
  queryset Queryset
}

func (m *Manager) All() Queryset {
  m.queryset.query = ""
}

func (m *Manager) Filter() Queryset {

}

func (m *Manager) Exclude() Queryset {

}
