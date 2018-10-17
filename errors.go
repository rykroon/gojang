package gojang

import (
	"errors"
)

var multipleObjectsReturned error = errors.New("Multiple objects returned")
var doesNotExist error = errors.New("Does not exist")
