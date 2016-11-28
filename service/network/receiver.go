package network

import (
	"github.com/the-anna-project/annad/service/clg/divide"
	"github.com/the-anna-project/annad/service/clg/greater"
	"github.com/the-anna-project/annad/service/clg/input"
	"github.com/the-anna-project/annad/service/clg/isbetween"
	"github.com/the-anna-project/annad/service/clg/isgreater"
	"github.com/the-anna-project/annad/service/clg/islesser"
	"github.com/the-anna-project/annad/service/clg/lesser"
	"github.com/the-anna-project/annad/service/clg/multiply"
	"github.com/the-anna-project/annad/service/clg/output"
	"github.com/the-anna-project/annad/service/clg/pairsyntactic"
	"github.com/the-anna-project/annad/service/clg/readinformationid"
	"github.com/the-anna-project/annad/service/clg/readseparator"
	"github.com/the-anna-project/annad/service/clg/round"
	"github.com/the-anna-project/annad/service/clg/splitfeatures"
	"github.com/the-anna-project/annad/service/clg/subtract"
	"github.com/the-anna-project/annad/service/clg/sum"
	servicespec "github.com/the-anna-project/spec/service"
)

// newCLGs returns a list of all CLGs which are configured and ready to be used
// within the neural network.
func (s *service) newCLGs() map[string]servicespec.CLGService {
	// TODO this should be initialized with the service collection
	list := []servicespec.CLGService{
		divide.New(),
		input.New(),
		divide.New(),
		greater.New(),
		input.New(),
		isbetween.New(),
		isgreater.New(),
		islesser.New(),
		lesser.New(),
		multiply.New(),
		output.New(),
		pairsyntactic.New(),
		readinformationid.New(),
		readseparator.New(),
		round.New(),
		splitfeatures.New(),
		subtract.New(),
		sum.New(),
	}

	newCLGs := map[string]servicespec.CLGService{}

	for _, CLG := range list {
		newCLGs[CLG.Metadata()["name"]] = CLG
	}

	for name := range newCLGs {
		newCLGs[name].SetServiceCollection(s.serviceCollection)
	}

	return newCLGs
}
