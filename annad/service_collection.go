package main

import (
	"github.com/cenk/backoff"

	"github.com/xh3b4sd/anna/service"
	"github.com/xh3b4sd/anna/service/fs/mem"
	"github.com/xh3b4sd/anna/service/id"
	"github.com/xh3b4sd/anna/service/permutation"
	"github.com/xh3b4sd/anna/service/random"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	"github.com/xh3b4sd/anna/service/textoutput"
	systemspec "github.com/xh3b4sd/anna/spec"
)

func newServiceCollection() (systemspec.ServiceCollection, error) {
	fileSystemService, err := newFileSystemService()
	if err != nil {
		return nil, maskAny(err)
	}
	randomService, err := newRandomService()
	if err != nil {
		return nil, maskAny(err)
	}
	idService, err := newIDService(randomService)
	if err != nil {
		return nil, maskAny(err)
	}
	permutationService, err := newPermutationService()
	if err != nil {
		return nil, maskAny(err)
	}
	textOutputService, err := newTextOutputService()
	if err != nil {
		return nil, maskAny(err)
	}

	newCollectionConfig := service.DefaultCollectionConfig()
	newCollectionConfig.FSService = fileSystemService
	newCollectionConfig.IDService = idService
	newCollectionConfig.PermutationService = permutationService
	newCollectionConfig.RandomService = randomService
	newCollectionConfig.TextOutputService = textOutputService
	newCollection, err := service.NewCollection(newCollectionConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newCollection, nil
}

// TODO make mem/os configurable
func newFileSystemService() (servicespec.FS, error) {
	newServiceConfig := mem.DefaultServiceConfig()
	newService, err := mem.NewService(newServiceConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newService, nil
}

func newIDService(randomService servicespec.Random) (servicespec.ID, error) {
	newServiceConfig := id.DefaultServiceConfig()
	newServiceConfig.RandomService = randomService
	newService, err := id.NewService(newServiceConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newService, nil
}

func newPermutationService() (servicespec.Permutation, error) {
	newServiceConfig := permutation.DefaultServiceConfig()
	newService, err := permutation.NewService(newServiceConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newService, nil
}

func newTextOutputService() (servicespec.TextOutput, error) {
	newServiceConfig := textoutput.DefaultServiceConfig()
	newService, err := textoutput.NewService(newServiceConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newService, nil
}

func newRandomService() (servicespec.Random, error) {
	newServiceConfig := random.DefaultServiceConfig()
	newServiceConfig.BackoffFactory = func() systemspec.Backoff {
		return backoff.NewExponentialBackOff()
	}
	newService, err := random.NewService(newServiceConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newService, nil
}
