package gojang

import (
	"database/sql/driver"
)

type IntegerField struct {
	*Column

	Valid bool
	Val   int32
}

func NewIntegerField() *IntegerField {
	field := &IntegerField{}
	field.Column = newColumn("INT4")
	field.Valid = true
	return field
}

// func (f *IntegerField) asAssignment() assignment {
// 	return assignment(f.Exact(int(f.Val)))
// }

func (f *IntegerField) Assign(value int32) assignment {
	return newAssignment(f, intAsSql(int(value)))
}

func (f *IntegerField) copy() *IntegerField {
	copy := NewIntegerField()
	copy.Column = f.Column.copy()
	return copy
}

func (f *IntegerField) copyField() field {
	return f.copy()
}

//
// Aggregates
//

func (f *IntegerField) Avg() *aggregate {
	return Avg(f, NewFloatField())
}

func (f *IntegerField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *IntegerField) Max() *aggregate {
	return Max(f, NewIntegerField())
}

func (f *IntegerField) Min() *aggregate {
	return Min(f, NewIntegerField())
}

func (f *IntegerField) Sum() *aggregate {
	return Sum(f, NewIntegerField())
}

//
// Lookups
//

func (f *IntegerField) Exact(value int) lookup {
	rhs := intAsSql(value)
	return exact(f, rhs)
}

func (f *IntegerField) In(values ...int) lookup {
	rhs := integersAsSql(values)
	return in(f, rhs)
}

func (f *IntegerField) Gt(value int) lookup {
	rhs := intAsSql(value)
	return gt(f, rhs)
}

func (f *IntegerField) Gte(value int) lookup {
	rhs := intAsSql(value)
	return gte(f, rhs)
}

func (f *IntegerField) Lt(value int) lookup {
	rhs := intAsSql(value)
	return lt(f, rhs)
}

func (f *IntegerField) Lte(value int) lookup {
	rhs := intAsSql(value)
	return lte(f, rhs)
}

func (f *IntegerField) Range(from, to int) lookup {
	rhs1 := intAsSql(from)
	rhs2 := intAsSql(to)
	return between(f, rhs1, rhs2)
}

func (f *IntegerField) IsNull(value bool) lookup {
	return isNull(f, value)
}

//
// Scanner, Valuer
//

func (f *IntegerField) Scan(value interface{}) error {
	result, ok := value.(int64)
	f.Val, f.Valid = int32(result), ok
	return nil
}

func (f *IntegerField) xValue() (driver.Value, error) {
	return int64(f.Val), nil
}

func (f *IntegerField) getValue() interface{} {
	return int(f.Val)
}

func (f *IntegerField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
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
		return intAsSql(int(f.Val))
	}
}
