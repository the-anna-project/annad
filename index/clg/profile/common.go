package profile

import (
	"reflect"
)

func typesToString(arguments []interface{}) []string {
	var newStrings []string

	for _, a := range arguments {
		newStrings = append(newStrings, reflect.TypeOf(a).String())
	}

	return newStrings
}
