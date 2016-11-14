package main

import (
	"os"

	"github.com/cenk/backoff"
	kitlog "github.com/go-kit/kit/log"

	"github.com/xh3b4sd/anna/service"
	"github.com/xh3b4sd/anna/service/fs/mem"
	"github.com/xh3b4sd/anna/service/id"
	"github.com/xh3b4sd/anna/service/log"
	"github.com/xh3b4sd/anna/service/permutation"
	"github.com/xh3b4sd/anna/service/random"
	"github.com/xh3b4sd/anna/service/spec"
	"github.com/xh3b4sd/anna/service/textinput"
	"github.com/xh3b4sd/anna/service/textoutput"
)

func (a *annactl) newServiceCollection() spec.Collection {
	// Set.
	collection := service.NewCollection()

	collection.SetFS(a.newFSService())
	collection.SetID(a.newIDService())
	collection.SetLog(a.newLogService())
	collection.SetPermutation(a.newPermutationService())
	collection.SetRandom(a.newRandomService())
	collection.SetTextInput(a.newTextInputService())
	collection.SetTextOutput(a.newTextOutputService())

	collection.FS().SetServiceCollection(collection)
	collection.ID().SetServiceCollection(collection)
	collection.Log().SetServiceCollection(collection)
	collection.Permutation().SetServiceCollection(collection)
	collection.Random().SetServiceCollection(collection)
	collection.TextInput().SetServiceCollection(collection)
	collection.TextOutput().SetServiceCollection(collection)

	return collection
}

// TODO make mem/os configurable
func (a *annactl) newFSService() spec.FS {
	return mem.New()
}

func (a *annactl) newIDService() spec.ID {
	return id.New()
}

func (a *annactl) newLogService() spec.Log {
	newService := log.New()

	newService.SetRootLogger(kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr)))

	return newService
}

func (a *annactl) newPermutationService() spec.Permutation {
	return permutation.New()
}

func (a *annactl) newRandomService() spec.Random {
	newService := random.New()

	newService.SetBackoffFactory(func() spec.Backoff {
		return backoff.NewExponentialBackOff()
	})

	return newService
}

func (a *annactl) newTextInputService() spec.TextInput {
	return textinput.New()
}

func (a *annactl) newTextOutputService() spec.TextOutput {
	return textoutput.New()
}
