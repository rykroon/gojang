package models

import (
	"strconv"
	//"reflect"
)

type Field struct {
	dbColumn   string
	dbType string
	goType string

	//specific for CharField
	maxLength int

	//specfic for DecimalField
	maxDigits     int
	decimalPlaces int

	//constraints
	null       bool
	primaryKey bool
	unique     bool
	//defaultType reflect.Kind
	defaultValue string

	foreignKey bool
	to         *Model
	onDelete   string
}

func AutoField() Field {
	return Field{dbType: "SERIAL", goType: "int32"}
}

func BooleanField() Field {
	return Field{dbType: "BOOLEAN", goType: "bool"}
}

func CharField(maxLength int) Field {
	n := strconv.Itoa(maxLength)
	dataType := "VARCHAR(" + n + ")"

	return Field{dbType: dataType, maxLength: maxLength, goType: "string"}
}

func DecimalField(maxDigits int, decimalPlaces int) Field {
	precision := strconv.Itoa(maxDigits)
	scale := strconv.Itoa(decimalPlaces)
	dataType := "NUMERIC(" + precision + ", " + scale + ")"

	return Field{dbType: dataType, maxDigits: maxDigits, decimalPlaces: decimalPlaces, goType: "float64"}
}

func FloatField() Field {
	return Field{dbType: "DOUBLE PRECISION", goType:"float64"}
}

func ForeignKey(m *Model, onDelete string) Field {
	return Field{dbType: "INTEGER", foreignKey: true, to: m, onDelete: onDelete, goType:"int32"}
}

func IntegerField() Field {
	return Field{dbType: "INTEGER", goType:"int32"}
}

func TextField() Field {
	return Field{dbType: "TEXT", goType:"string"}
}


func (f Field) create() string {
	s := f.dbColumn + " " + f.dbType

	if f.primaryKey {
		s += " PRIMARY KEY"
	} else {

		if f.foreignKey {
			s += " REFERENCES " + f.to.dbTable + " ON DELETE " + f.onDelete
		}

		if f.null {
			s += " NULL"
		} else {
			s += " NOT NULL"
		}

		if f.unique {
			s += " UNIQUE"
		}

		if f.defaultValue != "" {
			s += " DEFAULT " + f.defaultValue
		}
	}

	return s
}
