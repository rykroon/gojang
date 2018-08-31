package models

import ()

type Lookup struct {
  fieldLookup string
  value interface{}
}

func (l Lookup) toSql() string {
  field := ""
  lookup := ""

  //field := parse l.fieldLookup for the field
  //lookup := parse l .fieldLookup for the lookup

  s := field

  switch lookup {
  case "exact":
    s += exact(l.value)

  case "iexact":
    s += iexact(l.value)

  case "contains":
    s += contains(l.value)






  }

  return s
}


func exact(value interface{}) string {
  //if value is nil the cal isnull()
  return " = " + value.(string)
}

func iexact(value interface{}) string {
  return " ILIKE " + value.(string)
}

func contains(value interface{}) string {
  return " LIKE '%" + value.(string) + "%'"
}

func icontains(value interface{}) string {
  return " ILIKE '%" + value.(string) + "%'"
}
