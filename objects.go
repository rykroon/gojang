package gojang

import (
//"reflect"
)

//A Object has an underlying type of reflect.Value, whose underlying type is a struct
type Object map[string]interface{}

func newObj() Object {
	return make(map[string]interface{})
}

//Returns a list of the Objects attributes
func (o Object) Dirs() []string {
	var result []string

	for key, _ := range o {
		result = append(result, key)
	}

	return result
}

//returns the attribute if it exists
func (o Object) GetAttr(name string) interface{} {
	value, _ := o[name]
	return value
}

//Returns true if the Object has the attribute
func (o Object) HasAttr(name string) bool {
	_, ok := o[name]
	return ok
}

//Returns true if it was able to set the new value to the attribute, otherwise false
func (o Object) SetAttr(name string, v interface{}) {
	o[name] = v
}
