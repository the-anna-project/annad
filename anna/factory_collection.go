package main

import (
	"github.com/cenk/backoff"

	"github.com/xh3b4sd/anna/factory"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/factory/permutation"
	"github.com/xh3b4sd/anna/factory/random"
	"github.com/xh3b4sd/anna/spec"
)

func newFactoryCollection() (spec.FactoryCollection, error) {
	randomFactory, err := newRandomFactory()
	if err != nil {
		return nil, maskAny(err)
	}
	idFactory, err := newIDFactory(randomFactory)
	if err != nil {
		return nil, maskAny(err)
	}
	permutationFactory, err := newPermutationFactory()
	if err != nil {
		return nil, maskAny(err)
	}

	newCollectionConfig := factory.DefaultCollectionConfig()
	newCollectionConfig.IDFactory = idFactory
	newCollectionConfig.PermutationFactory = permutationFactory
	newCollectionConfig.RandomFactory = randomFactory
	newCollection, err := factory.NewCollection(newCollectionConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newCollection, nil
}

func newIDFactory(randomFactory spec.RandomFactory) (spec.IDFactory, error) {
	newFactoryConfig := id.DefaultFactoryConfig()
	newFactoryConfig.RandomFactory = randomFactory
	newFactory, err := id.NewFactory(newFactoryConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newFactory, nil
}

func newPermutationFactory() (spec.PermutationFactory, error) {
	newFactoryConfig := permutation.DefaultFactoryConfig()
	newFactory, err := permutation.NewFactory(newFactoryConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newFactory, nil
}

func newRandomFactory() (spec.RandomFactory, error) {
	newFactoryConfig := random.DefaultFactoryConfig()
	newFactoryConfig.BackOffFactory = func() spec.BackOff {
		return backoff.NewExponentialBackOff()
	}
	newFactory, err := random.NewFactory(newFactoryConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newFactory, nil
}
