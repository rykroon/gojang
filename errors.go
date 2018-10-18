package gojang

import (
//"errors"
)

type ObjectDoesNotExist struct {
	message string
}

type MultipleObjectsReturned struct {
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
