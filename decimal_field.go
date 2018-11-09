package gojang

import (
	"database/sql/driver"
	"fmt"
	"github.com/shopspring/decimal"
)

type DecimalField struct {
	*Column

	Valid bool
	Val   decimal.Decimal

	maxDigits     int
	decimalPlaces int
}

func NewDecimalField(maxDigits int, decimalPlaces int) *DecimalField {
	if maxDigits < decimalPlaces {
		err := NewFieldError("The maximum digits cannot be less than the number of decimal places.")
		panic(err)
	}

	field := &DecimalField{maxDigits: maxDigits, decimalPlaces: decimalPlaces}
	dataType := fmt.Sprintf("NUMERIC(%v, %v)", maxDigits, decimalPlaces)
	field.Column = newColumn(dataType)
	field.Val = decimal.New(0, 0)
	field.Valid = true
	return field
}

// func (f *DecimalField) asAssignment() assignment {
// 	return assignment(f.Exact(f.Val))
// }

func (f *DecimalField) Assign(value decimal.Decimal) assignment {
	return newAssignment(f, value.String())
}

func (f *DecimalField) copy() *DecimalField {
	copy := NewDecimalField(f.maxDigits, f.decimalPlaces)
	copy.Column = f.Column.copy()
	return copy
}

func (f *DecimalField) copyField() field {
	return f.copy()
}

func (f *DecimalField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
}

func (f *DecimalField) validate() {
	if f.primaryKey {
		panic(NewInvalidConstraint(f, "primary key"))
	}
}

func (f *DecimalField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return f.Val.String()
	}
}

//Aggregates

func (f *DecimalField) Avg() *aggregate {
	return Avg(f, NewDecimalField(f.maxDigits, f.decimalPlaces))
}

func (f *DecimalField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *DecimalField) Max() *aggregate {
	return Max(f, NewDecimalField(f.maxDigits, f.decimalPlaces))
}

func (f *DecimalField) Min() *aggregate {
	return Min(f, NewDecimalField(f.maxDigits, f.decimalPlaces))
}

func (f *DecimalField) Sum() *aggregate {
	return Sum(f, NewDecimalField(f.maxDigits, f.decimalPlaces))
}

//Lookups

func (f *DecimalField) Exact(value decimal.Decimal) lookup {
	rhs := value.String()
	return exact(f, rhs)
}

func (f *DecimalField) In(values ...decimal.Decimal) lookup {
	//rhs := boolsAsSql(values)
	return in(f, "do this later")
}

func (f *DecimalField) Gt(value decimal.Decimal) lookup {
	rhs := value.String()
	return gt(f, rhs)
}

func (f *DecimalField) Gte(value decimal.Decimal) lookup {
	rhs := value.String()
	return gte(f, rhs)
}

func (f *DecimalField) Lt(value decimal.Decimal) lookup {
	rhs := value.String()
	return lt(f, rhs)
}

func (f *DecimalField) Lte(value decimal.Decimal) lookup {
	rhs := value.String()
	return lte(f, rhs)
}

func (f *DecimalField) Range(from, to decimal.Decimal) lookup {
	rhs1 := from.String()
	rhs2 := to.String()
	return between(f, rhs1, rhs2)
}

func (f *DecimalField) IsNull(value bool) lookup {
	return isNull(f, value)
}

//
// Scanner Valuer
//

func (f *DecimalField) Scan(value interface{}) error {
	return f.Val.Scan(value)
}

func (f *DecimalField) Value() (driver.Value, error) {
	return f.Val.Value()
}

func (f *DecimalField) getValue() interface{} {
	return f.Val
}
