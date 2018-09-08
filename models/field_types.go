package models

import (
	"strconv"
)

//maybe?
type FieldInterface interface {
	isAField()
}

type AutoField struct {
	Field
	defaultVal int
}

type BooleanField struct {
	Field
	defaultVal bool
}

type CharField struct {
	Field
	maxLength int

	defaultVal string
}

type DecimalField struct {
	Field
	maxDigits     int
	decimalPlaces int

	defaultVal float64
}

type FloatField struct {
	Field
	defaultVal float64
}

// type ForeignKeyField struct {
// 	Field
// 	to       *Model
// 	onDelete string
// }

type IntegerField struct {
	Field
	defaultVal int
}

type TextField struct {
	Field
	defaultVal string
}

type Field struct {
	dbColumn string
	dbType   string
	model    *Model

	//constraints
	null       bool
	primaryKey bool
	unique     bool

	foreignKey bool
}

//Dummy function so that all Field can be under the FieldInterface interface

func (f AutoField) isAField() {
	return
}

func (f BooleanField) isAField() {
	return
}

func (f CharField) isAField()  {
	return
}

func (f DecimalField) isAField()  {
	return
}

func (f FloatField) isAField() {
	return
}

func (f IntegerField) isAField() {
	return
}

func (f TextField) isAField() {
	return
}

func NewField(options []fieldOption) Field {
	f := Field{}

	for _, option := range options {
		option(&f)
	}

	return f
}

func NewAutoField(options ...fieldOption) AutoField {
	f := AutoField{}
	f.Field = NewField(options)
	f.dbType = "SERIAL"
	return f
}

func NewBooleanField(options ...fieldOption) BooleanField {
	f := BooleanField{}
	f.Field = NewField(options)
	f.dbType = "BOOLEAN"
	return f
}

func NewCharField(maxLength int, options ...fieldOption) CharField {
	f := CharField{}
	f.Field = NewField(options)

	n := strconv.Itoa(maxLength)
	dataType := "VARCHAR(" + n + ")"
	f.dbType = dataType

	return f
}

func NewDecimalField(maxDigits int, decimalPlaces int, options ...fieldOption) DecimalField {
	f := DecimalField{maxDigits: maxDigits, decimalPlaces: decimalPlaces}
	f.Field = NewField(options)

	precision := strconv.Itoa(maxDigits)
	scale := strconv.Itoa(decimalPlaces)
	dataType := "NUMERIC(" + precision + ", " + scale + ")"

	f.dbType = dataType
	return f
}

func NewFloatField(options ...fieldOption) FloatField {
	f := FloatField{}
	f.Field = NewField(options)
	f.dbType = "DOUBLE PRECISION"
	return f
}

// func NewForeignKeyField(options ...fieldOption) ForeignKeyField {
// 	f := ForeignKeyField{}
// 	f.Field = NewField(options)
// 	f.dbType = "INTEGER"
// 	return f
// }

func NewIntegerField(options ...fieldOption) IntegerField {
	f := IntegerField{}
	f.Field = NewField(options)
	f.dbType = "INTEGER"
	return f
}

func NewTextField(options ...fieldOption) TextField {
	f := TextField{}
	f.Field = NewField(options)
	f.dbType = "TEXT"
	return f
}

//Field Options

// func (f Field) createString(dbColumn string) string {
// 	s := dbColumn + " " + f.dbType
//
// 	if f.primaryKey {
// 		s += " PRIMARY KEY"
// 	} else {
//
// 		if f.foreignKey {
// 			s += " REFERENCES " + f.to.dbTable + " ON DELETE " + f.onDelete
// 		}
//
// 		if f.null {
// 			s += " NULL"
// 		} else {
// 			s += " NOT NULL"
// 		}
//
// 		if f.unique {
// 			s += " UNIQUE"
// 		}
// 	}
//
// 	return s
// }
