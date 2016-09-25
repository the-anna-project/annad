package network

import (
	"github.com/xh3b4sd/anna/clg/divide"
	"github.com/xh3b4sd/anna/clg/greater"
	"github.com/xh3b4sd/anna/clg/input"
	"github.com/xh3b4sd/anna/clg/is-between"
	"github.com/xh3b4sd/anna/clg/is-greater"
	"github.com/xh3b4sd/anna/clg/is-lesser"
	"github.com/xh3b4sd/anna/clg/lesser"
	"github.com/xh3b4sd/anna/clg/multiply"
	"github.com/xh3b4sd/anna/clg/output"
	"github.com/xh3b4sd/anna/clg/pair-syntactic"
	"github.com/xh3b4sd/anna/clg/read-information-id"
	"github.com/xh3b4sd/anna/clg/read-separator"
	"github.com/xh3b4sd/anna/clg/round"
	"github.com/xh3b4sd/anna/clg/split-features"
	"github.com/xh3b4sd/anna/clg/subtract"
	"github.com/xh3b4sd/anna/clg/sum"
	"github.com/xh3b4sd/anna/spec"
)

// receiver

func (n *network) clgByName(name string) (spec.CLG, error) {
	ID, ok := n.CLGIDs[name]
	if !ok {
		return nil, maskAnyf(clgNotFoundError, "name: %s", name)
	}
	CLG, ok := n.CLGs[ID]
	if !ok {
		return nil, maskAnyf(clgNotFoundError, "ID: %s", ID)
	}

	return CLG, nil
}

func (n *network) configureCLGs(CLGs map[spec.ObjectID]spec.CLG) map[spec.ObjectID]spec.CLG {
	for ID := range CLGs {
		CLGs[ID].SetFactoryCollection(n.FactoryCollection)
		CLGs[ID].SetLog(n.Log)
		CLGs[ID].SetStorageCollection(n.StorageCollection)
	}

	return CLGs
}

func (n *network) mapCLGIDs(CLGs map[spec.ObjectID]spec.CLG) map[string]spec.ObjectID {
	clgIDs := map[string]spec.ObjectID{}

	for ID, CLG := range CLGs {
		clgIDs[CLG.GetName()] = ID
	}

	return clgIDs
}

// helper

func newCLGs() map[spec.ObjectID]spec.CLG {
	newList := []spec.CLG{
		divide.MustNew(),
		input.MustNew(),
		divide.MustNew(),
		greater.MustNew(),
		input.MustNew(),
		isbetween.MustNew(),
		isgreater.MustNew(),
		islesser.MustNew(),
		lesser.MustNew(),
		multiply.MustNew(),
		output.MustNew(),
		pairsyntactic.MustNew(),
		readinformationid.MustNew(),
		readseparator.MustNew(),
		round.MustNew(),
		splitfeatures.MustNew(),
		subtract.MustNew(),
		sum.MustNew(),
	}

	newCLGs := map[spec.ObjectID]spec.CLG{}

	for _, CLG := range newList {
		newCLGs[CLG.GetID()] = CLG
	}

	return newCLGs
}
