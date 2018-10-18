package gojang

import (
	//"strconv"
	//"reflect"
	"errors"
	//"fmt"
)

type constraint string

const Null constraint = "NULL"
const Unique constraint = "UNIQUE"

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
}

//Constructors

func NewAutoField(constraints ...constraint) *AutoField {
	field := &AutoField{dbType: "SERIAL4"}
	field.pointer = &field.value
	field.primaryKey = true
	return field
}

func NewBigAutoField() *BigAutoField {
	field := &BigAutoField{dbType: "SERIAL8"}
	field.pointer = &field.value
	field.primaryKey = true
	return field
}

func NewBigIntegerField(constraints ...constraint) *BigIntegerField {
	field := &BigIntegerField{dbType: "INT8"}

	for _, constraint := range constraints {
		switch constraint {
		case "NULL":
			field.null = true

		case "UNIQUE":
			field.unique = true
		}
	}

	if !field.null {
		field.pointer = &field.value
	}

	return field
}

func NewBooleanField(constraints ...constraint) *BooleanField {
	field := &BooleanField{dbType: "BOOL"}

	for _, constraint := range constraints {
		switch constraint {
		case "NULL":
			field.null = true

		case "UNIQUE":
			field.unique = true
		}
	}

	if !field.null {
		field.pointer = &field.value
	}

	return field
}

func NewFloatField(constraints ...constraint) *FloatField {
	field := &FloatField{dbType: "FLOAT8"}

	for _, constraint := range constraints {
		switch constraint {
		case "NULL":
			field.null = true

		case "UNIQUE":
			field.unique = true
		}
	}

	if !field.null {
		field.pointer = &field.value
	}

	return field
}

func NewIntegerField(constraints ...constraint) *IntegerField {
	field := &IntegerField{dbType: "INT4"}

	for _, constraint := range constraints {
		switch constraint {
		case "NULL":
			field.null = true

		case "UNIQUE":
			field.unique = true
		}
	}

	if !field.null {
		field.pointer = &field.value
	}

	return field
}

func NewTextField(constraints ...constraint) *TextField {
	field := &TextField{dbType: "TEXT"}

	for _, constraint := range constraints {
		switch constraint {
		case "NULL":
			field.null = true

		case "UNIQUE":
			field.unique = true
		}
	}

	if !field.null {
		field.pointer = &field.value
	}

	return field
}

func NewForeignKey(to *Model, onDelete onDelete, constraints ...constraint) *ForeignKey {
	field := &ForeignKey{dbType: "INT8"}
	field.isRelation = true
	field.manyToOne = true
	field.relatedModel = to
	field.onDelete = onDelete

	for _, constraint := range constraints {
		switch constraint {
		case "NULL":
			field.null = true

		case "UNIQUE":
			field.unique = true
		}
	}

	if !field.null {
		field.pointer = &field.value
	}

	return field
}

func NewOneToOneField(to *Model, onDelete onDelete, constraints ...constraint) *OneToOneField {
	field := &OneToOneField{dbType: "INT8"}
	field.isRelation = true
	field.oneToOne = true
	field.relatedModel = to
	field.onDelete = onDelete

	for _, constraint := range constraints {
		switch constraint {
		case "NULL":
			field.null = true

		case "UNIQUE":
			field.unique = true
		}
	}

	//unique constraint must be true foe OneToOne Field
	field.unique = true

	if !field.null {
		field.pointer = &field.value
	}

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
		return errors.New("Cannot set field with NOT NULL constraint to nil")
	}
}

func (f *BooleanField) SetNil() error {
	if f.hasNullConstraint() {
		f.pointer = nil
		f.value = false
		return nil
	} else {
		return errors.New("Cannot set field with NOT NULL constraint to nil")
	}
}

func (f *FloatField) SetNil() error {
	if f.hasNullConstraint() {
		f.pointer = nil
		f.value = 0
		return nil
	} else {
		return errors.New("Cannot set field with NOT NULL constraint to nil")
	}
}

func (f *IntegerField) SetNil() error {
	if f.hasNullConstraint() {
		f.pointer = nil
		f.value = 0
		return nil
	} else {
		return errors.New("Cannot set field with NOT NULL constraint to nil")
	}
}

func (f *TextField) SetNil() error {
	if f.hasNullConstraint() {
		f.pointer = nil
		f.value = ""
		return nil
	} else {
		return errors.New("Cannot set field with NOT NULL constraint to nil")
	}
}

func (f *ForeignKey) SetNil() error {
	if f.hasNullConstraint() {
		f.pointer = nil
		f.value = 0
		return nil
	} else {
		return errors.New("Cannot set field with NOT NULL constraint to nil")
	}
}

func (f *OneToOneField) SetNil() error {
	if f.hasNullConstraint() {
		f.pointer = nil
		f.value = 0
		return nil
	} else {
		return errors.New("Cannot set field with NOT NULL constraint to nil")
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
