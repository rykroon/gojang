package gojang

import (
	"strconv"
	"strings"
)

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

func boolAsSql(b bool) string {
	if b {
		return "TRUE"
	} else {
		return "FALSE"
	}
}

func intAsSql(i int) string {
	return strconv.Itoa(i)
}

func int16AsSql(i int16) string {
	return intAsSql(int(i))
}

func int32AsSql(i int32) string {
	return intAsSql(int(i))
}

func int64AsSql(i int64) string {
	return intAsSql(int(i))
}

func float64AsSql(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func stringAsSql(s string) string {
	return singleQuotes(s)
}

func boolSliceAsSql(booleans []bool) string {
	var valueList []string
	for _, value := range booleans {
		valueList = append(valueList, boolAsSql(value))
	}

	return "(" + strings.Join(valueList, ", ") + ")"
}

func float64SliceAsSql(floats []float64) string {
	var valueList []string
	for _, value := range floats {
		valueList = append(valueList, float64AsSql(value))
	}

	return "(" + strings.Join(valueList, ", ") + ")"
}

func integersAsSql(integers []int) string {
	var valueList []string
	for _, value := range integers {
		valueList = append(valueList, intAsSql(value))
	}

	return "(" + strings.Join(valueList, ", ") + ")"
}

func stringSliceAsSql(valueList []string) string {
	return "(" + strings.Join(valueList, ", ") + ")"
}
