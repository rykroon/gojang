package gojang

import (
	"errors"
	"reflect"
	"strconv"
)

func lookupDbTableTag(tag reflect.StructTag) string {
	dbTable, _ := tag.Lookup("dbTable")
	return dbTable
}

func lookupDbColumnTag(tag reflect.StructTag) string {
	dbColumn, _ := tag.Lookup("dbColumn")
	return dbColumn
}

func lookupPrimaryKeyTag(tag reflect.StructTag) (bool, error) {
	value, err := lookupBool(tag, "primaryKey")

	if err != nil {
		return false, errors.New("gojang: Invalid value for primaryKey tag")
	}

	return value, nil
}

func lookupNullTag(tag reflect.StructTag) (bool, error) {
	value, err := lookupBool(tag, "null")

	if err != nil {
		return false, errors.New("gojang: Invalid value for null tag")
	}

	return value, nil
}

func lookupUniqueTag(tag reflect.StructTag) (bool, error) {
	value, err := lookupBool(tag, "unique")
	if err != nil {
		return false, errors.New("gojang: Invalid value for unique tag")
	}

	return value, nil
}

//Returns the value associated with the key in the struct tag.
//Returns an error if the value cannot be Parsed to a bool
func lookupBool(tag reflect.StructTag, key string) (bool, error) {
	value, ok := tag.Lookup(key)

	if ok {
		boolValue, err := strconv.ParseBool(value)

		if err != nil {
			return false, err
		}

		return boolValue, nil
	}

	return false, nil
}
