package gojang

import (
  "strconv"
)

type FloatField struct {
	*Column

	Valid bool
	Value float64
}

func NewFloatField() *FloatField {
	field := &FloatField{}
	field.Column = newColumn("FLOAT8")
	field.Valid = true
	return field
}

func (f *FloatField) asAssignment() assignment {
	return assignment(f.Exact(f.Value))
}

func (f *FloatField) Assign(value float64) assignment {
	field := f.copy()
	field.Value = value
	return field.asAssignment()
}

func (f *FloatField) Avg() *aggregate {
	return Avg(f, NewFloatField())
}

func (f *FloatField) copy() *FloatField {
	copy := NewFloatField()
	copy.Column = f.Column.copy()
	return copy
}

func (f *FloatField) copyField() field {
	return f.copy()
}

func (f *FloatField) Count(distinct bool) *aggregate {
	return Count(f, distinct)
}

func (f *FloatField) Exact(value float64) lookup {
	return lookup{lhs: f, lookupName: "=", rhs: float64AsSql(value)}
}

func (f *FloatField) getValue() interface{} {
	return f.Value
}

func (f *FloatField) Gt(value float64) lookup {
	return lookup{lhs: f, lookupName: ">", rhs: float64AsSql(value)}
}

func (f *FloatField) Gte(value float64) lookup {
	return lookup{lhs: f, lookupName: ">=", rhs: float64AsSql(value)}
}

func (f *FloatField) In(values ...float64) lookup {
	return lookup{lhs: f, lookupName: "IN", rhs: floatsAsSql(values)}
}

func (f *FloatField) IsNull(value bool) lookup {
	return fieldIsNull(f, value)
}

func (f *FloatField) Lt(value float64) lookup {
	return lookup{lhs: f, lookupName: "<", rhs: float64AsSql(value)}
}

func (f *FloatField) Lte(value float64) lookup {
	return lookup{lhs: f, lookupName: "<=", rhs: float64AsSql(value)}
}

func (f *FloatField) Max() *aggregate {
	return Max(f, NewFloatField())
}

func (f *FloatField) Min() *aggregate {
	return Min(f, NewFloatField())
}

func (f *FloatField) Range(from, to float64) lookup {
	lookup := lookup{lhs: f, lookupName: "BETWEEN"}
	lookup.rhs = float64AsSql(from) + " AND " + float64AsSql(to)
	return lookup
}

func (f *FloatField) Scan(value interface{}) error {
	switch v := value.(type) {
	case float64:
		f.Value, f.Valid = v, true

	case int64:
		f.Value, f.Valid = float64(v), true

	case []uint8:
		float, err := strconv.ParseFloat(string(v), 64)
		f.Value = float
		f.Valid = err == nil
		return err

	default:
		f.Value, f.Valid = 0, false
	}

	return nil
}

func (f *FloatField) setNullConstraint(null bool) {
	f.null = null

	if f.null {
		f.Valid = false
	}
}

func (f *FloatField) Sum() *aggregate {
	return Sum(f, NewFloatField())
}

func (f *FloatField) validate() {
	if f.primaryKey {
		panic(NewInvalidConstraint(f, "primary key"))
	}
}

func (f *FloatField) valueAsSql() string {
	if f.null && !f.Valid {
		return "NULL"
	} else {
		return float64AsSql(f.Value)
	}
}
