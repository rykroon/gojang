package gojang

import ()

type IntegerField struct {
	*Column

	Valid bool
	Value int32
}

func NewIntegerField() *IntegerField {
	field := &IntegerField{}
	field.Column = newColumn("INT4")
	field.Valid = true
	return field
}

func (f *IntegerField) asAssignment() assignment {
	return assignment(f.Exact(int(f.Value)))
}

func (f *IntegerField) Avg() *aggregate {
	return Avg(f, NewFloatField())
}

func (f *IntegerField) copy() *IntegerField {
	copy := NewIntegerField()
	copy.Column = f.Column.copy()
	return copy
}

func (f *IntegerField) copyField() field {
	return f.copy()
}

func (f *IntegerField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *IntegerField) Exact(value int) lookup {
	return exactIntField(f, value)
}

func (f *IntegerField) getValue() interface{} {
	return int(f.Value)
}

func (f *IntegerField) Gt(value int) lookup {
	return gtIntField(f, value)
}

func (f *IntegerField) Gte(value int) lookup {
	return gteIntField(f, value)
}

func (f *IntegerField) In(values ...int) lookup {
	return inIntField(f, values)
}

func (f *IntegerField) IsNull(value bool) lookup {
	return fieldIsNull(f, value)
}

func (f *IntegerField) Lt(value int) lookup {
	return ltIntField(f, value)
}

func (f *IntegerField) Lte(value int) lookup {
	return lteIntField(f, value)
}

func (f *IntegerField) Max() *aggregate {
	return Max(f, NewIntegerField())
}

func (f *IntegerField) Min() *aggregate {
	return Min(f, NewIntegerField())
}

func (f *IntegerField) Range(from, to int) lookup {
	return rangeIntField(f, from, to)
}

func (f *IntegerField) Scan(value interface{}) error {
	result, ok := value.(int64)
	f.Value, f.Valid = int32(result), ok
	return nil
}

func (f *IntegerField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
}

func (f *IntegerField) Sum() *aggregate {
	return Sum(f, NewIntegerField())
}

func (f *IntegerField) validate() {
	if f.primaryKey && f.null {
		panic(NewConstraintConflict(f, "primary key", "null"))
	}
}

func (f *IntegerField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return intAsSql(int(f.Value))
	}
}
