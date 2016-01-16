package language

import (
	"fmt"
	"time"

	"github.com/xh3b4sd/anna/network"
)

func NewNetwork() network.Network {
	n := network{
		Gateway: network.NewGateway(),
		State:   network.NewState(),
	}

	go n.start()

	return n
}

type network struct {
	CreatedAt time.Duration `json:"created_at"`
	Gateway   network.Gateway
	State     network.State
}

func (n network) Age() time.Duration {
	return time.Since(lb.CreatedAt)
}

func (n network) Connections() ([]network.Connection, error) {
	return []network.Connection{}, nil
}

func (n network) Gateway() network.Gateway {
	return n.Gateway
}

func (n network) Load(state network.State) {
}

func (n network) Merge(dst, src interface{}) (interface{}, error) {
	return nil, nil
}

func (n network) Neurons() ([]network.Neuron, error) {
	return []network.Neuron{}, nil
}

func (n network) start() {
	stringGateway := n.Gateway().String()

	for {
		select {
		case input := <-stringGateway:
			fmt.Printf("language network received string input: %s\n", input)
		}
	}
}

func (n network) State() network.State {
	return n.State.Capture(n)
}
