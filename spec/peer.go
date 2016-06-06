package spec

import (
	"encoding/json"
)

// TODO implement

// Peer represents something like a node within the knowledge network. There
// peers are connected with each other. Thus, forming a network of learned
// knowledge.
type Peer interface {
	// AddPeer adds the given peer to the current peer. That way a connection
	// between both peers is created within the knowledge network by requesting
	// the connection creation over the configured gateway.
	AddPeer(peer Peer) error

	// GetGateway returns the configured gateway. This should be the gateway
	// connected to the knowledge network.
	GetGateway() Gateway

	json.Marshaler

	json.Unmarshaler

	Object
}
