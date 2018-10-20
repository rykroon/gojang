package gojang

import (
	"reflect"
	"unsafe"
)

type field interface {
	hasNullConstraint() bool
	setNullConstraint(bool)
	hasUniqueConstraint() bool
	setUniqueConstraint(bool)
	hasPrimaryKeyConstraint() bool
	setPrimaryKeyConstraint(bool)
	hasRelation() bool
	validate()

	setModel(*Model)
	getDbColumn() string
	setDbColumn(string)
	getDbType() string
	getGoType() string
	getPtr() unsafe.Pointer

	IsNil() bool

	Asc() sortExpression
	Desc() sortExpression

	toSql() string
	valueToSql() string
}

// type intField interface {
// 	Val() int
// }

type primaryKeyField interface {
	id() int
	Val() int
	Exact(int) lookup
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

func (f *SmallIntegerField) hasNullConstraint() bool {
	return f.null
}

func (f *TextField) hasNullConstraint() bool {
	return f.null
}

func (f *ForeignKey) hasNullConstraint() bool {
	return f.null
}

func (f *OneToOneField) hasNullConstraint() bool {
	return f.null
}

func (f *AutoField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.pointer = nil
	}
}

func (f *BigAutoField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.pointer = nil
	}
}

func (f *BigIntegerField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.pointer = nil
	}
}

func (f *BooleanField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.pointer = nil
	}
}

func (f *FloatField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.pointer = nil
	}
}

func (f *IntegerField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.pointer = nil
	}
}

func (f *SmallIntegerField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.pointer = nil
	}
}

func (f *TextField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.pointer = nil
	}
}

func (f *ForeignKey) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.pointer = nil
	}
}

func (f *OneToOneField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.pointer = nil
	}
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

func (f *SmallIntegerField) hasUniqueConstraint() bool {
	return f.unique
}

func (f *TextField) hasUniqueConstraint() bool {
	return f.unique
}

func (f *ForeignKey) hasUniqueConstraint() bool {
	return f.unique
}

func (f *OneToOneField) hasUniqueConstraint() bool {
	return f.unique
}

func (f *AutoField) setUniqueConstraint(unique bool) {
	f.unique = unique
}

func (f *BigAutoField) setUniqueConstraint(unique bool) {
	f.unique = unique
}

func (f *BigIntegerField) setUniqueConstraint(unique bool) {
	f.unique = unique
}

func (f *BooleanField) setUniqueConstraint(unique bool) {
	f.unique = unique
}

func (f *FloatField) setUniqueConstraint(unique bool) {
	f.unique = unique
}

func (f *IntegerField) setUniqueConstraint(unique bool) {
	f.unique = unique
}

func (f *SmallIntegerField) setUniqueConstraint(unique bool) {
	f.unique = unique
}

func (f *TextField) setUniqueConstraint(unique bool) {
	f.unique = unique
}

func (f *ForeignKey) setUniqueConstraint(unique bool) {
	f.unique = unique
}

//The 'Unique' field option is valid on all field types except ManyToManyField and OneToOneField.
func (f *OneToOneField) setUniqueConstraint(unique bool) {
	f.unique = unique
}

//The 'Unique' field option is valid on all field types except ManyToManyField and OneToOneField.
// func (f *ManyToManyField) setUniqueConstraint(unique bool) {
// 	f.unique = unique
//
// 	if !f.unique {
// 		panic("OneToOneField must be unique")
// 	}
// }

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

func (f *SmallIntegerField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *TextField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *ForeignKey) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *OneToOneField) hasPrimaryKeyConstraint() bool {
	return f.primaryKey
}

func (f *AutoField) setPrimaryKeyConstraint(primaryKey bool) {
	f.primaryKey = primaryKey
}

func (f *BigAutoField) setPrimaryKeyConstraint(primaryKey bool) {
	f.primaryKey = primaryKey
}

func (f *BigIntegerField) setPrimaryKeyConstraint(primaryKey bool) {
	f.primaryKey = primaryKey
}

func (f *BooleanField) setPrimaryKeyConstraint(primaryKey bool) {
	f.primaryKey = primaryKey
}

func (f *FloatField) setPrimaryKeyConstraint(primaryKey bool) {
	f.primaryKey = primaryKey
}

func (f *IntegerField) setPrimaryKeyConstraint(primaryKey bool) {
	f.primaryKey = primaryKey
}

func (f *SmallIntegerField) setPrimaryKeyConstraint(primaryKey bool) {
	f.primaryKey = primaryKey
}

func (f *TextField) setPrimaryKeyConstraint(primaryKey bool) {
	f.primaryKey = primaryKey
}

func (f *ForeignKey) setPrimaryKeyConstraint(primaryKey bool) {
	f.primaryKey = primaryKey
}

func (f *OneToOneField) setPrimaryKeyConstraint(primaryKey bool) {
	f.primaryKey = primaryKey
}

