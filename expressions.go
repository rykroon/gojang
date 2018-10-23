package gojang

import (
	"fmt"
	"reflect"
	//"unsafe"
)

//Expressions describe a value or a computation that can be used as part of an
//update, create, filter, order by, annotation, or aggregate.
type expression interface {
	asSql() string
}

type selectExpression interface {
	expression
	Scan(interface{}) error
	getValue() interface{}
	//getPtr() unsafe.Pointer
	getGoType() string
}

type sortExpression struct {
	field field
	desc  bool
}

type star string

func (s star) asSql() string {
	return "*"
}

func (a aggregate) asSql() string {
	return fmt.Sprintf(a.template, a.function, a.expression.asSql(), a.alias)
}

func (l lookup) asSql() string {
	sql := l.lhs.asSql() + " " + l.lookupName + " " + l.rhs

	if l.not {
		sql = "NOT(" + sql + ")"
	}
	return sql
}

func (e sortExpression) asSql() string {
	if e.desc {
		return e.field.asSql() + " DESC"
	} else {
		return e.field.asSql() + " ASC"
	}
}

func (f *AutoField) asSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *BigAutoField) asSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *BigIntegerField) asSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *BooleanField) asSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *FloatField) asSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *IntegerField) asSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *SmallIntegerField) asSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *TextField) asSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *ForeignKey) asSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (f *OneToOneField) asSql() string {
	return dbq(f.model.dbTable) + "." + dbq(f.dbColumn)
}

func (a aggregate) Scan(value interface{}) error {
	return a.outputField.Scan(value)
}

func (f *AutoField) Scan(value interface{}) error {
	result, ok := value.(int64)
	f.Value, f.Valid = int32(result), ok
	return nil
}

func (f *BigAutoField) Scan(value interface{}) error {
	f.Value, f.Valid = value.(int64)
	return nil
}

func (f *BigIntegerField) Scan(value interface{}) error {
	f.Value, f.Valid = value.(int64)
	return nil
}

func (f *BooleanField) Scan(value interface{}) error {
	f.Value, f.Valid = value.(bool)
	return nil
}

func (f *FloatField) Scan(value interface{}) error {
	f.Value, f.Valid = value.(float64)
	return nil
}

func (f *IntegerField) Scan(value interface{}) error {
	result, ok := value.(int64)
	f.Value, f.Valid = int32(result), ok
	return nil
}

func (f *SmallIntegerField) Scan(value interface{}) error {
	result, ok := value.(int64)
	f.Value, f.Valid = int16(result), ok
	return nil
}

func (f *TextField) Scan(value interface{}) error {
	f.Value, f.Valid = value.(string)
	return nil
}

func (f *ForeignKey) Scan(value interface{}) error {
	f.Value, f.Valid = value.(int64)
	return nil
}

func (f *OneToOneField) Scan(value interface{}) error {
	f.Value, f.Valid = value.(int64)
	return nil
}

func (a aggregate) getValue() interface{} {
	return a.outputField.getValue()
}

func (f *AutoField) getValue() interface{} {
	return f.Value
}

func (f *BigAutoField) getValue() interface{} {
	return f.Value
}

func (f *BigIntegerField) getValue() interface{} {
	return f.Value
}

func (f *BooleanField) getValue() interface{} {
	return f.Value
}

func (f *FloatField) getValue() interface{} {
	return f.Value
}

func (f *IntegerField) getValue() interface{} {
	return f.Value
}

func (f *SmallIntegerField) getValue() interface{} {
	return f.Value
}

func (f *TextField) getValue() interface{} {
	return f.Value
}

func (f *ForeignKey) getValue() interface{} {
	return f.Value
}

func (f *OneToOneField) getValue() interface{} {
	return f.Value
}

// func (a aggregate) getPtr() unsafe.Pointer {
// 	return a.outputField.getPtr()
// }
//
// func (f *AutoField) getPtr() unsafe.Pointer {
// 	return unsafe.Pointer(f.pointer)
// }
//
// func (f *BigAutoField) getPtr() unsafe.Pointer {
// 	return unsafe.Pointer(f.pointer)
// }
//
// func (f *BigIntegerField) getPtr() unsafe.Pointer {
// 	return unsafe.Pointer(f.pointer)
// }
//
// func (f *BooleanField) getPtr() unsafe.Pointer {
// 	return unsafe.Pointer(f.pointer)
// }
//
// func (f *FloatField) getPtr() unsafe.Pointer {
// 	return unsafe.Pointer(f.pointer)
// }
//
// func (f *IntegerField) getPtr() unsafe.Pointer {
// 	return unsafe.Pointer(f.pointer)
// }
//
// func (f *SmallIntegerField) getPtr() unsafe.Pointer {
// 	return unsafe.Pointer(f.pointer)
// }
//
// func (f *TextField) getPtr() unsafe.Pointer {
// 	return unsafe.Pointer(f.pointer)
// }
//
// func (f *ForeignKey) getPtr() unsafe.Pointer {
// 	return unsafe.Pointer(f.pointer)
// }
//
// func (f *OneToOneField) getPtr() unsafe.Pointer {
// 	return unsafe.Pointer(f.pointer)
// }

func (a aggregate) getGoType() string {
	return a.outputField.getGoType()
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
