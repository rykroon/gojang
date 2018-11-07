package gojang

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type onDelete string

const Cascade onDelete = "CASCADE"
const Protect onDelete = "RESTRICT"
const SetNull onDelete = "SET NULL"
const SetDefault onDelete = "SET DEFAULT"



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
	dbType := fmt.Sprintf("VARCHAR(%v)", maxLength)
	field.dbType = dbType
	return field
}

func NewDecimalField(maxDigits int, decimalPlaces int) *DecimalField {
	if maxDigits < decimalPlaces {
		err := NewFieldError("The maximum digits cannot be less than the number of decimal places.")
		panic(err)
	}

	field := &DecimalField{maxDigits: maxDigits, decimalPlaces: decimalPlaces}
	dbType := fmt.Sprintf("NUMERIC(%v, %v)", maxDigits, decimalPlaces)
	field.Column = newColumn(dbType)
	field.Value = decimal.New(0, 0)
	field.Valid = true
	return field
}

func NewAutoField() *AutoField {
	field := &AutoField{}
	field.IntegerField = NewIntegerField()
	field.dbType = "SERIAL4"
	return field
}

func NewBigAutoField() *BigAutoField {
	field := &BigAutoField{}
	field.BigIntegerField = NewBigIntegerField()
	field.dbType = "SERIAL8"
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
	s := dbq(f.DbColumn()) + " " + f.DbType()

	if f.HasPrimaryKeyConstraint() {
		s += " PRIMARY KEY"
	} else {

		if f.HasRelation() {
			fkey := f.(relatedField)
			s += " REFERENCES " + dbq(fkey.getRelatedModel().dbTable) + " ON DELETE " + fkey.getOnDelete()
		}

		if f.HasNullConstraint() {
			s += " NULL"
		} else {
			s += " NOT NULL"
		}

		if f.HasUniqueConstraint() {
			s += " UNIQUE"
		}
	}

	return s
}
