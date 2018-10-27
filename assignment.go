package gojang

import ()

// type assignment struct {
// 	lhs field
// 	rhs string
// }

//I don't think an assignment is technically a lookup, but they are similar in syntax
type assignment lookup

func (a assignment) asSql() string {
	return lookup(a).asSql()
}

func (f *BigIntegerField) Assign(value int64) assignment {
	f.Value = value
	return assignment(f.Exact(int(value)))
}

func (f *BooleanField) Assign(value bool) assignment {
	f.Value = value
	return assignment(f.Exact(value))
}

func (f *IntegerField) Assign(value int32) assignment {
	f.Value = value
	return assignment(f.Exact(int(value)))
}

func (f *SmallIntegerField) Assign(value int16) assignment {
	f.Value = value
	return assignment(f.Exact(int(value)))
}

func (f *TextField) Assign(value string) assignment {
	f.Value = value
	return assignment(f.Exact(value))
}
