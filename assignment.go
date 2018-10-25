package gojang

import ()

type assignment struct {
	lhs field
	rhs string
	//rhs expression
}

func (a assignment) asSql() string {
	return dbq(a.lhs.getDbColumn()) + " = " + a.rhs //a.rhs.asSql()
}

func (f *BigIntegerField) Assign(value int64) assignment {
	return assignment{lhs: f, rhs: int64ToSql(value)}
}
