package models

import (
	"strconv"
)


//type constraint bool

// type AutoField struct {
// 	dbColumn string
// 	null     bool
// 	primaryKey bool
// 	unique     bool
// }

// type BooleanField struct {
//   dbColumn string
// 	null     bool
// 	primaryKey bool
// 	unique     bool
// }

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

//CharField

func (f CharField) Init(column string, maxLength int) CharField {
	f.dbColumn = column
	f.maxLength = maxLength

	f.null = false
	//f.default_ = nil
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
	s += f.dbColumn + " varchar(" + n + ")"

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

//DecimalField

func (f DecimalField) Init(column string, maxDigits int, decimalPlaces int) DecimalField {
	f.dbColumn = column
	f.maxDigits = maxDigits
	f.decimalPlaces = decimalPlaces

	f.null = false
	//f.default_ = nil
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

	s += f.dbColumn + " NUMERIC(" + precision + " " + scale + ")"

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

//FloatField

func (f FloatField) Init(column string) FloatField {
	f.dbColumn = column
	f.null = false
	//f.default_ = nil
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
	s += f.dbColumn + " double precision"

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

//IntegerField

func (f IntegerField) Init(column string) IntegerField {
	f.dbColumn = column
	f.null = false
	//f.default_ = nil
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
	s += f.dbColumn + " double precision"

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

//TextField

func (f TextField) Init(column string) TextField {
	f.dbColumn = column
	f.null = false
	//f.default_ = nil
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
	s += f.dbColumn + " double precision"

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
