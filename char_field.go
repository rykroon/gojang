package gojang

import (
	"fmt"
)

type CharField struct {
	*TextField
	maxLength int
}

func NewCharField(maxLength int) *CharField {
	field := &CharField{maxLength: maxLength}
	field.TextField = NewTextField()
	dataType := fmt.Sprintf("VARCHAR(%v)", maxLength)
	field.dataType = dataType
	return field
}

func (f *CharField) copy() *CharField {
	copy := NewCharField(f.maxLength)
	copy.Column = f.Column.copy()
	return copy
}

func (f *CharField) copyField() field {
	return f.copy()
}
