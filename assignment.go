package gojang

import ()

//I don't think an assignment is technically a lookup, but they are similar in syntax
type assignment lookup

func (f *DecimalField) asAssignment() assignment {
	return assignment(f.Exact(f.Value))
}
