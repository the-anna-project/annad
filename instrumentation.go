package main

import (
	"github.com/xh3b4sd/anna/instrumentation/prometheus"
	"github.com/xh3b4sd/anna/spec"
)

func createPrometheusInstrumentation(prefixes []string) (spec.Instrumentation, error) {
	newPrometheusConfig := prometheus.DefaultConfig()
	newPrometheusConfig.Prefixes = append(newPrometheusConfig.Prefixes, prefixes...)
	newInstrumentation, err := prometheus.New(newPrometheusConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newInstrumentation, nil
}
