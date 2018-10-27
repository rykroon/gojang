package gojang

import (
//"reflect"
)

//A object has an underlying type of reflect.Value, whose underlying type is a struct
type object map[string]interface{}

func NewObj() object {
	return make(map[string]interface{})
}

//Returns a list of the objects attributes
func (o object) Dirs() []string {
	var result []string

	for key, _ := range o {
		result = append(result, key)
	}

	return result
}

//returns the attribute if it exists
func (o object) GetAttr(name string) interface{} {
	value, _ := o[name]
	return value
}

//Returns true if the object has the attribute
func (o object) HasAttr(name string) bool {
	_, ok := o[name]
	return ok
}

//Returns true if it was able to set the new value to the attribute, otherwise false
func (o object) SetAttr(name string, v interface{}) {
	o[name] = v
}
