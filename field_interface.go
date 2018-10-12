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

	setModel(*Model)
	getDbColumn() string
	setDbColumn(string)
	getDbType() string
	getGoType() string

	IsNil() bool
	//SetNil() error

	Asc() sortExpression
	Desc() sortExpression

	getPtr() unsafe.Pointer
	sqlField() string
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

func (f *AutoField) setModel(model *Model) {
	f.model = model
}

func (f *BigAutoField) setModel(model *Model) {
	f.model = model
}

func (f *BigIntegerField) setModel(model *Model) {
	f.model = model
}

func (f *BooleanField) setModel(model *Model) {
	f.model = model
}

func (f *FloatField) setModel(model *Model) {
	f.model = model
}

func (f *IntegerField) setModel(model *Model) {
	f.model = model
}

func (f *TextField) setModel(model *Model) {
	f.model = model
}

func (f *ForeignKeyField) setModel(model *Model) {
	f.model = model
}

func (f *OneToOneField) setModel(model *Model) {
	f.model = model
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

func (f *AutoField) sqlField() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *BigAutoField) sqlField() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *BigIntegerField) sqlField() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *BooleanField) sqlField() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *FloatField) sqlField() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *IntegerField) sqlField() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *TextField) sqlField() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *ForeignKeyField) sqlField() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *OneToOneField) sqlField() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *AutoField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *BigAutoField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *BigIntegerField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *BooleanField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *FloatField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *IntegerField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *TextField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *ForeignKeyField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *OneToOneField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *AutoField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *BigAutoField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *BigIntegerField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *BooleanField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *FloatField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *IntegerField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *TextField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *ForeignKeyField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *OneToOneField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}
