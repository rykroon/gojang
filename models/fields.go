package models

import (
	"strconv"
)

//type constraint bool

type AutoField struct {
	dbColumn   string
	null       bool
	primaryKey bool
	unique     bool
}

type BooleanField struct {
	dbColumn   string
	null       bool
	primaryKey bool
	unique     bool
}

type CharField struct {
	dbColumn  string
	maxLength int

	null       bool
	primaryKey bool
	unique     bool
}

type DecimalField struct {
	dbColumn      string
	maxDigits     int
	decimalPlaces int

	null       bool
	primaryKey bool
	unique     bool
}

type FloatField struct {
	dbColumn string

	null       bool
	primaryKey bool
	unique     bool
}

type IntegerField struct {
	dbColumn string

	null       bool
	primaryKey bool
	unique     bool
}

type TextField struct {
	dbColumn string

	null       bool
	primaryKey bool
	unique     bool
}

func constraintString(pkey, null, unique bool) string {
	s := ""

	if pkey {
		s += " PRIMARY KEY"
	} else {

		if null {
			s += " NULL"
		} else {
			s += " NOT NULL"
		}

		if unique {
			s += " UNIQUE"
		}

	}
	return s
}

func doubleQuote(s string) string {
	return "\"" + s + "\""
}

//AutoField

func (f AutoField) Init(column string) AutoField {
	f.dbColumn = column

	f.null = false
	f.primaryKey = false
	f.unique = false

	return f
}

func (f AutoField) PrimaryKey(b bool) AutoField {
	f.primaryKey = b
	return f
}

func (f AutoField) Null(b bool) AutoField {
	f.null = b
	return f
}

func (f AutoField) Unique(b bool) AutoField {
	f.unique = b
	return f
}

func (f AutoField) CreateString() string {
	s := ""
	s += doubleQuote(f.dbColumn) + " SERIAL"

	s += constraintString(f.primaryKey, f.null, f.unique)

	return s
}

//BooleanField

func (f BooleanField) Init(column string) BooleanField {
	f.dbColumn = column

	f.null = false
	f.primaryKey = false
	f.unique = false

	return f
}

func (f BooleanField) PrimaryKey(b bool) BooleanField {
	f.primaryKey = b
	return f
}

func (f BooleanField) Null(b bool) BooleanField {
	f.null = b
	return f
}

func (f BooleanField) Unique(b bool) BooleanField {
	f.unique = b
	return f
}

func (f BooleanField) CreateString() string {
	s := ""
	s += doubleQuote(f.dbColumn) + " boolean"

	s += constraintString(f.primaryKey, f.null, f.unique)

	return s
}

//CharField

func (f CharField) Init(column string, maxLength int) CharField {
	f.dbColumn = column
	f.maxLength = maxLength

	f.null = false
	f.primaryKey = false
	f.unique = false

	return f
}

func (f CharField) PrimaryKey(b bool) CharField {
	f.primaryKey = b
	return f
}

func (f CharField) Null(b bool) CharField {
	f.null = b
	return f
}

func (f CharField) Unique(b bool) CharField {
	f.unique = b
	return f
}

func (f CharField) CreateString() string {
	s := ""
	n := strconv.Itoa(f.maxLength)
	s += doubleQuote(f.dbColumn) + " varchar(" + n + ")"

	s += constraintString(f.primaryKey, f.null, f.unique)

	return s
}


//DecimalField

func (f DecimalField) Init(column string, maxDigits int, decimalPlaces int) DecimalField {
	f.dbColumn = column
	f.maxDigits = maxDigits
	f.decimalPlaces = decimalPlaces

	f.null = false
	f.primaryKey = false
	f.unique = false
	return f
}

func (f DecimalField) PrimaryKey(b bool) DecimalField {
	f.primaryKey = b
	return f
}

func (f DecimalField) Null(b bool) DecimalField {
	f.null = b
	return f
}

func (f DecimalField) Unique(b bool) DecimalField {
	f.unique = b
	return f
}

func (f DecimalField) CreateString() string {
	s := ""
	precision := strconv.Itoa(f.maxDigits)
	scale := strconv.Itoa(f.decimalPlaces)

	s += doubleQuote(f.dbColumn) + " NUMERIC(" + precision + " " + scale + ")"

	s += constraintString(f.primaryKey, f.null, f.unique)

	return s
}


//FloatField

func (f FloatField) Init(column string) FloatField {
	f.dbColumn = column
	f.null = false
	f.primaryKey = false
	f.unique = false
	return f
}

func (f FloatField) PrimaryKey(b bool) FloatField {
	f.primaryKey = b
	return f
}

func (f FloatField) Null(b bool) FloatField {
	f.null = b
	return f
}

func (f FloatField) Unique(b bool) FloatField {
	f.unique = b
	return f
}

func (f FloatField) CreateString() string {
	s := ""
	s += doubleQuote(f.dbColumn) + " double precision"

	s += constraintString(f.primaryKey, f.null, f.unique)

	return s
}

//IntegerField

func (f IntegerField) Init(column string) IntegerField {
	f.dbColumn = column
	f.null = false
	f.primaryKey = false
	f.unique = false
	return f
}

func (f IntegerField) PrimaryKey(b bool) IntegerField {
	f.primaryKey = b
	return f
}

func (f IntegerField) Null(b bool) IntegerField {
	f.null = b
	return f
}

func (f IntegerField) Unique(b bool) IntegerField {
	f.unique = b
	return f
}

func (f IntegerField) CreateString() string {
	s := ""
	s += doubleQuote(f.dbColumn) + " integer"

	s += constraintString(f.primaryKey, f.null, f.unique)

	return s
}

//TextField

func (f TextField) Init(column string) TextField {
	f.dbColumn = column
	f.null = false
	f.primaryKey = false
	f.unique = false
	return f
}

func (f TextField) PrimaryKey(b bool) TextField {
	f.primaryKey = b
	return f
}

func (f TextField) Null(b bool) TextField {
	f.null = b
	return f
}

func (f TextField) Unique(b bool) TextField {
	f.unique = b
	return f
}

func (f TextField) CreateString() string {
	s := ""
	s += doubleQuote(f.dbColumn) + " text"

	s += constraintString(f.primaryKey, f.null, f.unique)

	return s
}
