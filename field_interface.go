package gojang

import (
	"reflect"
	"unsafe"
)

type field interface {
	hasNullConstraint() bool
	hasUniqueConstraint() bool
	hasPrimaryKeyConstraint() bool
	hasRelation() bool

	getDbColumn() string
	setDbColumn(string)
	getDbType() string
	getGoType() string

	IsNil() bool
	//SetNil() error

	Asc() orderByExpression
	Desc() orderByExpression

	getPtr() unsafe.Pointer
	sqlValue() string
}


// type intField interface {
// 	Val() int
// }

type primaryKeyField interface {
	id() int
	Val() int
}

func (f *AutoField) id() int {
	return int(f.value)
}

func (f *BigAutoField) id() int {
	return int(f.value)
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

func (f *ForeignKeyField) hasNullConstraint() bool {
	return f.null
}

func (f *OneToOneField) hasNullConstraint() bool {
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

func (f *ForeignKeyField) hasUniqueConstraint() bool {
	return f.unique
}

func (f *OneToOneField) hasUniqueConstraint() bool {
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

func (f *ForeignKeyField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *OneToOneField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *AutoField) hasRelation() bool {
	return f.primaryKey
}

func (f *BigAutoField) hasRelation() bool {
	return f.isRelation
}

func (f *BigIntegerField) hasRelation() bool {
	return f.isRelation
}

func (f *BooleanField) hasRelation() bool {
	return f.isRelation
}

func (f *FloatField) hasRelation() bool {
	return f.isRelation
}

func (f *IntegerField) hasRelation() bool {
	return f.isRelation
}

func (f *TextField) hasRelation() bool {
	return f.isRelation
}

func (f *ForeignKeyField) hasRelation() bool {
	return f.isRelation
}

func (f *OneToOneField) hasRelation() bool {
	return f.isRelation
}

func (f *AutoField) getDbColumn() string {
	return f.dbColumn
}

func (f *BigAutoField) getDbColumn() string {
	return f.dbColumn
}

func (f *BigIntegerField) getDbColumn() string {
	return f.dbColumn
}

func (f *BooleanField) getDbColumn() string {
	return f.dbColumn
}

func (f *FloatField) getDbColumn() string {
	return f.dbColumn
}

func (f *IntegerField) getDbColumn() string {
	return f.dbColumn
}

func (f *TextField) getDbColumn() string {
	return f.dbColumn
}

func (f *ForeignKeyField) getDbColumn() string {
	return f.dbColumn
}

func (f *OneToOneField) getDbColumn() string {
	return f.dbColumn
}

func (f *AutoField) setDbColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *BigAutoField) setDbColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *BigIntegerField) setDbColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *BooleanField) setDbColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *FloatField) setDbColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *IntegerField) setDbColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *TextField) setDbColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *ForeignKeyField) setDbColumn(columnName string) {
	f.dbColumn = columnName
}

func (f OneToOneField) setDbColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *AutoField) getDbType() string {
	return f.dbType
}

func (f *BigAutoField) getDbType() string {
	return f.dbType
}

func (f *BigIntegerField) getDbType() string {
	return f.dbType
}

func (f *BooleanField) getDbType() string {
	return f.dbType
}

func (f *FloatField) getDbType() string {
	return f.dbType
}

func (f *IntegerField) getDbType() string {
	return f.dbType
}

func (f *TextField) getDbType() string {
	return f.dbType
}

func (f *ForeignKeyField) getDbType() string {
	return f.dbType
}

func (f *OneToOneField) getDbType() string {
	return f.dbType
}

func (f *AutoField) getGoType() string {
	return reflect.TypeOf(f.value).String()
}

func (f *BigAutoField) getGoType() string {
	return reflect.TypeOf(f.value).String()
}

func (f *BigIntegerField) getGoType() string {
	return reflect.TypeOf(f.value).String()
}

func (f *BooleanField) getGoType() string {
	return reflect.TypeOf(f.value).String()
}

func (f *FloatField) getGoType() string {
	return reflect.TypeOf(f.value).String()
}

func (f *IntegerField) getGoType() string {
	return reflect.TypeOf(f.value).String()
}

func (f *TextField) getGoType() string {
	return reflect.TypeOf(f.value).String()
}

func (f *ForeignKeyField) getGoType() string {
	return reflect.TypeOf(f.value).String()
}

func (f *OneToOneField) getGoType() string {
	return reflect.TypeOf(f.value).String()
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

func (f *ForeignKeyField) IsNil() bool {
	return f.pointer == nil
}

func (f *OneToOneField) IsNil() bool {
	return f.pointer == nil
}

func (f *AutoField) getPtr() unsafe.Pointer {
	return unsafe.Pointer(f.pointer)
}

func (f *BigAutoField) getPtr() unsafe.Pointer {
	return unsafe.Pointer(f.pointer)
}

func (f *BigIntegerField) getPtr() unsafe.Pointer {
	return unsafe.Pointer(f.pointer)
}

func (f *BooleanField) getPtr() unsafe.Pointer {
	return unsafe.Pointer(f.pointer)
}

func (f *FloatField) getPtr() unsafe.Pointer {
	return unsafe.Pointer(f.pointer)
}

func (f *IntegerField) getPtr() unsafe.Pointer {
	return unsafe.Pointer(f.pointer)
}

func (f *TextField) getPtr() unsafe.Pointer {
	return unsafe.Pointer(f.pointer)
}

func (f *ForeignKeyField) getPtr() unsafe.Pointer {
	return unsafe.Pointer(f.pointer)
}

func (f *OneToOneField) getPtr() unsafe.Pointer {
	return unsafe.Pointer(f.pointer)
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

func (f *ForeignKeyField) sqlValue() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return int64ToSql(f.value)
	}
}

func (f *OneToOneField) sqlValue() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return int64ToSql(f.value)
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

func (f *ForeignKeyField) Asc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "ASC")
}

func (f *OneToOneField) Asc() orderByExpression {
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

func (f *TextField) Desc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "DESC")
}

func (f *ForeignKeyField) Desc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "DESC")
}

func (f *OneToOneField) Desc() orderByExpression {
	return orderByExpression(doubleQuotes(f.dbColumn) + "DESC")
}
