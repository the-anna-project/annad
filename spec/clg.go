package spec

import (
	"reflect"
)

// CLG represents the CLGs interacting with each other within the neural
// network. Each CLG is registered in the Network. From there signal are
// dispatched in a dynamic fashion until some useful calculation took place.
type CLG interface {
	FactoryProvider

	// GetCalculate returns the CLG's calculate function which implements its
	// actual business logic.
	GetCalculate() interface{}

	// GetName returns the CLG's human readable name.
	GetName() string

	// GetInputChannel returns the CLG's input channel, which acts as
	// communication channel to reach the CLG inside of the neural network.
	GetInputChannel() chan NetworkPayload

	// GetInputTypes returns the CLG's underlying input types. These reflect the
	// real interface hidden behind the Calculate API.
	GetInputTypes() []reflect.Type

	Object

	// SetFactoryCollection configures the CLG's factory collection. This is done
	// for all CLGs, regardless if a CLG is making use of the factory collection
	// or not.
	SetFactoryCollection(factoryCollection FactoryCollection)

	// SetLog configures the CLG's logger. This is done for all CLGs, regardless
	// if a CLG is making use of the logger or not.
	SetLog(log Log)

	// SetStorageCollection configures the CLG's storage collection. This is done
	// for all CLGs, regardless if a CLG is making use of the storage collection
	// or not.
	SetStorageCollection(storageCollection StorageCollection)

	StorageProvider
}
