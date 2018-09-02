package models

import ()

type Model struct {
	dbTable string
	Objects Manager

	//Meta
	//uniqueTogether []string
}

//type instance map[string]interface{}

func (m *Model) Init(dbTable string) {
	m.dbTable = dbTable
	m.Objects.model = m
}
