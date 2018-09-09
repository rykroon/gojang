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
	return f
}

//DbColumn Field Option - So the user can choose a seperate db name
func (f Field) DbColumn(name string) Field {
	f.dbColumn = name
	return f
}

//Default Values
func (f Field) Default(i interface{}) Field {
	t := reflect.TypeOf(i)
	k := t.Kind()

	if k == f.defaultType {
			f.defaultValue = valueToSql(i)
	} else {
		return f
	}
	return f
}
