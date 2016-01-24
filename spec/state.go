package spec

import (
	"time"
)

type State interface {
	GetAge() time.Duration

	GetBytes(key string) ([]byte, error)

	GetCore() (Core, error)

	GetObjectID() ObjectID

	GetImpulses() map[ObjectID]Impulse

	GetNetworkByID(objectID ObjectID) (Network, error)

	GetNetworks() map[ObjectID]Network

	GetNeuronByID(objectID ObjectID) (Neuron, error)

	GetNeurons() map[ObjectID]Neuron

	GetObjectType() ObjectType

	SetBytes(key string, bytes []byte)

	SetCore(core Core)

	SetImpulse(imp Impulse)

	SetNetwork(neu Network)

	SetNeuron(neu Neuron)
}
