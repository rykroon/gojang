package models

import (
	"reflect"
)

//Primary Key Option
func (f field) PrimaryKey(value bool) field {
	f.primaryKey = value

	if f.primaryKey {
		f.null = false
		f.unique = true
	}

	return f
}

//Unique Field Option
func (f field) Unique(unique bool) field {
	if (f.manyToMany || f.oneToOne) && unique {
		panic("Unique option is not valid for ManyToMany and OneToOne Fields")
	}

	if f.primaryKey && !unique {
		panic("Primary Key must be unique")
	}

	f.unique = unique
	return f
}

//Null Field Option
func (f field) Null(null bool) field {
	f.null = null

	if f.primaryKey && null {
		panic("Primary Key cannot be null")
	}

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
func (f field) DbColumn(name string) field {
	f.dbColumn = name
	//check if dbcolumn already exists in another field
	return f
}

//Default Values
func (f field) Default(i interface{}) field {
	t := reflect.TypeOf(i).String()

	if t == f.goType {
		f.defaultValue = interfaceToSql(i)
	} else {
		panic("Default value is not the same type as the field")
	}

	return f
}
