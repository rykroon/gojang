package gojang

import ()

type BigIntegerField struct {
	*Column

	Valid bool
	Value int64
}

func NewBigIntegerField() *BigIntegerField {
	field := &BigIntegerField{}
	field.Column = newColumn("INT8")
	field.Valid = true
	return field
}

func (f *BigIntegerField) copy() *BigIntegerField {
	copy := NewBigIntegerField()
	copy.Column = f.Column.copy()
	return copy
}

func (f *BigIntegerField) copyField() field {
	return f.copy()
}

func (f *BigIntegerField) Exact(value int) lookup {
	return exactIntField(f, value)
}

func (f *BigIntegerField) getValue() interface{} {
	return int(f.Value)
}

func (f *BigIntegerField) Gt(value int) lookup {
	return gtIntField(f, value)
}

func (f *BigIntegerField) Gte(value int) lookup {
	return gteIntField(f, value)
	//return lookup{lhs: f, lookupName: ">=", rhs: intAsSql(value)}
}

func (f *BigIntegerField) In(values ...int) lookup {
	return inIntField(f, values)
}

func (f *BigIntegerField) IsNull(value bool) lookup {
	return fieldIsNull(f, value)
}

func (f *BigIntegerField) Lt(value int) lookup {
	return ltIntField(f, value)
	//return lookup{lhs: f, lookupName: "<", rhs: intAsSql(value)}
}

func (f *BigIntegerField) Lte(value int) lookup {
	return lteIntField(f, value)
	//return lookup{lhs: f, lookupName: "<=", rhs: intAsSql(value)}
}

func (f *BigIntegerField) Range(from, to int) lookup {
	return rangeIntField(f, from, to)
}

func (f *BigIntegerField) Scan(value interface{}) error {
	f.Value, f.Valid = value.(int64)
	return nil
}

func (f *BigIntegerField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
}

func (f *BigIntegerField) validate() {
	if f.primaryKey && f.null {
		panic(NewConstraintConflict(f, "primary key", "null"))
	}
}

func (f *BigIntegerField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return intAsSql(int(f.Value))
	}
}
