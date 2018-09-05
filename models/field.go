package models

import (
	"strconv"
)

//maybe?
type field interface {
}

type Field struct {
	dbColumn   string
	dbDataType string

	//specific for CharField
	maxLength int

	//specfic for DecimalField
	maxDigits     int
	decimalPlaces int

	//constraints
	null       bool
	primaryKey bool
	unique     bool

	foreignKey bool
	to         *Model
	onDelete   string
}

type fieldOption func(*Field)


// func newField(options ...func(*Field) error) (*Field, error) {
// 	f := &Field{}
//
// 	for _, option := range options{
//     err := option(f)
//     if err != nil {
//       return nil, err
//     }
//   }
//
//   return f, nil
// }
//
// func PrimaryKey(pkey bool) fieldOption {
// 	return func(f *Field) {
// 		f.primaryKey=pkey
// 	}
// }
//
// func Unique(unique bool) fieldOption {
// 	return func(f *Field) {
// 		f.unique=unique
// 	}
// }
//
// func Null(null bool) fieldOption {
// 	return func(f *Field) {
// 		f.null=null
// 	}
// }


func AutoField() Field {
	return Field{dbDataType: "SERIAL"}
}

func BooleanField() Field {
	return Field{dbDataType: "BOOLEAN"}
}

func CharField(maxLength int) Field {
	n := strconv.Itoa(maxLength)
	dataType := "VARCHAR(" + n + ")"

	return Field{dbDataType: dataType, maxLength: maxLength}
}

func DecimalField(maxDigits int, decimalPlaces int) Field {
	precision := strconv.Itoa(maxDigits)
	scale := strconv.Itoa(decimalPlaces)
	dataType := "NUMERIC(" + precision + ", " + scale + ")"

	return Field{dbDataType: dataType, maxDigits: maxDigits, decimalPlaces: decimalPlaces}
}

func FloatField() Field {
	return Field{dbDataType: "DOUBLE PRECISION"}
}

func ForeignKey(m *Model, onDelete string) Field {
	return Field{dbDataType: "INTEGER", foreignKey: true, to: m, onDelete: onDelete}
}

func IntegerField() Field {
	return Field{dbDataType: "INTEGER"}
}

// func IntegerField(options ...fieldOption) Field {
// 	f := &Field{}
//
// 	for _, option := range options{
//     option(f)
//   }
//
// 	f.dbDataType = "INTEGER"
// 	return *f
// }

func TextField() Field {
	return Field{dbDataType: "TEXT"}
}

//Constraints
func (f Field) PrimaryKey() Field {
	if !f.foreignKey {
		f.primaryKey = true
	}

	return f
}

func (f Field) Null() Field {
	f.null = true
	return f
}

func (f Field) Unique() Field {
	f.unique = true
	return f
}

//maybe make each Field their own type so I can specify the appropriate types as parameters
func (f Field) Default() Field {
	return f
}




func (f Field) createString(dbColumn string) string {
	s := dbColumn + " " + f.dbDataType

	if f.primaryKey {
		s += " PRIMARY KEY"
	} else {

		if f.foreignKey {
			s += " REFERENCES " + f.to.dbTable + " ON DELETE " + f.onDelete
		}

		if f.null {
			s += " NULL"
		} else {
			s += " NOT NULL"
		}

		if f.unique {
			s += " UNIQUE"
		}
	}

	return s
}
