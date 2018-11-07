package gojang

import (
	"fmt"
	"github.com/shopspring/decimal"
)

//A field type for each data type

type DecimalField struct {
	*Column

	Valid bool
	Value decimal.Decimal

	maxDigits     int
	decimalPlaces int
}

type AutoField struct {
	*IntegerField
}

type BigAutoField struct {
	*BigIntegerField
}

type CharField struct {
	*TextField
	maxLength int
}

type ForeignKey struct {
	*BigIntegerField

	//specific for related fields
	manyToMany   bool
	manyToOne    bool
	oneToMany    bool
	oneToOne     bool
	relatedModel *Model
	onDelete     onDelete
}

type OneToOneField struct {
	*ForeignKey
}

//Constructors

func NewCharField(maxLength int) *CharField {
	field := &CharField{maxLength: maxLength}
	field.TextField = NewTextField()
	dataType := fmt.Sprintf("VARCHAR(%v)", maxLength)
	field.dataType = dataType
	return field
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

func NewAutoField() *AutoField {
	field := &AutoField{}
	field.IntegerField = NewIntegerField()
	field.dataType = "SERIAL4"
	return field
}

func NewBigAutoField() *BigAutoField {
	field := &BigAutoField{}
	field.BigIntegerField = NewBigIntegerField()
	field.dataType = "SERIAL8"
	return field
}

func NewForeignKey(to *Model, onDelete onDelete) *ForeignKey {
	field := &ForeignKey{}
	field.BigIntegerField = NewBigIntegerField()

	field.isRelation = true
	field.manyToOne = true
	field.relatedModel = to
	field.onDelete = onDelete

	return field
}

func NewOneToOneField(to *Model, onDelete onDelete) *OneToOneField {
	field := &OneToOneField{}
	field.ForeignKey = NewForeignKey(to, onDelete)

	field.manyToOne = false
	field.oneToOne = true

	//unique constraint must be true for OneToOne Field
	field.unique = true

	return field
}

func (f *CharField) copy() *CharField {
	copy := NewCharField(f.maxLength)
	copy.Column = f.Column.copy()
	return copy
}

func (f *DecimalField) copy() *DecimalField {
	copy := NewDecimalField(f.maxDigits, f.decimalPlaces)
	copy.Column = f.Column.copy()
	return copy
}

func create(f field) string {
	s := dbq(f.ColumnName()) + " " + f.DataType()

	if f.PrimaryKey() {
		s += " PRIMARY KEY"
	} else {

		if f.HasRelation() {
			fkey := f.(relatedField)
			s += " REFERENCES " + dbq(fkey.getRelatedModel().dbTable) + " ON DELETE " + fkey.getOnDelete()
		}

		if f.Null() {
			s += " NULL"
		} else {
			s += " NOT NULL"
		}

		if f.Unique() {
			s += " UNIQUE"
		}
	}

	return s
}
