package gojang

import (
	"fmt"
)

//Expressions describe a value or a computation that can be used as part of an
//update, create, filter, order by, annotation, or aggregate.
type expression interface {
	asSql() string
}

type selectExpression interface {
	expression
	getValue() interface{}
	As(string) //is essentially the 'setter' method for alias
	Alias() string
	DataType() string
	Scan(interface{}) error
}

type orderByExpression string

type star string

func (f *function) asSql() string {
	return fmt.Sprintf(f.template, f.args...)
}

func (s star) asSql() string {
	return "*"
}

//
//Select Expression Method Set
//

func (s star) Alias() string {
	return "*"
}

func (s star) As(string) {
	return
}

func (s star) DataType() string {
	return ""
}

func (s star) Scan(interface{}) error {
	return nil
}

func (s star) getValue() interface{} {
	return nil
}
