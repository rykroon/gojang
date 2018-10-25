package gojang

import (
	"reflect"
)

//A object has an underlying type of reflect.Value, whose underlying type is a struct
type object reflect.Value

func NewObj(typ reflect.Type) (object, bool) {
	if typ.Kind() == reflect.Struct {
		return object(reflect.New(typ).Elem()), true
	}

	var obj object
	return obj, false
}

//Returns a list of the objects attributes
func (o object) Dirs() []string {
	var result []string
	t := o.Type()

	for i := 0; i < t.NumField(); i++ {
		result = append(result, t.Field(i).Name)
	}

	return result
}

//returns the attribute if it exists
func (o object) GetAttr(name string) reflect.Value {
	return reflect.Value(o).FieldByName(name)
}

//Returns true if the object has the attribute
func (o object) HasAttr(name string) bool {
	return reflect.Value(o).FieldByName(name).IsValid()
}

//Returns true if it was able to set the new value to the attribute, otherwise false
func (o object) SetAttr(name string, v interface{}) bool {
	if !o.HasAttr(name) {
		return false
	}

	attr := o.GetAttr(name)
	if !attr.CanSet() {
		return false
	}

	if attr.Type() != reflect.TypeOf(v) {
		return false
	}

	attr.Set(reflect.ValueOf(v))
	return true
}

//Returns the type of the object
func (o object) Type() reflect.Type {
	return reflect.Value(o).Type()
}

func newStruct(fields ...reflect.StructField) reflect.Type {
	return reflect.StructOf(fields)
}

func newStructField(name string, typ reflect.Type, tag string) reflect.StructField {
	return reflect.StructField{Name: name, Type: typ, Tag: reflect.StructTag(tag)}
}
