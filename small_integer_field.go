package gojang

import ()

type SmallIntegerField struct {
	*Column

	Valid bool
	Value int16
}

func NewSmallIntegerField() *SmallIntegerField {
	field := &SmallIntegerField{}
	field.Column = newColumn("INT2")
	field.Valid = true
	return field
}

func (f *SmallIntegerField) asAssignment() assignment {
	return assignment(f.Exact(int(f.Value)))
}

func (f *SmallIntegerField) Assign(value int16) assignment {
	field := f.copy()
	field.Value = value
	return field.asAssignment()
}

func (f *SmallIntegerField) Avg() *aggregate {
	return Avg(f, NewFloatField())
}

func (f *SmallIntegerField) copy() *SmallIntegerField {
	copy := NewSmallIntegerField()
	copy.Column = f.Column.copy()
	return copy
}

func (f *SmallIntegerField) copyField() field {
	return f.copy()
}

func (f *SmallIntegerField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *SmallIntegerField) Exact(value int) lookup {
	return exactIntField(f, value)
}

func (f *SmallIntegerField) getValue() interface{} {
	return int(f.Value)
}

func (f *SmallIntegerField) Gt(value int) lookup {
	return gtIntField(f, value)
}

func (f *SmallIntegerField) Gte(value int) lookup {
	return gteIntField(f, value)
}

func (f *SmallIntegerField) In(values ...int) lookup {
	return inIntField(f, values)
}

func (f *SmallIntegerField) Lt(value int) lookup {
	return ltIntField(f, value)
}

func (f *SmallIntegerField) Lte(value int) lookup {
	return lteIntField(f, value)
}

func (f *SmallIntegerField) Max() *aggregate {
	return Max(f, NewSmallIntegerField())
}

func (f *SmallIntegerField) Min() *aggregate {
	return Min(f, NewSmallIntegerField())
}

func (f *SmallIntegerField) Range(from, to int) lookup {
	return rangeIntField(f, from, to)
}

func (f *SmallIntegerField) Scan(value interface{}) error {
	result, ok := value.(int64)
	f.Value, f.Valid = int16(result), ok
	return nil
}

func (f *SmallIntegerField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
}

func (f *SmallIntegerField) Sum() *aggregate {
	return Sum(f, NewSmallIntegerField())
}

func (f *SmallIntegerField) validate() {
	if f.primaryKey && f.null {
		panic(NewConstraintConflict(f, "primary key", "null"))
	}
}

func (f *SmallIntegerField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return intAsSql(int(f.Value))
	}
}
