package main

import (
	"os"

	kitlog "github.com/go-kit/kit/log"

	"github.com/the-anna-project/collection"
	"github.com/the-anna-project/fs/memory"
	"github.com/the-anna-project/id"
	inputcollection "github.com/the-anna-project/input/collection"
	textinputservice "github.com/the-anna-project/input/service/text"
	"github.com/the-anna-project/log"
	outputcollection "github.com/the-anna-project/output/collection"
	textoutputservice "github.com/the-anna-project/output/service/text"
	"github.com/the-anna-project/permutation/service"
	servicespec "github.com/the-anna-project/spec/service"
)

func (a *annactl) newServiceCollection() servicespec.ServiceCollection {
	// Set.
	collection := collection.New()

	collection.SetFSService(a.newFSService())
	collection.SetIDService(a.newIDService())
	collection.SetInputCollection(a.newInputCollection())
	collection.SetLogService(a.newLogService())
	collection.SetOutputCollection(a.newOutputCollection())
	collection.SetPermutationService(a.newPermutationService())

	collection.FS().SetServiceCollection(collection)
	collection.ID().SetServiceCollection(collection)
	collection.Input().Text().SetServiceCollection(collection)
	collection.Log().SetServiceCollection(collection)
	collection.Output().Text().SetServiceCollection(collection)
	collection.Permutation().SetServiceCollection(collection)

	return collection
}

// TODO make mem/os configurable
func (a *annactl) newFSService() servicespec.FSService {
	return memory.New()
}

func (a *annactl) newIDService() servicespec.IDService {
	return id.New()
}

func (a *annactl) newInputCollection() servicespec.InputCollection {
	newCollection := inputcollection.New()

	newCollection.SetTextService(textinputservice.New())

	return newCollection
}

func (a *annactl) newLogService() servicespec.LogService {
	newService := log.New()

	newService.SetRootLogger(kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr)))

	return newService
}

func (a *annactl) newOutputCollection() servicespec.OutputCollection {
	newCollection := outputcollection.New()

	newCollection.SetTextService(textoutputservice.New())

	return newCollection
}

func (a *annactl) newPermutationService() servicespec.PermutationService {
	return permutation.New()
}
