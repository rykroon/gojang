package models

import (
	"reflect"
	"strconv"
)

//auxillary functions used in all programs

func doubleQuotes(s string) string {
	return "\"" + s + "\""
}

func singleQuotes(s string) string {
	return "'" + s + "'"
}

func valueToSql(i interface{}) string {
	v := reflect.ValueOf(i)
	k := v.Kind()

	switch k {
	case reflect.Int:
		return strconv.Itoa(i.(int))

	case reflect.String:
		return singleQuotes(i.(string))

	case reflect.Float64:
		return strconv.FormatFloat(i.(float64), 'E', -1, 64)

	case reflect.Bool:
		if i.(bool) {
			return "TRUE"
		} else {
			return "FALSE"
		}
	}

	return ""
}

//Constants

const Cascade string = "CASCADE"
const Protect string = "RESTRICT"
const SetNull string = "SET NULL"
const SetDefault string = "SET DEFAULT"
