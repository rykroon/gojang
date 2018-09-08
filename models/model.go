package models

import ()

type Model struct {
	dbTable string
	Objects Manager
	Fields  map[string]Field

	//Meta
	//uniqueTogether []string
}

func NewModel(dbTable string) Model {
	m := Model{dbTable: dbTable}
	m.Objects.model = &m
	m.Fields = make(map[string]Field)
	return m
}

func (m *Model) AddField(fieldName string, f Field) {
	m.Fields[fieldName] = f
}

func (m Model) NewInstance() Instance {
	i := Instance{}
	i.model = &m
	return i
}
