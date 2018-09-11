package models

import (
	"reflect"
)

//Primary Key Option
func (f Field) PrimaryKey(value bool) Field {
	f.primaryKey = value
	return f
}

//Unique Field Option
func (f Field) Unique(value bool) Field {
	f.unique = value
	return f
}

//Null Field Option
func (f Field) Null(value bool) Field {
	f.null = value

	if f.null {
		switch f.goType {
		case "bool":
			f.goType = "sql.NullBool"

		case "float64":
			f.goType = "sql.NullFloat64"

		case "int32":
			f.goType = "sql.NullInt64"

		case "string":
			f.goType = "sql.NullString"
		}
	}

	return f
}

//DbColumn Field Option - So the user can choose a seperate db name
func (f Field) DbColumn(name string) Field {
	f.dbColumn = name
	return f
}

//Default Values
func (f Field) Default(i interface{}) Field {
	t := reflect.TypeOf(i).String()

	if t == f.goType {
		f.defaultValue = interfaceToSql(i)
	}

	return f
}
