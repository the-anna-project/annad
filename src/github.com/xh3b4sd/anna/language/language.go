package languagenetwork

import (
	"fmt"

	"github.com/xh3b4sd/anna/network"
)

type LanguageNetwork interface {
	ImpulsGateway() chan network.Impuls
}

func NewNeuralNetwork() LanguageNetwork {
	ln := languageNetwork{
		impulsGateway: make(chan network.Impuls, 1000),
	}

	go ln.start()

	return ln
}

type languageNetwork struct {
	impulsGateway chan network.Impuls
}

func (ln languageNetwork) ImpulsGateway() chan network.Impuls {
	return ln.impulsGateway
}

func (ln languageNetwork) start() {
	for {
		select {
		case impuls := <-ln.impulsGateway:
			fmt.Printf("language network received impuls: %s\n", impuls.String())
		}
	}
}
