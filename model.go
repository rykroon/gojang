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
	PK      field

	db *sql.DB

	//Meta
	//uniqueTogether []string
}

func NewModel(db Database) Model {
	m := Model{}
	m.Fields = make(map[string]field)
	m.db, _ = db.toDB()
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

func (m Model) getPKField() primaryKeyField {
	for _, field := range m.Fields {
		if field.hasPrimaryKeyConstraint() {
			pk, _ := field.(primaryKeyField)
			return pk
		}
	}

	return NewAutoField()
}

// func (m Model) getPK() int {
// 	pkField := m.getPKField()
//
// 	//var i interface{} = pkField
// 	typ := reflect.TypeOf(pkField)
//
// 	if typ == reflect.TypeOf(*AutoField{}) {
// 		autoField := pkField.
// 	}
//
//
//
//
// 	return 0
//
// }

func (m *Model) Save() error {
	pk := m.getPKField()

	if pk.id() == 0 {
		var id int

		err := m.db.QueryRow(m.insert()).Scan(&id)
		if err != nil {
			return err
		}

		pk.setId(id)

	} else {
		_, err := m.db.Exec(m.update())

		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Model) insert() string {
	sql := "INSERT INTO " + dbq(m.DBTable) + " "

	columns := "("
	values := "("

	//var pkField field
	var pkFieldName string

	for _, field := range m.Fields {
		if field.hasPrimaryKeyConstraint() {
			pkFieldName = field.DBColumn()
			continue
		}

		columns += dbq(field.DBColumn()) + ", "
		values += field.sqlValue() + ", "
	}

	columns = columns[:len(columns)-2] + ")"
	values = values[:len(values)-2] + ")"

	sql += columns + " VALUES " + values + " RETURNING " + dbq(pkFieldName) + ";"

	return sql

	//execute sql and then store lastInsertId into the primary key value
	//pkField.set(lastInsertId)
}


func (m *Model) update() string {
	sql := "UPDATE " + dbq(m.DBTable) + " SET "

	var pk field

	for _, field := range m.Fields {
		if field.hasPrimaryKeyConstraint() {
			pk = field
			continue
		}

		sql += dbq(field.DBColumn()) + " = " + field.sqlValue() + ", "
	}

	sql = sql[:len(sql)-2]

	sql += " WHERE " + dbq(pk.DBColumn()) + " = " + pk.sqlValue()

	return sql
}

//Creates the table
func (m Model) Migrate() (sql.Result, error) {
	sql := m.createTable()
	result, err := m.db.Exec(sql)
	return result, err
}

//Creates an SQL statement that will create the table
func (m Model) createTable() string {
	sql := "CREATE TABLE IF NOT EXISTS " + m.DBTable + "("

	for _, field := range m.Fields {
		sql += create(field) + ", "
	}

	sql = sql[0:len(sql)-2] + ");"
	return sql
}
