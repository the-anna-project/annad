package spec

import (
	"time"
)

type State interface {
	Copy() State

	GetAge() time.Duration

	GetBytes(key string) ([]byte, error)

	GetCoreByID(objectID ObjectID) (Core, error)

	GetCores() map[ObjectID]Core

	GetObjectID() ObjectID

	GetImpulseByID(objectID ObjectID) (Impulse, error)

	GetImpulses() map[ObjectID]Impulse

	GetNetworkByID(objectID ObjectID) (Network, error)

	GetNetworks() map[ObjectID]Network

	GetNeuronByID(objectID ObjectID) (Neuron, error)

	GetNeurons() map[ObjectID]Neuron

	GetObjectType() ObjectType

	SetBytes(key string, bytes []byte)

	SetCore(core Core)

	SetImpulse(imp Impulse)

	SetNetwork(network Network)

	SetNeuron(neu Neuron)
}