func (f *AutoField) hasRelation() bool {
	return f.isRelation
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

func (f *SmallIntegerField) hasRelation() bool {
	return f.isRelation
}

func (f *TextField) hasRelation() bool {
	return f.isRelation
}

func (f *ForeignKey) hasRelation() bool {
	return f.isRelation
}

func (f *OneToOneField) hasRelation() bool {
	return f.isRelation
}

func (f *AutoField) validate() {
	if !f.primaryKey {
		panic(NewForcePrimaryKey(f))
	}

	if f.primaryKey && f.null {
		panic(NewNullPrimaryKeyErr())
	}
}

func (f *BigAutoField) validate() {
	if !f.primaryKey {
		panic(NewForcePrimaryKey(f))
	}

	if f.primaryKey && f.null {
		panic(NewNullPrimaryKeyErr())
	}
}

func (f *BigIntegerField) validate() {
	if f.primaryKey {
		panic(NewInvalidPrimaryKey(f))
	}

	if f.primaryKey && f.null {
		panic(NewNullPrimaryKeyErr())
	}
}

func (f *BooleanField) validate() {
	if f.primaryKey {
		panic(NewInvalidPrimaryKey(f))
	}

	if f.primaryKey && f.null {
		panic(NewNullPrimaryKeyErr())
	}
}

func (f *FloatField) validate() {
	if f.primaryKey {
		panic(NewInvalidPrimaryKey(f))
	}

	if f.primaryKey && f.null {
		panic(NewNullPrimaryKeyErr())
	}
}

func (f *IntegerField) validate() {
	if f.primaryKey {
		panic(NewInvalidPrimaryKey(f))
	}

	if f.primaryKey && f.null {
		panic(NewNullPrimaryKeyErr())
	}
}

func (f *SmallIntegerField) validate() {
	if f.primaryKey {
		panic(NewInvalidPrimaryKey(f))
	}

	if f.primaryKey && f.null {
		panic(NewNullPrimaryKeyErr())
	}
}

func (f *TextField) validate() {
	if f.primaryKey {
		panic(NewInvalidPrimaryKey(f))
	}

	if f.primaryKey && f.null {
		panic(NewNullPrimaryKeyErr())
	}
}

func (f *ForeignKey) validate() {
	if f.primaryKey {
		panic(NewInvalidPrimaryKey(f))
	}

	if f.primaryKey && f.null {
		panic(NewNullPrimaryKeyErr())
	}
}

func (f *OneToOneField) validate() {
	if f.primaryKey {
		panic(NewInvalidPrimaryKey(f))
	}

	if f.primaryKey && f.null {
		panic(NewNullPrimaryKeyErr())
	}

	if !f.unique {
		panic("OneToOneField must be unique")
	}
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

func (f *SmallIntegerField) setModel(model *Model) {
	f.model = model
}

func (f *TextField) setModel(model *Model) {
	f.model = model
}

func (f *ForeignKey) setModel(model *Model) {
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

func (f *SmallIntegerField) getDbColumn() string {
	return f.dbColumn
}

func (f *TextField) getDbColumn() string {
	return f.dbColumn
}

func (f *ForeignKey) getDbColumn() string {
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

func (f *SmallIntegerField) setDbColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *TextField) setDbColumn(columnName string) {
	f.dbColumn = columnName
}

func (f *ForeignKey) setDbColumn(columnName string) {
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

func (f *SmallIntegerField) getDbType() string {
	return f.dbType
}

func (f *TextField) getDbType() string {
	return f.dbType
}

func (f *ForeignKey) getDbType() string {
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

func (f *SmallIntegerField) getGoType() string {
	return reflect.TypeOf(f.value).String()
}

func (f *TextField) getGoType() string {
	return reflect.TypeOf(f.value).String()
}

func (f *ForeignKey) getGoType() string {
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

func (f *SmallIntegerField) IsNil() bool {
	return f.pointer == nil
}

func (f *TextField) IsNil() bool {
	return f.pointer == nil
}

func (f *ForeignKey) IsNil() bool {
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

func (f *SmallIntegerField) getPtr() unsafe.Pointer {
	return unsafe.Pointer(f.pointer)
}

func (f *TextField) getPtr() unsafe.Pointer {
	return unsafe.Pointer(f.pointer)
}

func (f *ForeignKey) getPtr() unsafe.Pointer {
	return unsafe.Pointer(f.pointer)
}

func (f *OneToOneField) getPtr() unsafe.Pointer {
	return unsafe.Pointer(f.pointer)
}

func (f *AutoField) valueToSql() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return int32ToSql(f.value)
	}
}

func (f *BigAutoField) valueToSql() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return int64ToSql(f.value)
	}
}

func (f *BigIntegerField) valueToSql() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return int64ToSql(f.value)
	}
}

func (f *BooleanField) valueToSql() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return boolToSql(f.value)
	}
}

func (f *FloatField) valueToSql() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return float64ToSql(f.value)
	}
}

func (f *IntegerField) valueToSql() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return int32ToSql(f.value)
	}
}

func (f *SmallIntegerField) valueToSql() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return int16ToSql(f.value)
	}
}

func (f *TextField) valueToSql() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return stringToSql(f.value)
	}
}

func (f *ForeignKey) valueToSql() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return int64ToSql(f.value)
	}
}

func (f *OneToOneField) valueToSql() string {
	if f.IsNil() {
		return "NULL"
	} else {
		return int64ToSql(f.value)
	}
}

func (f *AutoField) toSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *BigAutoField) toSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *BigIntegerField) toSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *BooleanField) toSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *FloatField) toSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *IntegerField) toSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *SmallIntegerField) toSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *TextField) toSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *ForeignKey) toSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *OneToOneField) toSql() string {
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

func (f *SmallIntegerField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *TextField) Asc() sortExpression {
	return sortExpression{field: f}
}

func (f *ForeignKey) Asc() sortExpression {
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

func (f *SmallIntegerField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *TextField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *ForeignKey) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}

func (f *OneToOneField) Desc() sortExpression {
	return sortExpression{field: f, desc: true}
}
