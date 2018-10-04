package gojang

import (
//"reflect"
)

type field interface {
	hasNullConstraint() bool
	hasUniqueConstraint() bool
	hasPrimaryKeyConstraint() bool
	DBColumn() string
	setDBColumn(string)
	getDBType() string
	IsNil() bool
	Asc() orderByExpression
	Desc() orderByExpression
	sqlValue() string
}

func (f *AutoField) hasNullConstraint() bool {
	return f.null
}

func (f *BigAutoField) hasNullConstraint() bool {
	return f.null
}

func (f *BigIntegerField) hasNullConstraint() bool {
	return f.null
}

func (f *BooleanField) hasNullConstraint() bool {
	return f.null
}

func (f *FloatField) hasNullConstraint() bool {
	return f.null
}

func (f *IntegerField) hasNullConstraint() bool {
	return f.null
}

func (f *TextField) hasNullConstraint() bool {
	return f.null
}


func (f *AutoField) hasUniqueConstraint() bool {
	return f.unique
}

func (f *BigAutoField) hasUniqueConstraint() bool {
	return f.unique
}

func (f *BigIntegerField) hasUniqueConstraint() bool {
	return f.unique
}

func (f *BooleanField) hasUniqueConstraint() bool {
	return f.unique
}

func (f *FloatField) hasUniqueConstraint() bool {
	return f.unique
}

func (f *IntegerField) hasUniqueConstraint() bool {
	return f.unique
}

func (f *TextField) hasUniqueConstraint() bool {
	return f.unique
}


func (f *AutoField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *BigAutoField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *BigIntegerField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *BooleanField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *FloatField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *IntegerField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *TextField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *AutoField) DBColumn() string {
	return f.dbColumn
}

func (f *BigAutoField) DBColumn() string {
	return f.dbColumn
}

func (f *BigIntegerField) DBColumn() string {
	return f.dbColumn
}

func (f *BooleanField) DBColumn() string {
	return f.dbColumn
}

func (f *FloatField) DBColumn() string {
	return f.dbColumn
}

func (f *IntegerField) DBColumn() string {
	return f.dbColumn
}

func (f *TextField) DBColumn() string {
	return f.dbColumn
}


func (f *AutoField) setDBColumn(columnName string)  {
	f.dbColumn = columnName
}

func (f *BigAutoField) setDBColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *BigIntegerField) setDBColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *BooleanField) setDBColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *FloatField) setDBColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *IntegerField) setDBColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *TextField) setDBColumn(columnName string) {
	f.dbColumn = columnName
}


func (f *AutoField) getDBType() string {
	return f.dbType
}

func (f *BigAutoField) getDBType() string {
	return f.dbType
}

func (f *BigIntegerField) getDBType() string {
	return f.dbType
}

func (f *BooleanField) getDBType() string {
	return f.dbType
}

func (f *FloatField) getDBType() string {
	return f.dbType
}

func (f *IntegerField) getDBType() string {
	return f.dbType
}

func (f *TextField) getDBType() string {
	return f.dbType
}

func (f *AutoField) IsNil() bool {
	return f.pointer == nil
}

func (f *BigAutoField) IsNil() bool {
	return f.pointer == nil
}

func (f *BigIntegerField) IsNil() bool {
	return f.pointer == nil
}

func (f *BooleanField) IsNil() bool {
	return f.pointer == nil
}

func (f *FloatField) IsNil() bool {
	return f.pointer == nil
}

func (f *IntegerField) IsNil() bool {
	return f.pointer == nil
}

func (f *TextField) IsNil() bool {
	return f.pointer == nil
}

func (f *AutoField) sqlValue() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return int32ToSql(f.value)
	}
}

func (f *BigAutoField) sqlValue() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return int64ToSql(f.value)
	}
}

func (f *BigIntegerField) sqlValue() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return int64ToSql(f.value)
	}
}

func (f *BooleanField) sqlValue() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return boolToSql(f.value)
	}
}

func (f *FloatField) sqlValue() string {
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

func (f *AutoField) Asc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "ASC")
}

func (f *BigAutoField) Asc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "ASC")
}

func (f *BigIntegerField) Asc() orderByExpression {
	return orderByExpression(dbq(f.dbColumn) + "ASC")
}

func (f *BooleanField) Asc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "ASC")
}

func (f *FloatField) Asc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "ASC")
}

func (f *IntegerField) Asc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "ASC")
}

func (f *TextField) Asc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "ASC")
}

func (f *AutoField) Desc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "DESC")
}

func (f *BigAutoField) Desc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "DESC")
}

func (f *BigIntegerField) Desc() orderByExpression {
	return orderByExpression(dbq(f.dbColumn) + "DESC")
}

func (f *BooleanField) Desc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "DESC")
}

func (f *FloatField) Desc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "DESC")
}

func (f *IntegerField) Desc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "DESC")
}

func (f TextField) Desc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "DESC")
}
