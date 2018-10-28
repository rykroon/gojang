package gojang

import (
	"database/sql"
	//"fmt"
	"reflect"
	//"strconv"
	"errors"
	"strings"
)

type Model struct {
	dbTable   string
	Objects   Manager
	fields    []field
	colToAttr map[string]string
	Pk        primaryKeyField

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
			model.addField(fieldType.Name, field)

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
		model.addField("id", model.Pk)
	}

	return nil
}

//Add Field to the model
func (m *Model) addField(attrName string, f field) {
	f.setModel(m)
	f.validate()

	columnName := f.getDbColumn()

	_, duplicate := m.colToAttr[columnName]
	if duplicate {
		panic(NewDuplicateColumnError(columnName))
	}

	m.colToAttr[columnName] = attrName

	//prepend primaryKeyField to the beginning of the slice
	_, isAPrimaryKeyField := f.(primaryKeyField)
	if isAPrimaryKeyField && len(m.fields) != 0 {
		m.fields = append([]field{f}, m.fields...)
	} else {
		m.fields = append(m.fields, f)
	}
}

func (m *Model) ToObj() object {
	obj := newObj()

	for _, field := range m.fields {
		attrName := m.colToAttr[field.getDbColumn()]
		obj.SetAttr(attrName, field.getValue())
	}

	return obj
}

//If instance does not have a primary key then it will insert into the database
//Otherwise it updates the record
func (m *Model) Save() error {
	if m.Pk.Id() == 0 {
		id, err := m.insert()
		if err != nil {
			return err
		}

		m.Pk.setInt(id)

	} else {
		err := m.update()

		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Model) insert() (int, error) {
	var createList []assignment

	for _, field := range m.fields {
		if !field.hasPrimaryKeyConstraint() {
			createList = append(createList, field.asAssignment())
		}
	}

	obj, err := m.Objects.Create(createList...)
	if err != nil {
		return 0, err
	}

	pkeyAttr := m.colToAttr[m.Pk.getDbColumn()]
	if obj.HasAttr(pkeyAttr) {
		return obj.GetAttr(pkeyAttr).(int), nil
	}

	return 0, errors.New("gojang: idk, we lost the key, sorry")
}

//
func (m *Model) update() error {
	var updateList []assignment

	for _, field := range m.fields {
		if !field.hasPrimaryKeyConstraint() {
			updateList = append(updateList, field.asAssignment())
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
