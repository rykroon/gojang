package gojang

import (
	"database/sql"
	"reflect"
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

//Returns a New Model
func NewModel(db Database) Model {
	m := Model{}
	m.Fields = make(map[string]field)
	m.db, _ = db.toDB()
	return m
}

//
func MakeModel(i interface{}) error {
	v := reflect.ValueOf(i)

	if v.Kind() != reflect.Ptr {
		panic("Value is not a pointer")
	}

	v = v.Elem()

	if v.Kind() != reflect.Struct {
		panic("Value does not point to a struct")
	}

	model := v.FieldByName("Model")

	if model.Type() != reflect.TypeOf(Model{}) {
		panic("Value does not have an embedded Model")
	}

	tableName := v.Type().String()
	dotIdx := strings.Index(tableName, ".") + 1
	tableName = strings.ToLower(tableName[dotIdx:])

	dbTable := model.FieldByName("DBTable")
	dbTable.SetString(tableName)

	fieldsMap := model.FieldByName("Fields")
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
		}
	}

	if numOfPKs < 1 {
		panic("Model must have a Primary Key")
	}

	return nil
}

//Perfect this later
// func (m *Model) Set(newModel Model) error {
// 	//check if models are of some type
// 	//are probably a few way to do this
//
// 	for key, _ := range m.Fields {
// 		m.Fields[key], _ = newModel.Fields[key]
// 	}
//
// 	return nil
// }

func (m Model) getPKField() primaryKeyField {
	for _, field := range m.Fields {
		if field.hasPrimaryKeyConstraint() {
			pk, _ := field.(primaryKeyField)
			return pk
		}
	}

	return NewAutoField()
}

func (m Model) getPointers(columns []string) []interface{} {
	result := make([]interface{}, 0)

	for _, col := range columns {
		field := m.getFieldByDbColumn(col)
		if field != nil {
			goType := field.getGoType()

			switch goType {
			case "int32":
				result = append(result, (*int32)(field.getPtr()))
			case "string":
				result = append(result, (*string)(field.getPtr()))
			}
		}
	}

	return result
}

func (m Model) getFieldByDbColumn(dbColumn string) field {
	for _, field := range m.Fields {
		if field.getDbColumn() == dbColumn {
			return field
		}
	}

	return nil
}

//If instance does not have a primary key then it will insert into the database
//Otherwise it updates the record
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
	var pkFieldName string

	for _, field := range m.Fields {
		if field.hasPrimaryKeyConstraint() {
			pkFieldName = field.getDbColumn()
			continue
		}

		columns += dbq(field.getDbColumn()) + ", "
		values += field.sqlValue() + ", "
	}

	columns = columns[:len(columns)-2] + ")"
	values = values[:len(values)-2] + ")"
	sql += columns + " VALUES " + values + " RETURNING " + dbq(pkFieldName) + ";"

	return sql
}

func (m *Model) update() string {
	sql := "UPDATE " + dbq(m.DBTable) + " SET "
	var pk field

	for _, field := range m.Fields {
		if field.hasPrimaryKeyConstraint() {
			pk = field
			continue
		}

		sql += dbq(field.getDbColumn()) + " = " + field.sqlValue() + ", "
	}

	sql = sql[:len(sql)-2]
	sql += " WHERE " + dbq(pk.getDbColumn()) + " = " + pk.sqlValue()

	return sql
}

//Creates the Database table
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
