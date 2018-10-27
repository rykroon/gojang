package gojang

import (
	"fmt"
)

type ObjectDoesNotExist struct {
	message string
}

type MultipleObjectsReturned struct {
	message string
}

type FieldError struct {
	message string
}

type ModelError struct {
	message string
}

func NewObjectDoesNotExist() ObjectDoesNotExist {
	msg := "gojang: The requested object does not exist"
	err := ObjectDoesNotExist{message: msg}
	return err
}

func (e ObjectDoesNotExist) Error() string {
	return e.message
}

func NewMultipleObjectsReturned() MultipleObjectsReturned {
	msg := "gojang: The query returned multiple objects when only one was expected."
	err := MultipleObjectsReturned{message: msg}
	return err
}

func (e MultipleObjectsReturned) Error() string {
	return e.message
}

func NewFieldError(msg string) FieldError {
	err := FieldError{}
	err.message = "gojang: " + msg
	return err
}

func (e FieldError) Error() string {
	return e.message
}

func NewConstraintConflict(fld field, const1, const2 string) FieldError {
	msg := fmt.Sprintf("%T cannot have both %v and %v constraints.", fld, const1, const2)
	return NewFieldError(msg)
}

func NewForceConstraint(fld field, constraint string) FieldError {
	msg := fmt.Sprintf("%T must have %v constraint.", fld, constraint)
	return NewFieldError(msg)
}

func NewInvalidConstraint(fld field, constraint string) FieldError {
	msg := fmt.Sprintf("%T cannot have %v constraint.", fld, constraint)
	return NewFieldError(msg)
}

func NewConstraintViolation(value, constraint string) FieldError {
	msg := fmt.Sprintf("Value %v violates %v constraint.", value, constraint)
	return NewFieldError(msg)
}

func NewNotNullConstraintViolation() FieldError {
	return NewConstraintViolation("null", "not-null")
}

func NewModelError(msg string) ModelError {
	err := ModelError{}
	err.message = "gojang: " + msg
	return err
}

func (e ModelError) Error() string {
	return e.message
}

func NewMultiplePrimaryKeyError(model Model) ModelError {
	msg := fmt.Sprintf("Multiple primary keys for model %v are not allowed.", model.dbTable)
	return NewModelError(msg)
}

func NewDuplicateColumnError(column string) ModelError {
	msg := fmt.Sprintf("Column \"%v\" specified more than once.", column)
	return NewModelError(msg)
}
