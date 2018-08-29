package models

import (
)

type Manager struct {

}

type Model struct {
  dbTable string
	fields []Field
  objects Manager
  uniqueTogether []string
}

type Field interface {
	CreateString() string
}

type instance map[string]interface{}

func (m Manager) All() {

}

func (m Model) CreateTable() {
  s := "CREATE TABLE " + doubleQuotes(m.dbTable) + "("

  for _,field := range m.fields {
    s += field.CreateString() + " "
  }

  s += ")"
}
