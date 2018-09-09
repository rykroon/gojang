package models

import (
	"reflect"
)

type Instance struct {
	model  *Model
	values map[string]FieldInstance
}

type FieldInstance struct {
	dataType reflect.Kind

	intValue    int
	floatValue  float64
	stringValue string
	boolValue   bool
}

func (i Instance) buildQuery() string {
	return ""
}

func newInstance(m *Model) Instance {
	i := Instance{}
	i.model = m
	return i
}

//Maybe
// func (i Instance) SetInt(fieldName string, number int ) {
// 	field,_ := i.model.Field(fieldName)
// 	if field.defaultType != reflect.Int {
// 		//thats a problem
// 	}
// }

func (i Instance) Save() {

}

func (i Instance) insert() {

}

func (i Instance) update() {

}
