package permutation

import (
	"github.com/xh3b4sd/anna/spec"
)

func createIndizesWithDelta(list spec.PermutationList, delta int) ([]int, error) {
	// Initialize scope variables.
	base := len(list.GetValues())
	newIndizes := list.GetIndizes()
	operation := 0

	// Check for the initial situation. This is special and the only exception
	// within the algorithm.
	if len(newIndizes) == 0 {
		newIndizes = append(newIndizes, 0)
	}

	for {
		// Check amount of operations in the first place. That way the initial
		// situation as well as all other operations are covered.
		operation++
		if operation > delta {
			break
		}

		// Increment the least significant digit. That is, the right most index.
		// This is the only incrementation being done on the index.
		i := len(newIndizes) - 1
		lsd := newIndizes[i]
		lsd++
		newIndizes[i] = lsd

		// Cap the indizes and shift them if necessary. In case the least
		// significant digit was incremented above the base capacity, indizes need
		// to be shifted from right to left. This is like counting a number.
		var msdShifted bool
		newIndizes, msdShifted = shiftIndizes(newIndizes, base)
		if msdShifted {
			// Make sure the permutation does not growth more than allowed.
			if len(newIndizes)+1 > list.GetMaxGrowth() {
				return nil, maskAny(maxGrowthReachedError)
			}

			// In case all the indizes where shifted, we zeroed out all indizes. Then
			// we need to prepend another zero as new most significant digit of the
			// index.
			newIndizes = prepend(newIndizes, 0, 0)
		}
	}

	return newIndizes, nil
}

func createMembers(list spec.PermutationList) []interface{} {
	// Map values for easy access.
	mapped := map[int]interface{}{}
	for i, v := range list.GetValues() {
		mapped[i] = v
	}

	// Create the new permutation based on the updated index.
	newMembers := make([]interface{}, len(list.GetIndizes()))
	for i, index := range list.GetIndizes() {
		newMembers[i] = mapped[index]
	}

	return newMembers
}

func prepend(s []int, i, x int) []int {
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = x

	return s
}

func shiftIndizes(indizes []int, base int) ([]int, bool) {
	var msdShifted bool
	var reminder int

	for i := len(indizes) - 1; i >= 0; i-- {
		if reminder > 0 {
			current := indizes[i] + reminder
			reminder = 0
			indizes[i] = current
		}

		if indizes[i] >= base {
			indizes[i] = 0
			reminder = 1

			if i == 0 {
				msdShifted = true
			}
		}
	}

	return indizes, msdShifted
}
