package models

import (
  "strings"
)

type Lookup struct {
  fieldLookup string
  value interface{}
}

func lookupToSql(fieldLookup string, value interface{}) string {
  idx := strings.Index(fieldLookup, "__")

  if idx < 0 {
    return ""
  }

  field := fieldLookup[0:idx]
  lookup :=  fieldLookup[idx+2:]

  s := field

  switch lookup {
  case "exact":
    s += exact(value)

  case "iexact":
    s += iexact(value)

  case "contains":
    s += contains(value)

  case "icontains":
    s += icontains(value)






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
