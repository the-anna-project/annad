package spec

// GatewayCollection represents a collection of gateways. This scopes
// different gateway implementations in a simple container, which can easily
// be passed around.
type GatewayCollection interface {
	// TextOutput returns an text output gateway. It is used to send text
	// responses back to the client.
	TextOutput() TextOutputGateway

	// TODO add TextInput
}

// GatewayProvider should be implemented by every object which wants to use
// gateways. This then creates an API between gateway implementations and
// gateway users.
type GatewayProvider interface {
	Gateway() GatewayCollection
}

// TextOutputGateway provides a communication channel to send information
// sequences back to the client.
type TextOutputGateway interface {
	// GetChannel returns a channel which is used to send text responses back to
	// the client.
	GetChannel() chan TextResponse
}
