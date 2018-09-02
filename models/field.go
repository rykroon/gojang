package models

import (
	"strconv"
)

type Field struct {
	dbColumn   string
	dbDataType string

	//specific for CharField
	maxLength int

	//specfic for DecimalField
	maxDigits     int
	decimalPlaces int

	//constraints
	null       bool
	primaryKey bool
	unique     bool
}

//maybe?
// type FldOptn struct {
// 	dbColumn   string
// 	null       bool
// 	primaryKey bool
// 	unique     bool
// }

func AutoField() Field {
	return Field{dbDataType: "SERIAL"}
}

func BooleanField() Field {
	return Field{dbDataType: "boolean"}
}

func CharField(maxLength int) Field {
	n := strconv.Itoa(maxLength)
	dataType := "varchar(" + n + ")"

	return Field{dbDataType: dataType, maxLength: maxLength}
}

func DecimalField(maxDigits int, decimalPlaces int) Field {
	precision := strconv.Itoa(maxDigits)
	scale := strconv.Itoa(decimalPlaces)
	dataType := "NUMERIC(" + precision + ", " + scale + ")"

	return Field{dbDataType: dataType, maxDigits: maxDigits, decimalPlaces: decimalPlaces}
}

func FloatField() Field {
	return Field{dbDataType: "double precision"}
}

func IntegerField() Field {
	return Field{dbDataType: "integer"}
}

func TextField() Field {
	return Field{dbDataType: "text"}
}

//Constraints
func (f Field) PrimaryKey() Field {
	f.primaryKey = true
	return f
}

func (f Field) IsNull() Field {
	f.null = true
	return f
}

func (f Field) Unique() Field {
	f.unique = true
	return f
}

func (f Field) createString(dbColumn string) string {

	s := dbColumn + " " + f.dbDataType

	if f.primaryKey {
		s += " PRIMARY KEY"
	} else {

		if f.null {
			s += " NULL"
		} else {
			s += " NOT NULL"
		}

		if f.unique {
			s += " UNIQUE"
		}
	}

	return s
}
