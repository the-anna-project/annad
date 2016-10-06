package textoutput

import (
	"github.com/xh3b4sd/anna/spec"
)

// GatewayConfig represents the configuration used to create a new text output
// gateway object.
type GatewayConfig struct {
	// Settings.
	Channel chan spec.TextResponse
}

// DefaultGatewayConfig provides a default configuration to create a new text
// output gateway object by best effort.
func DefaultGatewayConfig() GatewayConfig {
	newConfig := GatewayConfig{
		// Settings.
		Channel: make(chan spec.TextResponse, 1000),
	}

	return newConfig
}

// NewGateway creates a new configured text output gateway object.
func NewGateway(config GatewayConfig) (spec.TextOutputGateway, error) {
	newGateway := &gateway{
		GatewayConfig: config,
	}

	if newGateway.Channel == nil {
		return nil, maskAnyf(invalidConfigError, "channel must not be empty")
	}

	return newGateway, nil
}

// MustNewGateway creates either a new default configured id gateway object, or
// panics.
func MustNewGateway() spec.TextOutputGateway {
	newTextOutputGateway, err := NewGateway(DefaultGatewayConfig())
	if err != nil {
		panic(err)
	}

	return newTextOutputGateway
}

type gateway struct {
	GatewayConfig
}

func (g *gateway) GetChannel() chan spec.TextResponse {
	return g.Channel
}
