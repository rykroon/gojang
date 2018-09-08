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

	foreignKey bool
	to         *Model
	onDelete   string
}

func AutoField() Field {
	return Field{dbDataType: "SERIAL"}
}

func BooleanField() Field {
	return Field{dbDataType: "BOOLEAN"}
}

func CharField(maxLength int) Field {
	n := strconv.Itoa(maxLength)
	dataType := "VARCHAR(" + n + ")"

	return Field{dbDataType: dataType, maxLength: maxLength}
}

func DecimalField(maxDigits int, decimalPlaces int) Field {
	precision := strconv.Itoa(maxDigits)
	scale := strconv.Itoa(decimalPlaces)
	dataType := "NUMERIC(" + precision + ", " + scale + ")"

	return Field{dbDataType: dataType, maxDigits: maxDigits, decimalPlaces: decimalPlaces}
}

func FloatField() Field {
	return Field{dbDataType: "DOUBLE PRECISION"}
}

func ForeignKey(m *Model, onDelete string) Field {
	return Field{dbDataType: "INTEGER", foreignKey: true, to: m, onDelete: onDelete}
}

func IntegerField() Field {
	return Field{dbDataType: "INTEGER"}
}

func TextField() Field {
	return Field{dbDataType: "TEXT"}
}


func (f Field) createString(dbColumn string) string {
	s := dbColumn + " " + f.dbDataType

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
	}

	return s
}
