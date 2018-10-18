package gojang

import (
	"database/sql"
	//"fmt"
	"reflect"
	"strings"
)

type Model struct {
	dbTable string
	Objects Manager
	fields  map[string]field
	//colsToFields map[string]string
	Pk      primaryKeyField

	db *sql.DB

	//Meta
	//uniqueTogether []string
}

//Returns a New Model
func NewModel(db Database) *Model {
	model := &Model{}
	model.fields = make(map[string]field)
	model.db, _ = db.toDB()
	model.Objects = newManager(model)
	return model
}

//initializes a Model
func MakeModel(i interface{}) error {
	v := reflect.ValueOf(i)

	if v.Kind() != reflect.Ptr {
		panic("Value is not a pointer")
	}

	v = v.Elem()

	if v.Kind() != reflect.Struct {
		panic("Value does not point to a struct")
	}

	modelVal := v.FieldByName("Model")

	if modelVal.Type() != reflect.TypeOf(&Model{}) {
		panic("Value does not have an embedded Model")
	}

	model := modelVal.Interface().(*Model)

	tableName := v.Type().String()
	dotIdx := strings.Index(tableName, ".") + 1
	tableName = tableName[dotIdx:]
	model.dbTable = snakeCase(tableName)

	numOfPKs := 0

	for idx := 0; idx < v.NumField(); idx++ {
		fieldVal := v.Field(idx)
		fieldType := v.Type().Field(idx)

		field, isAField := fieldVal.Interface().(field)

		if isAField {
			if field.hasPrimaryKeyConstraint() {
				numOfPKs += 1

				if numOfPKs > 1 {
					panic("Model cannot have more than one primary key")
				}

				model.Pk = field.(primaryKeyField)
			}

			field.setDbColumn(snakeCase(fieldType.Name))
			field.setModel(model)
			model.fields[fieldType.Name] = field
		}
	}

	if numOfPKs < 1 {
		model.Pk = NewAutoField()
		model.Pk.(field).setDbColumn("id")
		model.fields["id"] = model.Pk.(field)
	}

	return nil
}

//Sets a Model's fields from Rows.Scan()
//Assumes that Rows.Next() had been previously called
func (m *Model) setFromRows(rows *sql.Rows) error {
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	pointers := m.getPointers(columns)

	err = rows.Scan(pointers...)
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) getPointers(columns []string) []interface{} {
	result := make([]interface{}, 0)

	for _, col := range columns {
		field := m.getFieldByDbColumn(col)
		if field != nil {
			goType := field.getGoType()
			var ptr interface{}

			switch goType {
			case "int64":
				ptr = (*int64)(field.getPtr())
			case "int32":
				ptr = (*int32)(field.getPtr())
			case "int16":
				ptr = (*int16)(field.getPtr())
			case "float64":
				ptr = (*float64)(field.getPtr())
			case "bool":
				ptr = (*bool)(field.getPtr())
			case "string":
				ptr = (*string)(field.getPtr())
			}

			result = append(result, ptr)
		}
	}

	return result
}

func (m Model) getFieldByDbColumn(dbColumn string) field {
	for _, field := range m.fields {
		if field.getDbColumn() == dbColumn {
			return field
		}
	}

	return nil
}

//If instance does not have a primary key then it will insert into the database
//Otherwise it updates the record
func (m *Model) Save() error {
	if m.Pk.Val() == 0 {
		var err error
		row := m.db.QueryRow(m.insert())
		goType := m.Pk.(field).getGoType()

		switch goType {
		case "int64":
			ptr := (*int64)(m.Pk.(field).getPtr())
			err = row.Scan(ptr)
		case "int32":
			ptr := (*int32)(m.Pk.(field).getPtr())
			err = row.Scan(ptr)
		}

		if err != nil {
			return err
		}

	} else {
		_, err := m.db.Exec(m.update())

		if err != nil {
			return err
		}
	}

	return nil
}

//
func (m *Model) insert() string {
	sql := "INSERT INTO " + dbq(m.dbTable) + " "
	columns := "("
	values := "("
	var pkFieldName string

	for _, field := range m.fields {
		if field.hasPrimaryKeyConstraint() {
			pkFieldName = field.getDbColumn()
			continue
		}

		columns += dbq(field.getDbColumn()) + ", "
		values += field.valueToSql() + ", "
	}

	columns = columns[:len(columns)-2] + ")"
	values = values[:len(values)-2] + ")"
	sql += columns + " VALUES " + values + " RETURNING " + dbq(pkFieldName) + ";"

	return sql
}

//
func (m *Model) update() string {
	sql := "UPDATE " + dbq(m.dbTable) + " SET "
	var pk field

	for _, field := range m.fields {
		if field.hasPrimaryKeyConstraint() {
			pk = field
			continue
		}

		sql += dbq(field.getDbColumn()) + " = " + field.valueToSql() + ", "
	}

	sql = sql[:len(sql)-2]
	sql += " WHERE " + dbq(pk.getDbColumn()) + " = " + pk.valueToSql()

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
	sql := "CREATE TABLE IF NOT EXISTS " + dbq(m.dbTable) + " ("

	for _, field := range m.fields {
		sql += create(field) + ", "
	}

	sql = sql[0:len(sql)-2] + ");"
	return sql
}
