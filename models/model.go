package models

import (
)

type model struct {
	dbTable string
	//Fields  map[string]Field
  Fields []Field
	Objects Manager

	//Meta
	uniqueTogether []string
}


type instance map[string]interface{}

func Model(dbTable string) model {
  m := model{dbTable: dbTable}
  m.Fields = make([]Field,7)
  return m
}

func (m model) Init(dbTable string) model {
	m.dbTable = dbTable
	return m

}

func (m model) CreateTable() string {
	//s := "CREATE TABLE " + DoubleQuotes(m.dbTable) + "("
	s := "CREATE TABLE " + doubleQuotes(m.dbTable) + " ("

	for _, field := range m.Fields {
		s += field.CreateString() + ", "
	}

	s += ")"

	return s
}
