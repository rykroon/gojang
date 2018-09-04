package models

import ()

//auxillary functions used in all programs

func doubleQuotes(s string) string {
	return "\"" + s + "\""
}

func singleQuotes(s string) string {
	return "'" + s + "'"
}

//Constants

const Cascade string = "CASCADE"
const Protect string = "RESTRICT"
const SetNull string = "SET NULL"
const SetDefault string = "SSET DEFAULT"
