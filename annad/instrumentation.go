package main

import (
	"github.com/xh3b4sd/anna/instrumentation/prometheus"
	"github.com/xh3b4sd/anna/spec"
)

func newPrometheusInstrumentation(prefixes []string) (spec.Instrumentation, error) {
	newInstrumentationConfig := prometheus.DefaultInstrumentationConfig()
	newInstrumentationConfig.Prefixes = append(newInstrumentationConfig.Prefixes, prefixes...)
	newInstrumentation, err := prometheus.NewInstrumentation(newInstrumentationConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newInstrumentation, nil
}
