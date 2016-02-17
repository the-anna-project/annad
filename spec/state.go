package spec

import (
	"encoding/json"
	"time"
)

type StateType string

type State interface {
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

	GetVersion() string

	json.Unmarshaler

	SetBytes(key string, bytes []byte)

	SetCore(core Core)

	SetImpulse(imp Impulse)

	SetNetwork(network Network)

	SetNeuron(neu Neuron)

	// SetStateFromObjectBytes takes the byte slice provided by an object. The
	// raw state structure is pulled out of the object structure and set as the
	// current state. So when restoring a backup any object state can be restored
	// by extracting the raw state structure and applying it as new object state.
	SetStateFromObjectBytes(bytes []byte) error

	StateReader

	StateWriter

	SetVersion(version string)
}

type StateReader interface {
	// Read loads state based on the given state reader configuration.
	Read() error

	ReadFile(filename string) error
}

type StateWriter interface {
	// Write persists the current state based on the given state writer
	// configuration.
	Write() error

	WriteFile(filename string) error
}
