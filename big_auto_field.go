package gojang

import ()

type BigAutoField struct {
	*BigIntegerField
}

func NewBigAutoField() *BigAutoField {
	field := &BigAutoField{}
	field.BigIntegerField = NewBigIntegerField()
	field.dataType = "SERIAL8"
	return field
}

func (f *BigAutoField) Id() int {
	return f.Val()
}

func (f *BigAutoField) isAutoField() bool {
	return true
}

func (f *BigAutoField) validate() {
	if !f.primaryKey {
		panic(NewForceConstraint(f, "primary key"))
	}

	if f.primaryKey && f.null {
		panic(NewConstraintConflict(f, "primary key", "null"))
	}
}
