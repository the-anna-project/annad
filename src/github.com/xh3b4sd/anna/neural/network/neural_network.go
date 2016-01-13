package neuralnetwork

import (
	"fmt"
)

type NeuralNetwork interface {
	StringGateway() chan string
}

func NewNeuralNetwork() NeuralNetwork {
	nn := neuralNetwork{
		stringGateway: make(chan string, 1000),
	}

	go nn.start()

	return nn
}

type neuralNetwork struct {
	stringGateway chan string
}

func (nn neuralNetwork) StringGateway() chan string {
	return nn.stringGateway
}

func (nn neuralNetwork) start() {
	for {
		select {
		case input := <-nn.stringGateway:
			fmt.Printf("neural network received string input: %s\n", input)
		}
	}
}
