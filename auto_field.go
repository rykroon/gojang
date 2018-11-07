package gojang

import ()

type AutoField struct {
	*IntegerField
}

func NewAutoField() *AutoField {
	field := &AutoField{}
	field.IntegerField = NewIntegerField()
	field.dataType = "SERIAL4"
	return field
}

func (f *AutoField) Id() int {
	return f.Val()
}

func (f *AutoField) isAutoField() bool {
	return true
}

func (f *AutoField) validate() {
	if !f.primaryKey {
		panic(NewForceConstraint(f, "primary key"))
	}

	if f.primaryKey && f.null {
		panic(NewConstraintConflict(f, "primary key", "null"))
	}
}
