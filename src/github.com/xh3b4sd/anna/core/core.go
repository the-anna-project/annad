package core

import (
	"fmt"
)

type CoreConfig struct {
	Gateway Gateway

	LanguageNetwork Network

	State State
}

func DefaultCoreConfig() CoreConfig {
	return CoreConfig{
		Gateway:         NewGateway(),
		LanguageNetwork: nil,
		State:           NewState(),
	}
}

type Core interface {
	// Boot initializes and starts the whole core like booting a machine. The
	// call to Boot blocks until the core is completely initialized, so you might
	// want to call it in a separate goroutine.
	Boot()

	GetGateway() Gateway

	SetState(state State)

	// Shutdown ends all processes of the core like shutting down a machine. The
	// call to Boot blocks until the core is completely shut down, so you might
	// want to call it in a separate goroutine.
	Shutdown()

	GetState() State
}

func NewCore(config CoreConfig) Core {
	return core{
		CoreConfig: config,
	}
}

type core struct {
	CoreConfig
}

func (c core) Boot() {
	go c.listen()
}

func (c core) listen() {
	stringChannel := c.GetGateway().GetTextGateway().GetStringChannel()

	for {
		select {
		case input := <-stringChannel:
			fmt.Printf("core received string input: %s\n", input)
		}
	}
}

func (c core) GetGateway() Gateway {
	return c.Gateway
}

func (c core) SetState(state State) {
	c.State = state
}

func (c core) Shutdown() {
}

func (c core) GetState() State {
	return c.State
}
