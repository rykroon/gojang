package models

import (
  "fmt"
)

type Model struct {
  dbTable string
	fields map[string]Field
  objects Manager

  //Meta
  uniqueTogether []string
}

// type Field interface {
// 	CreateString() string
// }

type instance map[string]interface{}

func (m Model) Init(dbTable string) Model {
  m.dbTable = dbTable
  return m

}


func (m Model) CreateTable() string {
  //s := "CREATE TABLE " + DoubleQuotes(m.dbTable) + "("
  s := "CREATE TABLE " + doubleQuotes(m.dbTable) + "("

  for _,field := range m.fields {
    s += field.CreateString() + ", "
  }

  fmt.Println(len(m.fields))

  s += ")"

  return s
}
