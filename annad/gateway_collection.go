package main

import (
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/gateway/text-output"
	"github.com/xh3b4sd/anna/spec"
)

func newGatewayCollection(newTextOutput chan spec.TextResponse) (spec.GatewayCollection, error) {
	textOutputGateway, err := newTextOutputGateway(newTextOutput)
	if err != nil {
		return nil, maskAny(err)
	}

	newCollectionConfig := gateway.DefaultCollectionConfig()
	newCollectionConfig.TextOutputGateway = textOutputGateway
	newCollection, err := gateway.NewCollection(newCollectionConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newCollection, nil
}

func newTextOutputGateway(newTextOutput chan spec.TextResponse) (spec.TextOutputGateway, error) {
	newGatewayConfig := textoutput.DefaultGatewayConfig()
	newGatewayConfig.Channel = newTextOutput
	newGateway, err := textoutput.NewGateway(newGatewayConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newGateway, nil
}
