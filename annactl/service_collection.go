package main

import (
	"os"

	kitlog "github.com/go-kit/kit/log"

	"github.com/the-anna-project/collection"
	"github.com/the-anna-project/fs/memory"
	"github.com/the-anna-project/id"
	"github.com/the-anna-project/log"
	servicespec "github.com/the-anna-project/spec/service"
	"github.com/xh3b4sd/anna/service/permutation"
	"github.com/xh3b4sd/anna/service/textinput"
	"github.com/xh3b4sd/anna/service/textoutput"
)

func (a *annactl) newServiceCollection() servicespec.ServiceCollection {
	// Set.
	collection := collection.New()

	collection.SetFSService(a.newFSService())
	collection.SetIDService(a.newIDService())
	collection.SetLogService(a.newLogService())
	collection.SetPermutationService(a.newPermutationService())
	collection.SetTextInputService(a.newTextInputService())
	collection.SetTextOutputService(a.newTextOutputService())

	collection.FS().SetServiceCollection(collection)
	collection.ID().SetServiceCollection(collection)
	collection.Log().SetServiceCollection(collection)
	collection.Permutation().SetServiceCollection(collection)
	collection.TextInput().SetServiceCollection(collection)
	collection.TextOutput().SetServiceCollection(collection)

	return collection
}

// TODO make mem/os configurable
func (a *annactl) newFSService() servicespec.FSService {
	return memory.New()
}

func (a *annactl) newIDService() servicespec.IDService {
	return id.New()
}

func (a *annactl) newLogService() servicespec.LogService {
	newService := log.New()

	newService.SetRootLogger(kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr)))

	return newService
}

func (a *annactl) newPermutationService() servicespec.PermutationService {
	return permutation.New()
}

func (a *annactl) newTextInputService() servicespec.TextInputService {
	return textinput.New()
}

func (a *annactl) newTextOutputService() servicespec.TextOutputService {
	return textoutput.New()
}
