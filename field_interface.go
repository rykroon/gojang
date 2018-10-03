package gojang

import (
//"reflect"
)

type field interface {
	HasNullConstraint() bool
	HasUniqueConstraint() bool
	HasPrimaryKeyConstraint() bool
	getDBColumn() string
	getDBType() string
	IsNil() bool
	Asc() orderByExpression
	Desc() orderByExpression
	sqlValue() string
}

func (f BooleanField) HasNullConstraint() bool {
	return f.null
}

func (f FloatField) HasNullConstraint() bool {
	return f.null
}

func (f *IntegerField) HasNullConstraint() bool {
	return f.null
}

func (f *TextField) HasNullConstraint() bool {
	return f.null
}



func (f BooleanField) HasUniqueConstraint() bool {
	return f.unique
}

func (f FloatField) HasUniqueConstraint() bool {
	return f.unique
}

func (f *IntegerField) HasUniqueConstraint() bool {
	return f.unique
}

func (f *TextField) HasUniqueConstraint() bool {
	return f.unique
}



func (f BooleanField) HasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f FloatField) HasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *IntegerField) HasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *TextField) HasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f BooleanField) getDBColumn() string {
	return f.dbColumn
}

func (f FloatField) getDBColumn() string {
	return f.dbColumn
}

func (f *IntegerField) getDBColumn() string {
	return f.dbColumn
}

func (f *TextField) getDBColumn() string {
	return f.dbColumn
}



func (f BooleanField) getDBType() string {
	return f.dbType
}

func (f FloatField) getDBType() string {
	return f.dbType
}

func (f *IntegerField) getDBType() string {
	return f.dbType
}

func (f *TextField) getDBType() string {
	return f.dbType
}


func (f BooleanField) IsNil() bool {
	return f.pointer == nil
}

func (f FloatField) IsNil() bool {
	return f.pointer == nil
}

func (f *IntegerField) IsNil() bool {
	return f.pointer == nil
}

func (f *TextField) IsNil() bool {
	return f.pointer == nil
}

func (f BooleanField) sqlValue() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return boolToSql(f.value)
	}
}

func (f FloatField) sqlValue() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return float64ToSql(f.value)
	}
}

func (f *IntegerField) sqlValue() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return int32ToSql(f.value)
	}
}

func (f *TextField) sqlValue() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return stringToSql(f.value)
	}
}



type orderByExpression string

func (f BooleanField) Asc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "ASC")
}

func (f FloatField) Asc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "ASC")
}

func (f *IntegerField) Asc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "ASC")
}

func (f *TextField) Asc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "ASC")
}



func (f BooleanField) Desc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "DESC")
}

func (f FloatField) Desc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "DESC")
}

func (f *IntegerField) Desc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "DESC")
}

func (f *TextField) Desc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "DESC")
}
