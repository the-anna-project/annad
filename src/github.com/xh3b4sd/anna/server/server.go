package server

import (
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/neural/network"
	"github.com/xh3b4sd/anna/server/interface/text"
)

func Listen() {
	ctx := context.Background()

	// neural network
	nn := neuralnetwork.NewNeuralNetwork()

	// text interface
	config := textinterface.NewTextInterfaceConfig{
		StringGateway: nn.StringGateway(),
	}
	ti := textinterface.NewTextInterface(config)
	handlers := textinterface.NewHandlers(ctx, ti)

	// http
	for url, handler := range handlers {
		http.Handle(url, handler)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}
