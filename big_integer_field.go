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

// func (f *BigIntegerField) asAssignment() assignment {
// 	return assignment(f.Exact(int(f.Value)))
// }

func (f *BigIntegerField) Assign(value int64) assignment {
	return newAssignment(f, intAsSql(int(value)))
}

func (f *BigIntegerField) Avg() *aggregate {
	return Avg(f, NewFloatField())
}

func (f *BigIntegerField) copy() *BigIntegerField {
	copy := NewBigIntegerField()
	copy.Column = f.Column.copy()
	return copy
}

func (f *BigIntegerField) copyField() field {
	return f.copy()
}

func (f *BigIntegerField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *BigIntegerField) getValue() interface{} {
	return int(f.Value)
}

//
// Lookups
//

func (f *BigIntegerField) Exact(value int) lookup {
	rhs := intAsSql(value)
	return exact(f, rhs)
}

func (f *BigIntegerField) In(values ...int) lookup {
	rhs := integersAsSql(values)
	return in(f, rhs)
}

func (f *BigIntegerField) Gt(value int) lookup {
	rhs := intAsSql(value)
	return gt(f, rhs)
}

func (f *BigIntegerField) Gte(value int) lookup {
	rhs := intAsSql(value)
	return gte(f, rhs)
}

func (f *BigIntegerField) Lt(value int) lookup {
	rhs := intAsSql(value)
	return lt(f, rhs)
}

func (f *BigIntegerField) Lte(value int) lookup {
	rhs := intAsSql(value)
	return lte(f, rhs)
}

func (f *BigIntegerField) Range(from, to int) lookup {
	rhs1 := intAsSql(from)
	rhs2 := intAsSql(to)
	return between(f, rhs1, rhs2)
}

func (f *BigIntegerField) IsNull(value bool) lookup {
	return isNull(f, value)
}

func (f *BigIntegerField) Max() *aggregate {
	return Max(f, NewBigIntegerField())
}

func (f *BigIntegerField) Min() *aggregate {
	return Min(f, NewBigIntegerField())
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

func (f *BigIntegerField) Sum() *aggregate {
	return Sum(f, NewBigIntegerField())
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
