package permutationlist

import (
	objectspec "github.com/the-anna-project/spec/object"
)

func permuteValues(list objectspec.PermutationList) []interface{} {
	permutedValues := make([]interface{}, len(list.GetIndizes()))

	for i, index := range list.GetIndizes() {
		permutedValues[i] = list.GetRawValues()[index]
	}

	return permutedValues
}
