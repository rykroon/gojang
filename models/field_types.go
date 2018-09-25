package models

import (
	//"strconv"
	//"reflect"
)

type field struct { //Required Attributes

	dbColumn string
	dbType   string
	goType   string
	model    *Model

	//specific for CharField
	maxLength int

	//specfic for DecimalField
	maxDigits     int
	decimalPlaces int

	//Constraint Attributes
	null         bool
	primaryKey   bool
	unique       bool
	defaultValue string

	//Other Attributes
	autoCreated bool
	concrete    bool
	hidden      bool

	//Relation Attributes
	isRelation   bool
	manyToMany   bool
	manyToOne    bool
	oneToMany    bool
	oneToOne     bool
	relatedModel *Model

	//foreignKey bool
	onDelete onDelete
}

type onDelete string

const Cascade onDelete = "CASCADE"
const Protect onDelete = "RESTRICT"
const SetNull onDelete = "SET NULL"
const SetDefault onDelete = "SET DEFAULT"

//Constructors

func AutoField() field {
	return field{dbType: "SERIAL4", goType: "int32"}
}

func BooleanField() field {
	return field{dbType: "BOOL", goType: "bool"}
}

// func CharField(maxLength int) field {
// 	n := strconv.Itoa(maxLength)
// 	dataType := "VARCHAR(" + n + ")"
//
// 	return field{dbType: dataType, goType: "string", maxLength: maxLength}
// }

// func DecimalField(maxDigits int, decimalPlaces int) field {
// 	precision := strconv.Itoa(maxDigits)
// 	scale := strconv.Itoa(decimalPlaces)
// 	dataType := "DECIMAL(" + precision + ", " + scale + ")"
//
// 	field := field{dbType: dataType, goType: "float64"}
// 	field.maxDigits = maxDigits
// 	field.decimalPlaces = decimalPlaces
// 	return field
// }

func FloatField() field {
	return field{dbType: "FLOAT8", goType: "float64"}
}

func IntegerField() field {
	return field{dbType: "INT4", goType: "int32"}
}

func TextField() field {
	return field{dbType: "TEXT", goType: "string"}
}

//Relation Fields
func ForeignKey(to *Model, onDelete onDelete) field {
	relatedPkey := to.getPrimaryKey()
	field := field{dbType: relatedPkey.dbType, goType: relatedPkey.goType}
	field.isRelation = true
	field.relatedModel = to
	field.onDelete = onDelete
	field.manyToOne = true
	return field
}

// func OneToOneField(to *Model, onDelete onDelete) field {
// 	relatedPkey := to.getPrimaryKey()
// 	field := field{dbType: relatedPkey.dbType, goType: relatedPkey.goType}
// 	field.isRelation = true
// 	field.relatedModel = to
// 	field.onDelete = onDelete
// 	field.oneToOne = true
// 	return field
// }
//
// func ManyToManyField(to *Model, onDelete onDelete) field {
// 	relatedPkey := to.getPrimaryKey()
// 	field := field{dbType: relatedPkey.dbType, goType: relatedPkey.goType}
// 	field.isRelation = true
// 	field.relatedModel = to
// 	field.onDelete = onDelete
// 	field.manyToMany = true
// 	return field
// }

func (f field) create() string {
	s := doubleQuotes(f.dbColumn) + " " + f.dbType

	if f.primaryKey {
		s += " PRIMARY KEY"
	} else {

		//if f.foreignKey {
		if f.isRelation {
			s += " REFERENCES " + f.relatedModel.dbTable + " ON DELETE " + string(f.onDelete)
		}

		if f.null {
			s += " NULL"
		} else {
			s += " NOT NULL"
		}

		if f.unique {
			s += " UNIQUE"
		}

		if f.defaultValue != "" {
			s += " DEFAULT " + f.defaultValue
		}
	}

	return s
}
