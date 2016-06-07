package main

import (
	"github.com/xh3b4sd/anna/index/clg"
	"github.com/xh3b4sd/anna/index/clg/collection"
	"github.com/xh3b4sd/anna/index/clg/profile"
	"github.com/xh3b4sd/anna/spec"
)

func createCLGIndex(newLog spec.Log, newStorage spec.Storage) (spec.CLGIndex, error) {
	// CLG collection
	newCLGCollectionConfig := collection.DefaultConfig()
	newCLGCollectionConfig.Log = newLog
	newCLGCollection, err := collection.New(newCLGCollectionConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	// CLG profile generator
	newGeneratorConfig := profile.DefaultGeneratorConfig()
	newGeneratorConfig.Collection = newCLGCollection
	newGeneratorConfig.Instrumentation, err = createPrometheusInstrumentation([]string{"CLGProfileGenerator"})
	if err != nil {
		return nil, maskAny(err)
	}
	newGeneratorConfig.Log = newLog
	newGeneratorConfig.Storage = newStorage
	newGenerator, err := profile.NewGenerator(newGeneratorConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	// CLG index
	newCLGIndexConfig := clg.DefaultIndexConfig()
	newCLGIndexConfig.Generator = newGenerator
	newCLGIndexConfig.Instrumentation, err = createPrometheusInstrumentation([]string{"CLGIndex"})
	if err != nil {
		return nil, maskAny(err)
	}
	newCLGIndexConfig.Log = newLog
	newCLGIndex, err := clg.NewIndex(newCLGIndexConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newCLGIndex, nil
}
