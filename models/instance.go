package models

import (
	//"reflect"
)

type modelInstance struct {
	model  *Model
	values map[string]fieldInstance
}

type fieldInstance struct {
	dbColumn string
	//dbType string
	goType string
	primaryKey bool
	value interface{} //?
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
	i.values = make(map[string]fieldInstance)

	for key, field := range m.fields {
		field := fieldInstance{dbColumn:field.dbColumn, goType:field.goType, primaryKey: field.primaryKey}
		i.values[key] = field
	}

	return i
}

func (i *modelInstance) Get(attr string) (interface{}, bool) {
	field, ok := i.values[attr]

	if ok {
		return field.value, ok
	}

	return 0, ok
}

func (i *modelInstance) Set(attr string, value interface{}) bool {
	field, ok := i.values[attr]

	if ok {
		field.value = value
		i.values[attr] = field
	}

	return ok
}


func (i modelInstance) Save() {

}

func (i modelInstance) insert() {

}

func (i modelInstance) update() {

}
