package permutationlist

import (
	objectspec "github.com/xh3b4sd/anna/object/spec"
)

func permuteValues(list objectspec.PermutationList) []interface{} {
	permutedValues := make([]interface{}, len(list.GetIndizes()))

	for i, index := range list.GetIndizes() {
		permutedValues[i] = list.GetRawValues()[index]
	}

	return permutedValues
}
