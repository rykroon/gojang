package models

import (
  "strconv"
)

type Field struct {
	dbColumn string
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

func AutoField(dbColumn string) Field {
	return Field{dbColumn: dbColumn, dbDataType: "SERIAL"}
}

func BooleanField(dbColumn string) Field {
  return Field{dbColumn: dbColumn, dbDataType: "boolean"}
}

func CharField(dbColumn string, maxLength int) Field {
  n := strconv.Itoa(maxLength)
  dataType := "varchar(" + n + ")"

  return Field{dbColumn: dbColumn, dbDataType: dataType, maxLength: maxLength}
}

func DecimalField(dbColumn string, maxDigits int, decimalPlaces int) Field {
  precision := strconv.Itoa(maxDigits)
  scale := strconv.Itoa(decimalPlaces)
  dataType := "NUMERIC(" + precision + ", " + scale + ")"

  return Field{dbColumn: dbColumn, dbDataType: dataType, maxDigits: maxDigits, decimalPlaces: decimalPlaces}
}

func FloatField(dbColumn string) Field {
  return Field{dbColumn: dbColumn, dbDataType: "double precision"}
}

func IntegerField(dbColumn string) Field {
  return Field{dbColumn: dbColumn, dbDataType: "integer"}
}

func TextField(dbColumn string) Field {
  return Field{dbColumn: dbColumn, dbDataType: "text"}
}


func doubleQuotes(s string) string {
  return "\"" + s + "\""
}

func (f Field) CreateString() string {
  s := doubleQuotes(f.dbColumn) + " " + f.dbDataType

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
