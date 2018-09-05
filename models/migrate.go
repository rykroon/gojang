package models

import (
	"fmt"
	"reflect"
)



func Migrate(i interface{}) {

	v := reflect.ValueOf(i)

	if isAModel(v) {

		t := v.Type()
		fmt.Println(t.String()) //get the table name from t.String()

		hasPrimaryKey := false
		numOfFields := v.NumField()
		sql := "CREATE TABLE  ("
		sqlFields := ""

		for idx := 0; idx < numOfFields; idx++ {
			structField := t.Field(idx)

			if structField.Type.String() == "models.Field" {
				columnName := t.Field(idx).Name
				field := v.Field(idx).Interface().(Field)

				if field.foreignKey {
					columnName += "_id"
				}

				hasPrimaryKey = hasPrimaryKey || field.primaryKey
				sqlFields += field.createString(columnName) + ", "
			}
		}

		if !hasPrimaryKey {
			//idField := AutoField().PrimaryKey()
			//sqlFields = idField.createString("id") + ", " + sqlFields
		}

		sql += sqlFields[0:len(sqlFields)-2] + ");"
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
			//dbTable := v.Field(idx).Interface().(Model).dbTable
			return true
		}
	}

	return false
}
