package gojang

import (
	//"database/sql"
	"reflect"
	"strconv"
	"strings"
)

//auxillary functions used in all programs

type value reflect.Value

// type value struct {
// 	outputField field
// }

func newValue(v interface{}) value {
	return value(reflect.ValueOf(v))
}

func (v value) asSql() string {
	return ""
}

func doubleQuotes(s string) string {
	return "\"" + s + "\""
}

func dbq(s string) string {
	return doubleQuotes(s)
}

func singleQuotes(s string) string {
	return "'" + s + "'"
}

//Transforms a 'CamelCase' string into a 'snake_case' string
func snakeCase(s string) string {
	result := ""

	for idx, byte := range s {
		char := string(byte)
		lowerChar := strings.ToLower(char)

		if char != lowerChar && idx != 0 {
			result += "_" + lowerChar
		} else {
			result += lowerChar
		}
	}

	return result
}

func stringSliceToSql(slice []string) string {
	s := "("

	for _, value := range slice {
		s += stringToSql(value) + ", "
	}

	s = s[0:len(s)-2] + ")"

	return s
}

func stringToSql(s string) string {
	return singleQuotes(s)
}

func intSliceToSql(slice []int) string {
	s := "("

	for _, value := range slice {
		s += intToSql(value) + ", "
	}

	s = s[0:len(s)-2] + ")"

	return s
}

func intToSql(i int) string {
	return strconv.Itoa(i)
}

func int16ToSql(i int16) string {
	return strconv.Itoa(int(i))
}

func int32ToSql(i int32) string {
	return strconv.Itoa(int(i))
}

func int64ToSql(i int64) string {
	return strconv.Itoa(int(i))
}

func float64SliceToSql(slice []float64) string {
	s := "("

	for _, value := range slice {
		s += float64ToSql(value) + ", "
	}

	s = s[0:len(s)-2] + ")"

	return s
}

func float64ToSql(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func boolSliceToSql(slice []bool) string {
	s := "("

	for _, value := range slice {
		s += boolToSql(value) + ", "
	}

	s = s[0:len(s)-2] + ")"

	return s
}

func boolToSql(b bool) string {
	if b {
		return "TRUE"
	} else {
		return "FALSE"
	}
}
