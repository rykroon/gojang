package gojang

import (
	//"errors"
	"reflect"
)

type ObjectDoesNotExist struct {
	message string
}

type MultipleObjectsReturned struct {
	message string
}

type CannotSetNil struct {
	message string
}

type NullPrimaryKeyErr struct {
	message string
}

type InvalidPrimaryKey struct {
	message string
}

type ForcePrimaryKey struct {
	message string
}

func NewObjectDoesNotExist() ObjectDoesNotExist {
	msg := "The requested object does not exist"
	e := ObjectDoesNotExist{message: msg}
	return e
}

func (e ObjectDoesNotExist) Error() string {
	return e.message
}

func NewMultipleObjectsReturned() MultipleObjectsReturned {
	msg := "The query returned multiple objects when only one was expected."
	e := MultipleObjectsReturned{message: msg}
	return e
}

func (e MultipleObjectsReturned) Error() string {
	return e.message
}

func NewCannotSetNil() CannotSetNil {
	msg := "Cannot set field with Not-Null Constraint to nil"
	e := CannotSetNil{message: msg}
	return e
}

func (e CannotSetNil) Error() string {
	return e.message
}

func NewNullPrimaryKeyErr() NullPrimaryKeyErr {
	msg := "Field cannot have both Primary Key and Null Constraints"
	e := NullPrimaryKeyErr{message: msg}
	return e
}

func (e NullPrimaryKeyErr) Error() string {
	return e.message
}

func NewInvalidPrimaryKey(field field) InvalidPrimaryKey {
	fieldType := reflect.TypeOf(field).String()
	msg := fieldType + " cannot have Primary Key Constraint"
	e := InvalidPrimaryKey{message: msg}
	return e
}

func (e InvalidPrimaryKey) Error() string {
	return e.message
}

func NewForcePrimaryKey(field field) ForcePrimaryKey {
	fieldType := reflect.TypeOf(field).String()
	msg := fieldType + " must have Primary Key Constraint"
	e := ForcePrimaryKey{message: msg}
	return e
}

func (e ForcePrimaryKey) Error() string {
	return e.message
}
