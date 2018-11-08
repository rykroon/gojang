package gojang

import (
	"fmt"
)

//I don't think an assignment is technically a lookup, but they are similar in syntax
//type assignment lookup
type assignment struct {
	lhs columner
	rhs string
}

func newAssignment(lhs columner, rhs string) assignment {
	return assignment{lhs: lhs, rhs: rhs}
}

func (a assignment) asSql() string {
	return fmt.Sprintf("%v = %v", a.lhs.asSql(), a.rhs)
}
