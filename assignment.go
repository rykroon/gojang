package gojang

import (
	"fmt"
)

//I don't think an assignment is technically a lookup, but they are similar in syntax
//type assignment lookup
type assignment struct {
	lhs field
	rhs string
}

type columnAssigner interface {
	ColumnName() string
	//asSqlValue() string
	getValue() interface{}
}

func newAssignment(lhs field, rhs string) assignment {
	return assignment{lhs: lhs, rhs: rhs}
}

func (a assignment) asSql() string {
	return fmt.Sprintf("%v = %v", dbq(a.lhs.ColumnName()), a.rhs)
}
