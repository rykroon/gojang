package gojang

import ()

//I don't think an assignment is technically a lookup, but they are similar in syntax
type assignment lookup

func (a assignment) asSql() string {
	return dbq(a.lhs.getDbColumn()) + " " + a.lookupName + " " + a.rhs
}

func (f *BigIntegerField) asAssignment() assignment {
	return assignment(f.Exact(int(f.Value)))
}

func (f *BooleanField) asAssignment() assignment {
	return assignment(f.Exact(f.Value))
}

func (f *FloatField) asAssignment() assignment {
	return assignment(f.Exact(f.Value))
}

func (f *IntegerField) asAssignment() assignment {
	return assignment(f.Exact(int(f.Value)))
}

func (f *SmallIntegerField) asAssignment() assignment {
	return assignment(f.Exact(int(f.Value)))
}

func (f *TextField) asAssignment() assignment {
	return assignment(f.Exact(f.Value))
}

func (f *BigIntegerField) Assign(value int64) assignment {
	f.Value = value
	f.valid = true
	return f.asAssignment()
}

func (f *BooleanField) Assign(value bool) assignment {
	f.Value = value
	f.valid = true
	return f.asAssignment()
}

func (f *FloatField) Assign(value float64) assignment {
	f.Value = value
	f.valid = true
	return f.asAssignment()
}

func (f *IntegerField) Assign(value int32) assignment {
	f.Value = value
	f.valid = true
	return f.asAssignment()
}

func (f *SmallIntegerField) Assign(value int16) assignment {
	f.Value = value
	f.valid = true
	return f.asAssignment()
}

func (f *TextField) Assign(value string) assignment {
	f.Value = value
	f.valid = true
	return f.asAssignment()
}
