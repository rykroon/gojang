package gojang

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type DecimalField struct {
	*Column

	Valid bool
	Value decimal.Decimal

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
	field.Value = decimal.New(0, 0)
	field.Valid = true
	return field
}

func (f *DecimalField) asAssignment() assignment {
	return assignment(f.Exact(f.Value))
}

func (f *DecimalField) copy() *DecimalField {
	copy := NewDecimalField(f.maxDigits, f.decimalPlaces)
	copy.Column = f.Column.copy()
	return copy
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
	return lookup{lhs: f, lookupName: "=", rhs: f.Value.String()}
}

func (f *DecimalField) Gt(value decimal.Decimal) lookup {
	return lookup{lhs: f, lookupName: ">", rhs: f.Value.String()}
}

func (f *DecimalField) Gte(value decimal.Decimal) lookup {
	return lookup{lhs: f, lookupName: "<=", rhs: f.Value.String()}
}

func (f *DecimalField) Lt(value decimal.Decimal) lookup {
	return lookup{lhs: f, lookupName: "<", rhs: f.Value.String()}
}

func (f *DecimalField) Lte(value decimal.Decimal) lookup {
	return lookup{lhs: f, lookupName: "<=", rhs: f.Value.String()}
}

func (f *DecimalField) Range(from, to decimal.Decimal) lookup {
	lookup := lookup{lhs: f, lookupName: "BETWEEN"}
	lookup.rhs = from.String() + " AND " + to.String()
	return lookup
}

func (f *DecimalField) IsNull(value bool) lookup {
	return fieldIsNull(f, value)
}

func (f *DecimalField) Scan(value interface{}) error {
	return f.Value.Scan(value)
}

func (f *DecimalField) getValue() interface{} {
	return f.Value
}
