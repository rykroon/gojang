package gojang

import (
	"database/sql"
	"reflect"
	//"unsafe"
	"errors"
	//"fmt"
	"strings"
)

type Model struct {
	DBTable string
	Objects Manager
	Fields  map[string]field
	PK field

	db *sql.DB

	//Meta
	//uniqueTogether []string
}

func NewModel() Model {
	m := Model{}
	m.Fields = make(map[string]field)
	return m
}


func MakeModel(i interface{}) error {
	v := reflect.ValueOf(i)

	if v.Kind() != reflect.Ptr {
		panic("Value is not a pointer")
		//return errors.New("Value is not a pointer")
	}

	v = v.Elem()

	if v.Kind() != reflect.Struct {
		return errors.New("Value does not point to a struct")
	}

	// if !v.CanSet() {
	// 	return errors.New("value is not settable")
	// }

	modelVal := v.FieldByName("Model")

	if modelVal.Type() != reflect.TypeOf(Model{}) {
		panic("Value does not have an embedded Model")
	}

	tableName := v.Type().String()
	dotIdx := strings.Index(tableName, ".") + 1
	tableName = strings.ToLower(tableName[dotIdx:])

	dbTableVal := modelVal.FieldByName("DBTable")
	dbTableVal.SetString(tableName)

	//var fieldsMap reflect.Value

	fieldsMap := modelVal.FieldByName("Fields")
	numOfPKs := 0

	for idx := 0; idx < v.NumField(); idx++ {
		fieldVal := v.Field(idx)
		fieldType := v.Type().Field(idx)

		_, isAField := fieldVal.Interface().(field)

		if isAField {
			if fieldVal.Interface().(field).hasPrimaryKeyConstraint() {
				numOfPKs += 1

				if numOfPKs > 1 {
					panic("Model cannot have more than one primary key")
				}
			}

			fieldVal.Interface().(field).setDBColumn(strings.ToLower(fieldType.Name))
			fieldsMap.SetMapIndex(reflect.ValueOf(fieldType.Name), fieldVal)
			//save primaryKey to model.PK

		}
	}

	if numOfPKs < 1 {
		panic("Model must have a Primary Key")
	}

	return nil
}


func (m *Model) Save() string {
	return m.insert()
}


func (m *Model) insert() string {
	sql := "INSERT INTO " + dbq(m.DBTable) + " "

	columns := "("
	values := "("

	//var pkField field

	for _, field := range m.Fields {
		if field.hasPrimaryKeyConstraint() {
			//pkField = field
			continue
		}

		columns += dbq(field.DBColumn()) + ", "
		values += field.sqlValue() + ", "
	}

	columns = columns[:len(columns) - 2] + ")"
	values = values[:len(values) - 2] + ")"

	sql += columns + " VALUES " + values + ";"

	return sql

	//execute sql and then store lastInsertId into the primary key value
	//pkField.set(lastInsertId)
}

func (m *Model) update() {
	return

}


//returns a list of the model's fields in SQL format
// func (m Model) sqlFieldList() []string {
// 	list := []string{}
//
// 	for _, field := range m.fields {
// 		list = append(list, field.toSql())
// 	}
//
// 	return list
// }

//Returns a map with attribute names as the keys and database columns as the values
// func (m Model) attrToDBColumnMap() map[string]string {
// 	result := make(map[string]string)
//
// 	for key, field := range m.fields {
// 		result[key] = field.dbColumn
// 	}
//
// 	return result
// }

//Returns a map of database columns as the keys and attribute names as the values
// func (m Model) dbColumnToAttrMap() map[string]string {
// 	result := make(map[string]string)
//
// 	for key, field := range m.fields {
// 		result[field.dbColumn] = key
// 	}
//
// 	return result
// }

//Creates a table
func (m Model) CreateTable() string {
	sql := "CREATE TABLE IF NOT EXISTS " + m.DBTable + "("

	for _, field := range m.Fields {
		sql += create(field) + ", "
	}

	sql = sql[0:len(sql)-2] + ");"
	return sql

	//return s
	//fmt.Println(sql)
	// _, err := m.db.Exec(sql)
	//
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	//fmt.Println(result.LastInsertId())
	// 	//fmt.Println(result.RowsAffected())
	// }

}
