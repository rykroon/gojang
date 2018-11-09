package gojang

import (
	"database/sql/driver"
)

type BooleanField struct {
	*Column

	Valid bool
	Val   bool
}

func NewBooleanField() *BooleanField {
	field := &BooleanField{}
	field.Column = newColumn("BOOL")
	field.Val = true
	return field
}

// func (f *BooleanField) asAssignment() assignment {
// 	return assignment(f.Exact(f.Val))
// }

func (f *BooleanField) Assign(value bool) assignment {
	return newAssignment(f, boolAsSql(value))
}

func (f *BooleanField) copy() *BooleanField {
	copy := NewBooleanField()
	copy.Column = f.Column.copy()
	return copy
}

func (f *BooleanField) copyField() field {
	return f.copy()
}

//
// Aggregates
//

func (f *BooleanField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *BooleanField) Max() *aggregate {
	return Max(f, NewBooleanField())
}

func (f *BooleanField) Min() *aggregate {
	return Min(f, NewBooleanField())
}

//
// Lookups
//

func (f *BooleanField) Exact(value bool) lookup {
	rhs := boolAsSql(value)
	return exact(f, rhs)
}

func (f *BooleanField) In(values ...bool) lookup {
	rhs := boolsAsSql(values)
	return in(f, rhs)
}

func (f *BooleanField) Gt(value bool) lookup {
	rhs := boolAsSql(value)
	return gt(f, rhs)
}

func (f *BooleanField) Gte(value bool) lookup {
	rhs := boolAsSql(value)
	return gte(f, rhs)
}

func (f *BooleanField) Lt(value bool) lookup {
	rhs := boolAsSql(value)
	return lt(f, rhs)
}

func (f *BooleanField) Lte(value bool) lookup {
	rhs := boolAsSql(value)
	return lte(f, rhs)
}

func (f *BooleanField) Range(from, to bool) lookup {
	rhs1 := boolAsSql(from)
	rhs2 := boolAsSql(to)
	return between(f, rhs1, rhs2)
}

func (f *BooleanField) IsNull(value bool) lookup {
	return isNull(f, value)
}

//
// Scanner, Valuer
//

func (f *BooleanField) Scan(value interface{}) error {
	f.Val, f.Valid = value.(bool)
	return nil
}

func (f *BooleanField) xValue() (driver.Value, error) {
	return f.Val, nil
}

func (f *BooleanField) getValue() interface{} {
	return f.Val
}

func (f *BooleanField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
}

func (f *BooleanField) validate() {
	if f.primaryKey {
		panic(NewInvalidConstraint(f, "primary key"))
	}
}

func (f *BooleanField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return boolAsSql(f.Val)
	}
}
