package gojang

import (
//"strconv"
//"reflect"
//"fmt"
)

type onDelete string

const Cascade onDelete = "CASCADE"
const Protect onDelete = "RESTRICT"
const SetNull onDelete = "SET NULL"
const SetDefault onDelete = "SET DEFAULT"

type AutoField struct {
	model    *Model
	dbColumn string
	dbType   string

	null       bool
	unique     bool
	primaryKey bool
	isRelation bool

	expr expression

	valid bool
	Value int32
}

type BigAutoField struct {
	model    *Model
	dbColumn string
	dbType   string

	null       bool
	unique     bool
	primaryKey bool
	isRelation bool

	expr expression

	valid bool
	Value int64
}

type BigIntegerField struct {
	model    *Model
	dbColumn string
	dbType   string

	null       bool
	unique     bool
	primaryKey bool
	isRelation bool

	expr expression

	valid bool
	Value int64
}

type BooleanField struct {
	model    *Model
	dbColumn string
	dbType   string

	null       bool
	unique     bool
	primaryKey bool
	isRelation bool

	expr expression

	valid bool
	Value bool
}

type FloatField struct {
	model    *Model
	dbColumn string
	dbType   string

	null       bool
	unique     bool
	primaryKey bool
	isRelation bool

	expr expression

	valid bool
	Value float64
}

type IntegerField struct {
	model    *Model
	dbColumn string
	dbType   string

	null       bool
	unique     bool
	primaryKey bool
	isRelation bool

	expr expression

	valid bool
	Value int32
}

type SmallIntegerField struct {
	model    *Model
	dbColumn string
	dbType   string

	null       bool
	unique     bool
	primaryKey bool
	isRelation bool

	expr expression

	valid bool
	Value int16
}

type TextField struct {
	model    *Model
	dbColumn string
	dbType   string

	null       bool
	unique     bool
	primaryKey bool
	isRelation bool

	expr expression

	valid bool
	Value string
}

type ForeignKey struct {
	model    *Model
	dbColumn string
	dbType   string

	null       bool
	unique     bool
	primaryKey bool
	isRelation bool

	expr expression

	//specific for related fields
	manyToMany   bool
	manyToOne    bool
	oneToMany    bool
	oneToOne     bool
	relatedModel *Model
	onDelete     onDelete

	valid bool
	Value int64
}

type OneToOneField struct {
	model    *Model
	dbColumn string
	dbType   string

	null       bool
	unique     bool
	primaryKey bool
	isRelation bool

	expr expression

	//specific for related fields
	manyToMany   bool
	manyToOne    bool
	oneToMany    bool
	oneToOne     bool
	relatedModel *Model
	onDelete     onDelete

	valid bool
	Value int64
}

//Constructors

func NewAutoField() *AutoField {
	field := &AutoField{dbType: "SERIAL4"}
	field.valid = true
	return field
}

func NewBigAutoField() *BigAutoField {
	field := &BigAutoField{dbType: "SERIAL8"}
	field.valid = true
	return field
}

func NewBigIntegerField() *BigIntegerField {
	field := &BigIntegerField{dbType: "INT8"}
	field.valid = true
	return field
}

func NewBooleanField() *BooleanField {
	field := &BooleanField{dbType: "BOOL"}
	field.valid = true
	return field
}

func NewFloatField() *FloatField {
	field := &FloatField{dbType: "FLOAT8"}
	field.valid = true
	return field
}

func NewIntegerField() *IntegerField {
	field := &IntegerField{dbType: "INT4"}
	field.valid = true
	return field
}

func NewSmallIntegerField() *SmallIntegerField {
	field := &SmallIntegerField{dbType: "INT2"}
	field.valid = true
	return field
}

func NewTextField() *TextField {
	field := &TextField{dbType: "TEXT"}
	field.valid = true
	return field
}

func NewForeignKey(to *Model, onDelete onDelete) *ForeignKey {
	field := &ForeignKey{dbType: "INT8"}
	field.valid = true

	field.isRelation = true
	field.manyToOne = true
	field.relatedModel = to
	field.onDelete = onDelete

	return field
}

func NewOneToOneField(to *Model, onDelete onDelete) *OneToOneField {
	field := &OneToOneField{dbType: "INT8"}
	field.valid = true

	field.isRelation = true
	field.oneToOne = true
	field.relatedModel = to
	field.onDelete = onDelete

	//unique constraint must be true for OneToOne Field
	field.unique = true

	return field
}

func (f *AutoField) SetNil() error {
	return NewNotNullConstraintViolation()
}

func (f *BigAutoField) SetNil() error {
	return NewNotNullConstraintViolation()
}

func (f *BigIntegerField) SetNil() error {
	if f.hasNullConstraint() {
		f.valid = false
		f.Value = 0
		return nil
	} else {
		return NewNotNullConstraintViolation()
	}
}

func (f *BooleanField) SetNil() error {
	if f.hasNullConstraint() {
		f.valid = false
		f.Value = false
		return nil
	} else {
		return NewNotNullConstraintViolation()
	}
}

func (f *FloatField) SetNil() error {
	if f.hasNullConstraint() {
		f.valid = false
		f.Value = 0
		return nil
	} else {
		return NewNotNullConstraintViolation()
	}
}

func (f *IntegerField) SetNil() error {
	if f.hasNullConstraint() {
		f.valid = false
		f.Value = 0
		return nil
	} else {
		return NewNotNullConstraintViolation()
	}
}

func (f *SmallIntegerField) SetNil() error {
	if f.hasNullConstraint() {
		f.valid = false
		f.Value = 0
		return nil
	} else {
		return NewNotNullConstraintViolation()
	}
}

func (f *TextField) SetNil() error {
	if f.hasNullConstraint() {
		f.valid = false
		f.Value = ""
		return nil
	} else {
		return NewNotNullConstraintViolation()
	}
}

func (f *ForeignKey) SetNil() error {
	if f.hasNullConstraint() {
		f.valid = false
		f.Value = 0
		return nil
	} else {
		return NewNotNullConstraintViolation()
	}
}

func (f *OneToOneField) SetNil() error {
	if f.hasNullConstraint() {
		f.valid = false
		f.Value = 0
		return nil
	} else {
		return NewNotNullConstraintViolation()
	}
}

func (f AutoField) UnSetNil() {
	f.valid = true
}

func (f BigAutoField) UnSetNil() {
	f.valid = true
}

func (f BigIntegerField) UnSetNil() {
	f.valid = true
}

func (f BooleanField) UnSetNil() {
	f.valid = true
}

func (f FloatField) UnSetNil() {
	f.valid = true
}

func (f IntegerField) UnSetNil() {
	f.valid = true
}

func (f SmallIntegerField) UnSetNil() {
	f.valid = true
}

func (f TextField) UnSetNil() {
	f.valid = true
}

func (f ForeignKey) UnSetNil() {
	f.valid = true
}

func (f OneToOneField) UnSetNil() {
	f.valid = true
}

func create(f field) string {
	s := dbq(f.getDbColumn()) + " " + f.getDbType()

	if f.hasPrimaryKeyConstraint() {
		s += " PRIMARY KEY"
	} else {

		if f.hasRelation() {
			fkey := f.(relatedField)
			s += " REFERENCES " + dbq(fkey.getRelatedModel().dbTable) + " ON DELETE " + fkey.getOnDelete()
		}

		if f.hasNullConstraint() {
			s += " NULL"
		} else {
			s += " NOT NULL"
		}

		if f.hasUniqueConstraint() {
			s += " UNIQUE"
		}
	}

	return s
}
