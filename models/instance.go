package models

import (
//"reflect"
)

type modelInstance struct {
	model  *Model
	values map[string]instanceField
	pkAttr string
}

type instanceField struct {
	dbColumn string
	//dbType string
	goType     string
	primaryKey bool
	value      interface{} //?
}

// type fieldInstance struct {
// 	dataType reflect.Kind
//
// 	intValue    int
// 	floatValue  float64
// 	stringValue string
// 	boolValue   bool
// }

func (i modelInstance) buildQuery() string {
	return ""
}

func (m *Model) NewInstance() modelInstance {
	i := modelInstance{}
	i.model = m
	i.values = make(map[string]instanceField)

	for key, field := range m.fields {
		field := instanceField{dbColumn: field.dbColumn, goType: field.goType, primaryKey: field.primaryKey}
		i.values[key] = field

		if field.primaryKey {
			i.pkAttr = key
		}
	}

	return i
}

//Get the value of an attribute
func (i *modelInstance) Get(attr string) (interface{}, bool) {
	field, ok := i.values[attr]

	if ok {
		return field.value, ok
	}

	return 0, ok
}


//Get Value of PrimaryKey
func (i *modelInstance) GetPK() (interface{}, bool) {
	value, ok := i.Get(i.pkAttr)

	if ok {
		return value, ok
	}

	return 0, ok
}

//Set a value of an attribute
func (i *modelInstance) Set(attr string, value interface{}) bool {
	field, ok := i.values[attr]

	if ok {
		field.value = value
		i.values[attr] = field
	}

	return ok
}

func (i *modelInstance) SetPK(value interface{}) bool {
	return i.Set(i.pkAttr, value)
}

func (i *modelInstance) Save() {
	//check if primary key is 0 value
	
}

func (i *modelInstance) insert() {

}

func (i *modelInstance) update() {

}
