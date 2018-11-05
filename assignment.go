package gojang

import ()

//I don't think an assignment is technically a lookup, but they are similar in syntax
type assignment lookup

func (f *BigIntegerField) asAssignment() assignment {
	return assignment(f.Exact(int(f.Value)))
}

func (f *BooleanField) asAssignment() assignment {
	return assignment(f.Exact(f.Value))
}

func (f *DecimalField) asAssignment() assignment {
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
	field := f.copy()
	field.Value = value
	return field.asAssignment()
}

func (f *BooleanField) Assign(value bool) assignment {
	field := f.copy()
	field.Value = value
	return field.asAssignment()
}

func (f *FloatField) Assign(value float64) assignment {
	field := f.copy()
	field.Value = value
	return field.asAssignment()
}

func (f *IntegerField) Assign(value int32) assignment {
	field := f.copy()
	field.Value = value
	return field.asAssignment()
}

func (f *SmallIntegerField) Assign(value int16) assignment {
	field := f.copy()
	field.Value = value
	return field.asAssignment()
}

func (f *TextField) Assign(value string) assignment {
	field := f.copy()
	field.Value = value
	return field.asAssignment()
}
