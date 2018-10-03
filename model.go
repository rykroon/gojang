package gojang

import (
	"database/sql"
	"reflect"
	//"unsafe"
	"errors"
	"fmt"
	"strings"
)

type Model struct {
	DBTable string
	Objects Manager
	Fields  map[string]field

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
		return errors.New("Value is not a pointer")
	}

	v = v.Elem()

	if v.Kind() != reflect.Struct {
		return errors.New("Value is not a struct")
	}

	if !v.CanSet() {
		return errors.New("value is not settable")
	}

	tableName := v.Type().String()
	dotIndex := strings.Index(tableName, ".") + 1
	tableName = tableName[dotIndex:]

	modelValue := v.FieldByName("Model")

	dbTableValue := modelValue.FieldByName("DBTable")
	dbTableValue.SetString(tableName)

	var fieldsMap reflect.Value

	if modelValue.Type() == reflect.TypeOf(Model{}) {
		fieldsMap = modelValue.FieldByName("Fields")

	} else {
		return errors.New("Struct is not a model")
	}

	var validTypes []reflect.Type

	validTypes = append(validTypes, reflect.TypeOf(BooleanField{}))
	validTypes = append(validTypes, reflect.TypeOf(FloatField{}))
	validTypes = append(validTypes, reflect.TypeOf(IntegerField{}))
	validTypes = append(validTypes, reflect.TypeOf(TextField{}))

	for idx := 0; idx < v.NumField(); idx++ {
		fieldValue := v.Field(idx)
		fieldType := v.Type().Field(idx)

		isAField := false

		for _, validType := range validTypes {
			if fieldValue.Type() == validType {
				isAField = true
				break
			}
		}

		if isAField {
			fieldsMap.SetMapIndex(reflect.ValueOf(fieldType.Name), fieldValue.Addr())
		}
	}

	return nil
}

func (m Model) getPrimaryKeyField() field {
	for _, field := range m.Fields {
		if field.HasPrimaryKeyConstraint() {
			return field
		}
	}
	return nil
}

func (m *Model) Save() string {
	pkField := m.getPrimaryKeyField()
	pkIntField := pkField.(*IntegerField)
	fmt.Println(pkIntField)

	if pkIntField.Val() == 0 {
		return m.insert()
	}

	return ""

}

func (m *Model) insert() string {
	sql := "INSERT INTO " + dbq(m.DBTable) + " "

	columns := "("
	values := "("

	for key, field := range m.Fields {
		columns += dbq(key) + ", "
		values += field.sqlValue() + ", "
	}

	columns = columns[:len(columns) - 2] + ")"
	values = values[:len(values) - 2] + ")"

	sql += columns + " VALUES " + values + ";"

	return sql
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
func (m Model) CreateTable() {
	sql := "CREATE TABLE IF NOT EXISTS " + m.DBTable + "("

	for _, field := range m.Fields {
		sql += create(field) + ", "
	}

	sql = sql[0:len(sql)-2] + ");"

	//return s
	//fmt.Println(sql)
	_, err := m.db.Exec(sql)

	if err != nil {
		panic(err)
	} else {
		//fmt.Println(result.LastInsertId())
		//fmt.Println(result.RowsAffected())
	}

}