package gojang

import (
//"strconv"
//"reflect"
	"errors"
)

// type field struct { //Required Attributes
//
// 	dbColumn string
// 	dbType   string
// 	goType   string
// 	model    *Model
//
// 	//specific for CharField
// 	maxLength int
//
// 	//specfic for DecimalField
// 	maxDigits     int
// 	decimalPlaces int
//
// 	//Constraint Attributes
// 	null         bool
// 	primaryKey   bool
// 	unique       bool
// 	defaultValue string
//
// 	//Other Attributes
// 	autoCreated bool
// 	concrete    bool
// 	hidden      bool
//
// 	//Relation Attributes
// 	isRelation   bool
// 	manyToMany   bool
// 	manyToOne    bool
// 	oneToMany    bool
// 	oneToOne     bool
// 	relatedModel *Model
//
// 	//foreignKey bool
// 	onDelete onDelete
// }

type constraint string

const Null constraint = "NULL"
const Unique constraint = "UNIQUE"
const PrimaryKey constraint = "PRIMARY KEY"

type onDelete string

const Cascade onDelete = "CASCADE"
const Protect onDelete = "RESTRICT"
const SetNull onDelete = "SET NULL"
const SetDefault onDelete = "SET DEFAULT"


type BigIntegerField struct {
	dbColumn string
	dbType   string

	null       bool
	unique     bool
	primaryKey bool

	pointer *int64
	value int64
}

type BooleanField struct {
	dbColumn string
	dbType   string

	null       bool
	unique     bool
	primaryKey bool

	pointer *bool
	value bool
}

type FloatField struct {
	dbColumn string
	dbType   string

	null       bool
	unique     bool
	primaryKey bool

	pointer *float64
	value float64
}

type IntegerField struct {
	dbColumn string
	dbType   string

	null       bool
	unique     bool
	primaryKey bool

	pointer *int32
	value int32
}

type TextField struct {
	dbColumn string
	dbType   string

	null       bool
	unique     bool
	primaryKey bool

	pointer *string
	value string
}

//Constructors

// func AutoField() field {
// 	return field{dbType: "SERIAL4", goType: "int32"}
// }

func NewBooleanField(constraints ...constraint) BooleanField {
	field := BooleanField{dbType: "BOOL"}

	for _, constraint := range constraints {
		switch constraint {
		case "NULL":
			field.null = true

		case "UNIQUE":
			field.unique = true

		case "PRIMARY KEY":
			field.primaryKey = true
		}
	}

	if !field.null {
		field.pointer = &field.value
	}

	return field
}

func NewFloatField(constraints ...constraint) FloatField {
	field := FloatField{dbType: "FLOAT8"}

	for _, constraint := range constraints {
		switch constraint {
		case "NULL":
			field.null = true

		case "UNIQUE":
			field.unique = true

		case "PRIMARY KEY":
			field.primaryKey = true
		}
	}

	if !field.null {
		field.pointer = &field.value
	}

	return field
}

func NewIntegerField(constraints ...constraint) IntegerField {
	field := IntegerField{dbType: "INT4"}

	for _, constraint := range constraints {
		switch constraint {
		case "NULL":
			field.null = true

		case "UNIQUE":
			field.unique = true

		case "PRIMARY KEY":
			field.primaryKey = true
		}
	}

	if !field.null {
		field.pointer = &field.value
	}

	return field
}

func NewTextField(constraints ...constraint) TextField {
	field := TextField{dbType: "TEXT"}

	for _, constraint := range constraints {
		switch constraint {
		case "NULL":
			field.null = true

		case "UNIQUE":
			field.unique = true

		case "PRIMARY KEY":
			field.primaryKey = true
		}
	}

	if !field.null {
		field.pointer = &field.value
	}

	return field
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


func (f *BooleanField) SetNil() error {
	if f.HasNullConstraint() {
		f.pointer = nil
		f.value = false
		return nil
	} else {
		return errors.New("Cannot set field with NOT NULL constraint to nil")
	}
}

func (f *FloatField) SetNil() error {
	if f.HasNullConstraint() {
		f.pointer = nil
		f.value = 0
		return nil
	} else {
		return errors.New("Cannot set field with NOT NULL constraint to nil")
	}
}

func (f *IntegerField) SetNil() error {
	if f.HasNullConstraint() {
		f.pointer = nil
		f.value = 0
		return nil
	} else {
		return errors.New("Cannot set field with NOT NULL constraint to nil")
	}
}

func (f *TextField) SetNil() error {
	if f.HasNullConstraint() {
		f.pointer = nil
		f.value = ""
		return nil
	} else {
		return errors.New("Cannot set field with NOT NULL constraint to nil")
	}
}



//Relation Fields
// func ForeignKey(to *Model, onDelete onDelete) field {
// 	relatedPkey := to.getPrimaryKey()
// 	field := field{dbType: relatedPkey.dbType, goType: relatedPkey.goType}
// 	field.isRelation = true
// 	field.relatedModel = to
// 	field.onDelete = onDelete
// 	field.manyToOne = true
// 	return field
// }



func create(f field) string {
	s := doubleQuotes(f.getDBColumn()) + " " + f.getDBType()

	if f.HasPrimaryKeyConstraint() {
		s += " PRIMARY KEY"
	} else {

		//if f.foreignKey {
		// if f.isRelation {
		// 	s += " REFERENCES " + f.relatedModel.dbTable + " ON DELETE " + string(f.onDelete)
		// }

		if f.HasNullConstraint() {
			s += " NULL"
		} else {
			s += " NOT NULL"
		}

		if f.HasUniqueConstraint() {
			s += " UNIQUE"
		}

		// if f.defaultValue != "" {
		// 	s += " DEFAULT " + f.defaultValue
		// }
	}

	return s
}