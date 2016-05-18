package profile

import (
	"crypto/rand"
	"math/big"
	"reflect"

	"github.com/xh3b4sd/anna/spec"
)

var (
	// argTypes represents a list of well known types used to identify CLG input-
	// and output types. Here we want to have a list of types only.
	//
	// TODO identify types by strings
	argTypes = []interface{}{
		// Simple types.
		*new(bool),
		*new(int),
		*new(float64),
		*new(string),
		*new(spec.Distribution),
		*new(string),

		// Slices of simple types.
		*new([]int),
		*new([]float64),
		*new([]string),

		// Slices of slices of simple types.
		*new([][]int),
		*new([][]float64),
	}

	// maxExamples represents the maximum number of inputsOutputs samples
	// provided in a CLG profile. A CLG profile may contain only one sample in
	// case the CLG interface is very strict. Nevertheless there might be CLGs
	// that accept a variadic amount of input parameters or return a variadic
	// amount of output results. The number of possible inputsOutputs samples can
	// be infinite in theory. Thus we cap the amount of inputsOutputs samples by
	// maxSamples.
	maxSamples = 10

	// numArgs is an ordered list of numbers used to find out how many input
	// arguments a CLG expects. Usually CLGs do not expect more than 5 input
	// arguments. For special cases we try to find out how many they expect
	// beyond 5 arguments. Here we assume that a CLG might expect 10 or even 50
	// arguments. In case a CLG expects 50 or more arguments, we assume it
	// expects infinite arguments.
	numArgs = []int{0, 1, 2, 3, 4, 5, 10, 20, 30, 40, 50}
)

type inputsOutputs struct {
	Inputs  []string
	Outputs []string
}

func (g *generator) CreateInputsOutputs(clgName string) ([]inputsOutputs, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call CreateInputs")

	min, max := g.getNumInputArgs(clgName)
	if err != nil {
		return nil, maskAny(err)
	}

	insOuts, err := g.getInputsOutputs(clgName, min, max)
	if err != nil {
		return nil, maskAny(err)
	}

	cappedInsOuts, err := limitInputsOutputs(insOuts, maxSamples)
	if err != nil {
		return nil, maskAny(err)
	}

	return cappedInsOuts, nil
}

// TODO test
func (g *generator) getNumInputArgs(clgName string) (int, int) {
	var min int
	var max int

	var minSet bool
	var maxSet bool

	for i, n := range numArgs {
		nArgs := getNArgs(n)

		_, err := g.Collection.CallByNameMethod(nArgs...)
		if collection.IsNotEnoughArguments(err) {
			// In case N args was not enough, we continue with the next number.
			continue
		} else if collection.IsTooManyArguments(err) {
			if i-1 < 0 {
				max = numArgs[0]
			} else {
				max = numArgs[i-1]
			}
		} else {
			// In case there was some other error or no error at all we
			if !minSet {
				min = n
				minSet = true
				max = n
			}
		}
	}

	return min, max
}

// TODO test
func (g *generator) getInputsOutputs(clgName string, min, max int) ([]inputsOutputs, error) {
	inputCombinations := getInputCombinations(min, max)
	var inputsOutputsList []inputsOutputs

	for _, in := range inputCombinations {
		out, err := g.Collection.CallByNameMethod(in...)
		if err != nil {
			continue
		}

		newInputsOutputs := inputsOutputs{
			Inputs:  typesToString(in),
			Outputs: typesToString(out),
		}
		inputsOutputsList = append(inputsOutputsList, newInputsOutputs)
	}

	return inputsOutputsList, nil
}

// TODO
func getInputCombinations(min, max int) [][]interface{} {
	// Create list of input types first.
	var typeCombinations [][]interface{}

	if matchesLength(argTypes, minLength, maxLength) {
		// Add the whole list to the set of combinations.
		typeCombinations = append(typeCombinations, argTypes)
	}

	for i, _ := range argTypes {
		comb := []interface{}{argTypes[i]}
		if !containsType(typeCombinations, comb) && matchesLength(comb, minLength, maxLength) {
			// Add the single type to the set of combinations.
			typeCombinations = append(typeCombinations, comb)
		}

		j := i
		for range argTypes {
			j++

			if j > len(argTypes) {
				break
			}

			comb := argTypes[i:j]
			if !containsType(typeCombinations, comb) && matchesLength(comb, minLength, maxLength) {
				// Add the matching subset to the list of combinations.
				typeCombinations = append(typeCombinations, comb)
			}
		}
	}

	// Transform list of input types to list of inputs consisting of real values.
	var newInputCombinations [][]interface{}

	for _, tc := range typeCombinations {
		var newCombination []interface{}

		for _, t := range tc {
			switch t.(type) {
			case bool:
			case int:
			case int64:
			case float64:
			case string:
			}
		}

		newInputCombinations = append(newInputCombinations, newCombination)
	}

	return newInputCombinations
}

func matchesLength(list []interface{}, min, max int) bool {
	if min != -1 && len(list) < min {
		return false
	}
	if max != -1 && len(list) > max {
		return false
	}
	return true
}

func containsType(list [][]interface{}, comb []interface{}) bool {
	for _, i := range list {
		if relfect.DeepEqual(i, comb) {
			return true
		}
	}

	return false
}

// TODO test
func limitInputsOutputs(inputsOutputsList []inputsOutputs, maxSamples int) ([]inputsOutputs, error) {
	l := len(inputsOutputsList)
	if l < maxSamples {
		return inputsOutputsList
	}

	// Depending on the number of inputsOutputs in relation to the given
	// maxSamples it is more efficient to either remove a few samples or to
	// select a few samples. In case maxSamples is set to 10 and there are 11
	// inputsOutputs in the given list, it would be smarter to remove one instead
	// of selecting 10.
	var numToSelect int
	var numToRemove int
	var selectRandom bool
	if l-maxSamples < 0 {
		selectRandom = true
	}
	if selectRandom {
		numToSelect = maxSamples
	} else {
		numToRemove = l - maxSamples
	}

	// Now that we know what strategy to apply we are going to find the required
	// amount of random indizes.
	alreadySeen := map[int64]struct{}{}
	var indizes []int
	for {
		max := big.NewInt(int64(l))
		j, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, maskAny(err)
		}
		r := j.Int64()
		if _, ok := alreadySeen[r]; ok {
			continue
		}
		alreadySeen[r] = struct{}{}

		indizes = append(indizes, r)

		if selectRandom {
			if len(indizes) >= numToSelect {
				break
			}
		} else {
			if len(indizes) >= numToRemove {
				break
			}
		}
	}

	// Finally we are ready to apply the prepared strategy.
	var newInputsOutputsList []inputsOutputs

	if selectRandom {
		for _, r := range indizes {
			newInputsOutputsList = append(newInputsOutputsList, inputsOutputsList[r])
		}
	} else {
		for _, r := range indizes {
			for i, insOuts := range inputsOutputsList {
				if i == r {
					continue
				}
				newInputsOutputsList = append(newInputsOutputsList, insOuts)
			}
		}
	}

	return newInputsOutputsList
}
