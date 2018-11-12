package gojang

import (
	"fmt"
	"reflect"
	"strings"
)

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

//auxillary print function for testing purposes
func print(args ...interface{}) {
	fmt.Println(args...)
}

func typeOf(i interface{}) string {
	return reflect.TypeOf(i).String()
}

func doubleQuotes(s string) string {
	return "\"" + s + "\""
}

func dbq(s string) string {
	return doubleQuotes(s)
}
