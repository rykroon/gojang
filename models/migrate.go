package models

import (
  "fmt"
  "reflect"
)

func Migrate(i interface{}) {

  v := reflect.ValueOf(i)
  if isAModel(v) {

    t := v.Type()
    numOfFields := v.NumField()
    sql := "CREATE TABLE table ("

    for idx := 0; idx < numOfFields; idx++ {
      structField := t.Field(idx)

      if structField.Type.String() == "models.Field" {
        columnName := t.Field(idx).Name
        sqlField := v.Field(idx).Interface().(Field).createString(columnName)
        sql += sqlField + ", "
      }
    }

    sql = sql[0:len(sql)-2] + ");"
    fmt.Println(sql)

  } else {
    fmt.Println("This is not a model")
  }
}


func isAModel(v reflect.Value) bool {

  if v.Kind().String() != "struct" {
    return false
  }

  t := v.Type()
  numOfFields := v.NumField()

  for idx := 0; idx < numOfFields; idx++ {
    fieldName := t.Field(idx).Name
    anonymous := t.Field(idx).Anonymous

    if (fieldName == "Model") && anonymous {
      return true
    }
  }

  return false
}
