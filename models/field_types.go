package models

import (
	"reflect"
	"strconv"
)

type Field struct {
	dbColumn string
	dbType   string

	//specific for CharField
	maxLength int

	//specfic for DecimalField
	maxDigits     int
	decimalPlaces int

	//constraints
	null         bool
	primaryKey   bool
	unique       bool
	defaultType  reflect.Kind
	defaultValue string

	foreignKey bool
	to         *Model
	onDelete   string
}

func AutoField() Field {
	return Field{dbType: "SERIAL", defaultType: reflect.Int}
}

func BooleanField() Field {
	return Field{dbType: "BOOLEAN", defaultType: reflect.Bool}
}

func CharField(maxLength int) Field {
	n := strconv.Itoa(maxLength)
	dataType := "VARCHAR(" + n + ")"

	return Field{dbType: dataType, maxLength: maxLength, defaultType: reflect.String}
}

func DecimalField(maxDigits int, decimalPlaces int) Field {
	precision := strconv.Itoa(maxDigits)
	scale := strconv.Itoa(decimalPlaces)
	dataType := "NUMERIC(" + precision + ", " + scale + ")"

	return Field{dbType: dataType, maxDigits: maxDigits, decimalPlaces: decimalPlaces, defaultType: reflect.Float64}
}

func FloatField() Field {
	return Field{dbType: "DOUBLE PRECISION", defaultType: reflect.Float64}
}

func ForeignKey(m *Model, onDelete string) Field {
	return Field{dbType: "INTEGER", foreignKey: true, to: m, onDelete: onDelete, defaultType: reflect.Int}
}

func IntegerField() Field {
	return Field{dbType: "INTEGER", defaultType: reflect.Int}
}

func TextField() Field {
	return Field{dbType: "TEXT", defaultType: reflect.String}
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
