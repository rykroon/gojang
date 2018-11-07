package gojang

import (
	"database/sql"
	//"fmt"
	"reflect"
	"strings"
)

type Model struct {
	dbTable   string
	Objects   Manager
	fields    []field
	colToAttr map[string]string
	Pk        PrimaryKeyField

	db *sql.DB

	//Meta
	//uniqueTogether []string
}

//Returns a New Model
func NewModel(db Database) *Model {
	model := &Model{}
	model.fields = make([]field, 0)
	model.colToAttr = make(map[string]string)
	model.db, _ = db.toDB()
	model.Objects = NewManager(model)
	return model
}

//initializes a Model
//func MakeModel(i interface{}) error {
func MakeModel(i ModelInstance) error {
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
			model.addField(fieldType.Name, field)

			if field.HasPrimaryKeyConstraint() {
				//Use type assertion to make sure that even if the field has a
				//Primary Key Constraint that it can still implement the
				//primaryKeyField interface
				pkeyField, ok := field.(PrimaryKeyField)

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
		model.Pk.setColumnName("id")
		model.Pk.setPrimaryKeyConstraint(true)
		model.addField("id", model.Pk)
	}

	return nil
}

//Add Field to the model
func (m *Model) addField(attrName string, f field) {
	f.As(attrName)
	f.setModel(m)
	f.validate()

	columnName := f.ColumnName()

	_, duplicate := m.colToAttr[columnName]
	if duplicate {
		panic(NewDuplicateColumnError(columnName))
	}

	m.colToAttr[columnName] = attrName

	//prepend primaryKeyField to the beginning of the slice
	_, isAPrimaryKeyField := f.(PrimaryKeyField)
	if isAPrimaryKeyField && len(m.fields) != 0 {
		m.fields = append([]field{f}, m.fields...)
	} else {
		m.fields = append(m.fields, f)
	}
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
