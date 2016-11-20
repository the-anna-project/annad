package list

import (
	objectspec "github.com/the-anna-project/spec/object"
)

func permuteValues(list objectspec.PermutationList) []interface{} {
	permutedValues := make([]interface{}, len(list.Indizes()))

	for i, index := range list.Indizes() {
		permutedValues[i] = list.RawValues()[index]
	}

	return permutedValues
}
