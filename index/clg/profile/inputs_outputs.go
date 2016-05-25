package profile

import (
	"github.com/xh3b4sd/anna/factory/permutation"
	"github.com/xh3b4sd/anna/index/clg/collection"
	"github.com/xh3b4sd/anna/spec"
)

var (
	// maxSamples represents the maximum number of InOut samples provided in a
	// CLG profile. A CLG profile may contain only one sample in case the CLG
	// interface is very strict. Nevertheless there might be CLGs that accept a
	// variadic amount of input parameters. The number of possible InOut samples
	// can be infinite in theory. Thus we cap the amount of InOut samples by
	// maxSamples.
	//
	// Note that this number is experimental and might change if necessary.
	maxSamples = 10
)

// TODO test
func (g *generator) CreateInputsOutputs(clgName string) (spec.InputsOutputs, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call CreateInputs")

	newInputsOutputs, err := g.getInputsOutputs(clgName)
	if err != nil {
		return spec.InputsOutputs{}, maskAny(err)
	}

	newInputsOutputs.InsOuts, err = g.limitInputsOutputs(newInputsOutputs.InsOuts, maxSamples)
	if err != nil {
		return spec.InputsOutputs{}, maskAny(err)
	}

	return newInputsOutputs, nil
}

func (g *generator) getInputsOutputs(clgName string) (spec.InputsOutputs, error) {
	newInputsOutputs := spec.InputsOutputs{}

	for {
		// Perform the permutations to fetch possible combinations of input
		// arguments for the CLG execution.
		err := g.PermutationFactory.PermuteBy(g.ArgumentList, 1)
		if permutation.IsMaxGrowthReached(err) {
			// We are through with all possible combinations. Thus return what we have so far.
			return newInputsOutputs, nil
		} else if err != nil {
			return spec.InputsOutputs{}, maskAny(err)
		}

		in := g.ArgumentList.GetMembers()
		out, err := g.Collection.CallByNameMethod(in...)
		if collection.IsNotEnoughArguments(err) {
			// The number of input arguments is lesser than the CLG interface
			// actually requires. Thus we go ahead to check the next permutation.
			continue
		} else if collection.IsTooManyArguments(err) {
			// The number of input arguments is greater than the CLG interface
			// actually requires. Thus we return what we have so far, because there
			// is nothing to add.
			return newInputsOutputs, nil
		} else if err != nil {
			// Some unknown error happened. Keep trying.
			continue
		}

		newInOut := spec.InOut{
			In:  typesToString(in),
			Out: typesToString(out),
		}
		newInputsOutputs.InsOuts = append(newInputsOutputs.InsOuts, newInOut)
	}
}

func (g *generator) limitInputsOutputs(insOuts []spec.InOut, maxSamples int) ([]spec.InOut, error) {
	if len(insOuts) < maxSamples {
		return insOuts, nil
	}

	n := maxSamples
	max := len(insOuts)

	newRandomNumbers, err := g.RandomFactory.CreateNMax(n, max)
	if err != nil {
		return nil, maskAny(err)
	}

	newInsOuts := make([]spec.InOut, len(newRandomNumbers))

	for i, r := range newRandomNumbers {
		newInsOuts[i] = insOuts[r]
	}

	return newInsOuts, nil
}
