package language

import (
	"fmt"
	"time"

	"github.com/xh3b4sd/anna/network"
)

func NewLanguageNetwork() network.Network {
	ln := languageNetwork{
		NetworkGateway: network.NewGateway(),
		NetworkState:   network.NewState(),
	}

	go ln.start()

	return ln
}

type languageNetwork struct {
	NetworkGateway network.Gateway
	NetworkState   network.State
}

func (ln languageNetwork) Age() time.Time {
	return time.Time{}
}

func (ln languageNetwork) Connections() ([]network.Connection, error) {
	return []network.Connection{}, nil
}

func (ln languageNetwork) Gateway() network.Gateway {
	return ln.NetworkGateway
}

func (ln languageNetwork) Load(state network.State) {
}

func (ln languageNetwork) Merge(dst, src interface{}) (interface{}, error) {
	return nil, nil
}

func (ln languageNetwork) Neurons() ([]network.Neuron, error) {
	return []network.Neuron{}, nil
}

func (ln languageNetwork) start() {
	stringGateway := ln.Gateway().String()

	for {
		select {
		case input := <-stringGateway:
			fmt.Printf("language network received string input: %s\n", input)
		}
	}
}

func (ln languageNetwork) State() network.State {
	return ln.NetworkState
}
