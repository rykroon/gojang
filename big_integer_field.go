package gojang

import (
	"database/sql/driver"
)

type BigIntegerField struct {
	*Column

	Valid bool
	Val   int64
}

func NewBigIntegerField() *BigIntegerField {
	field := &BigIntegerField{}
	field.Column = newColumn("INT8")
	field.Valid = true
	return field
}

// func (f *BigIntegerField) asAssignment() assignment {
// 	return assignment(f.Exact(int(f.Val)))
// }

func (f *BigIntegerField) Assign(value int64) assignment {
	return newAssignment(f, intAsSql(int(value)))
}

func (f *BigIntegerField) asSqlValue() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return intAsSql(int(f.Val))
	}
}

func (f *BigIntegerField) copy() *BigIntegerField {
	copy := NewBigIntegerField()
	*copy = *f
	return copy
}

func (f *BigIntegerField) copyField() field {
	return f.copy()
}

func (f *BigIntegerField) Set(value int64) columnAssigner {
	copy := f.copy()
	copy.Val = value
	return copy
}

//
// Aggregates
//

func (f *BigIntegerField) Avg() *aggregate {
	return Avg(f, NewFloatField())
}

func (f *BigIntegerField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *BigIntegerField) Max() *aggregate {
	return Max(f, NewBigIntegerField())
}

func (f *BigIntegerField) Min() *aggregate {
	return Min(f, NewBigIntegerField())
}

func (f *BigIntegerField) Sum() *aggregate {
	return Sum(f, NewBigIntegerField())
}

//
// Lookups
//

func (f *BigIntegerField) Exact(value int) lookup {
	// field := f.copy()
	// field.Val = int64(value)

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

//
// Scanner, Valuer
//

func (f *BigIntegerField) Scan(value interface{}) error {
	f.Val, f.Valid = value.(int64)
	return nil
}

func (f *BigIntegerField) Value() (driver.Value, error) {
	return f.Val, nil
}

func (f *BigIntegerField) getValue() interface{} {
	return int(f.Val)
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
