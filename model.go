package gojang

import (
	"database/sql"
	//"fmt"
	"reflect"
	//"strconv"
	"strings"
)

type Model struct {
	dbTable     string
	Objects     Manager
	fields      []field
	colToFields map[string]field
	Pk          primaryKeyField

	db *sql.DB

	//Meta
	//uniqueTogether []string
}

//Returns a New Model
func NewModel(db Database) *Model {
	model := &Model{}
	model.fields = make([]field, 0)
	model.colToFields = make(map[string]field)
	model.db, _ = db.toDB()
	model.Objects = newManager(model)
	return model
}

//initializes a Model
func MakeModel(i interface{}) error {
	v := reflect.ValueOf(i)

	if v.Kind() != reflect.Ptr {
		panic("gojang: Value is not a pointer")
	}

	v = v.Elem()

	if v.Kind() != reflect.Struct {
		panic("gojang: Value does not point to a struct")
	}

	modelVal := v.FieldByName("Model")
	model, ok := modelVal.Interface().(*Model)

	if !ok {
		panic("gojang: Value does not have an embedded Model")
	}

	//Get table name
	modelStruct, _ := v.Type().FieldByName("Model")
	model.dbTable = lookupDbTableTag(modelStruct.Tag)

	if model.dbTable == "" {
		tableName := strings.ToLower(v.Type().String())
		model.dbTable = strings.Replace(tableName, ".", "_", -1)
	}

	numOfPKs := 0

	for idx := 0; idx < v.NumField(); idx++ {
		fieldVal := v.Field(idx)
		fieldType := v.Type().Field(idx)

		field, isAField := fieldVal.Interface().(field)

		if isAField {

			options, err := getFieldOptions(fieldType)
			if err != nil {
				panic(err)
			}

			setFieldOptions(field, options)
			field.setModel(model)
			field.validate()
			model.addField(field)

			if field.hasPrimaryKeyConstraint() {
				//Use type assertion to make sure that even if the field has a
				//Primary Key Constraint that it can still implement the
				//primaryKeyField interface
				pkeyField, ok := field.(primaryKeyField)

				if !ok {
					panic(NewInvalidConstraint(field, "primary key"))
				}

				numOfPKs += 1

				if numOfPKs > 1 {
					panic(NewMultiplePrimaryKeyError(*model))
				}

				model.Pk = pkeyField
			}
		}
	}

	if numOfPKs < 1 {
		model.Pk = NewAutoField()
		model.Pk.setDbColumn("id")
		model.Pk.setPrimaryKeyConstraint(true)
		model.Pk.setModel(model)
		model.Pk.validate()
		model.addField(model.Pk)
	}

	return nil
}

//Add Field to the model
func (m *Model) addField(f field) {
	columnName := f.getDbColumn()

	_, duplicate := m.colToFields[columnName]
	if duplicate {
		panic(NewDuplicateColumnError(columnName))
	}

	m.colToFields[columnName] = f

	//prepend primaryKeyField to the beginning of the slice
	_, isAPrimaryKeyField := f.(primaryKeyField)
	if isAPrimaryKeyField && len(m.fields) != 0 {
		m.fields = append([]field{f}, m.fields...)
	} else {
		m.fields = append(m.fields, f)
	}
}

//If instance does not have a primary key then it will insert into the database
//Otherwise it updates the record
func (m *Model) Save() error {
	if m.Pk.Id() == 0 {
		row := m.db.QueryRow(m.insert())

		result := 0
		err := row.Scan(&result)
		if err != nil {
			return err
		}

		m.Pk.setInt(result)

	} else {
		err := m.update()

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
func (m *Model) update() error {
	var updateList []field
	for _, field := range m.fields {
		if !field.hasPrimaryKeyConstraint() {
			updateList = append(updateList, field)
		}
	}

	qs := m.Objects.Filter(m.Pk.Exact(m.Pk.Id()))
	_, err := qs.Update(updateList...)
	return err
}

//Creates the Database table
func (m Model) Migrate() error {
	return m.CreateTable()
}

//Creates an SQL statement that will create the table
func (m Model) CreateTable() error {
	sql := "CREATE TABLE IF NOT EXISTS " + dbq(m.dbTable) + " ("

	for _, field := range m.fields {
		sql += create(field) + ", "
	}

	sql = sql[0:len(sql)-2] + ");"
	_, err := m.db.Exec(sql)
	return err
}

func (m Model) DropTable() error {
	sql := "DROP TABLE " + dbq(m.dbTable) + ";"
	_, err := m.db.Exec(sql)
	return err
}
