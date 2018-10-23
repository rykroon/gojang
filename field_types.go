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

	pointer *int32
	Ptr     *int32
	value   int32

	Valid bool
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

	pointer *int64
	value   int64

	Valid bool
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

	pointer *int64
	value   int64

	Valid bool
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

	pointer *bool
	value   bool

	Valid bool
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

	pointer *float64
	value   float64

	Valid bool
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

	pointer *int32
	value   int32

	Valid bool
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

	pointer *int16
	value   int16

	Valid bool
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

	pointer *string
	Ptr     *string
	value   string

	Valid bool
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

	//specific for related fields
	manyToMany   bool
	manyToOne    bool
	oneToMany    bool
	oneToOne     bool
	relatedModel *Model
	onDelete     onDelete

	pointer *int64
	value   int64

	Valid bool
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

	//specific for related fields
	manyToMany   bool
	manyToOne    bool
	oneToMany    bool
	oneToOne     bool
	relatedModel *Model
	onDelete     onDelete

	pointer *int64
	value   int64

	Valid bool
	Value int64
}

//Constructors

func NewAutoField() *AutoField {
	field := &AutoField{dbType: "SERIAL4"}
	field.pointer = &field.value
	return field
}

func NewBigAutoField() *BigAutoField {
	field := &BigAutoField{dbType: "SERIAL8"}
	field.pointer = &field.value
	return field
}

func NewBigIntegerField() *BigIntegerField {
	field := &BigIntegerField{dbType: "INT8"}
	field.pointer = &field.value
	return field
}

func NewBooleanField() *BooleanField {
	field := &BooleanField{dbType: "BOOL"}
	field.pointer = &field.value
	return field
}

func NewFloatField() *FloatField {
	field := &FloatField{dbType: "FLOAT8"}
	field.pointer = &field.value
	return field
}

func NewIntegerField() *IntegerField {
	field := &IntegerField{dbType: "INT4"}
	field.pointer = &field.value
	return field
}

func NewSmallIntegerField() *SmallIntegerField {
	field := &SmallIntegerField{dbType: "INT2"}
	field.pointer = &field.value
	return field
}

func NewTextField() *TextField {
	field := &TextField{dbType: "TEXT"}
	field.pointer = &field.value
	return field
}

func NewForeignKey(to *Model, onDelete onDelete) *ForeignKey {
	field := &ForeignKey{dbType: "INT8"}
	field.pointer = &field.value

	field.isRelation = true
	field.manyToOne = true
	field.relatedModel = to
	field.onDelete = onDelete

	return field
}

func NewOneToOneField(to *Model, onDelete onDelete) *OneToOneField {
	field := &OneToOneField{dbType: "INT8"}
	field.pointer = &field.value

	field.isRelation = true
	field.oneToOne = true
	field.relatedModel = to
	field.onDelete = onDelete

	//unique constraint must be true for OneToOne Field
	field.unique = true

	return field
}

func (f AutoField) Val() int {
	return int(f.value)
}

func (f BigAutoField) Val() int {
	return int(f.value)
}

func (f BigIntegerField) Val() int {
	return int(f.value)
}

func (f BooleanField) Val() bool {
	return f.value
}

func (f FloatField) Val() float64 {
	return f.value
}

func (f IntegerField) Val() int {
	return int(f.value)
}

func (f SmallIntegerField) Val() int {
	return int(f.value)
}

func (f TextField) Val() string {
	return f.value
}

func (f ForeignKey) Val() int {
	return int(f.value)
}

func (f OneToOneField) Val() int {
	return int(f.value)
}

func (f *BigIntegerField) Set(value int64) {
	if f.pointer == nil {
		f.pointer = &f.value
	}

	f.value = value
}

func (f *BooleanField) Set(value bool) {
	if f.pointer == nil {
		f.pointer = &f.value
	}

	f.value = value
}

func (f *FloatField) Set(value float64) {
	if f.pointer == nil {
		f.pointer = &f.value
	}

	f.value = value
}

func (f *IntegerField) Set(value int32) {
	if f.pointer == nil {
		f.pointer = &f.value
	}

	f.value = value
}

func (f *SmallIntegerField) Set(value int16) {
	if f.pointer == nil {
		f.pointer = &f.value
	}

	f.value = value
}

func (f *TextField) Set(value string) {
	if f.pointer == nil {
		f.pointer = &f.value
	}

	f.value = value
}

func (f *ForeignKey) Set(value int64) {
	if f.pointer == nil {
		f.pointer = &f.value
	}

	f.value = value
}

func (f *OneToOneField) Set(value int64) {
	if f.pointer == nil {
		f.pointer = &f.value
	}

	f.value = value
}

func (f *BigIntegerField) SetNil() error {
	if f.hasNullConstraint() {
		f.pointer = nil
		f.value = 0
		return nil
	} else {
		return NewCannotSetNil()
	}
}

func (f *BooleanField) SetNil() error {
	if f.hasNullConstraint() {
		f.pointer = nil
		f.value = false
		return nil
	} else {
		return NewCannotSetNil()
	}
}

func (f *FloatField) SetNil() error {
	if f.hasNullConstraint() {
		f.pointer = nil
		f.value = 0
		return nil
	} else {
		return NewCannotSetNil()
	}
}

func (f *IntegerField) SetNil() error {
	if f.hasNullConstraint() {
		f.pointer = nil
		f.value = 0
		return nil
	} else {
		return NewCannotSetNil()
	}
}

func (f *SmallIntegerField) SetNil() error {
	if f.hasNullConstraint() {
		f.pointer = nil
		f.value = 0
		return nil
	} else {
		return NewCannotSetNil()
	}
}

func (f *TextField) SetNil() error {
	if f.hasNullConstraint() {
		f.pointer = nil
		f.value = ""
		return nil
	} else {
		return NewCannotSetNil()
	}
}

func (f *ForeignKey) SetNil() error {
	if f.hasNullConstraint() {
		f.pointer = nil
		f.value = 0
		return nil
	} else {
		return NewCannotSetNil()
	}
}

func (f *OneToOneField) SetNil() error {
	if f.hasNullConstraint() {
		f.pointer = nil
		f.value = 0
		return nil
	} else {
		return NewCannotSetNil()
	}
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
