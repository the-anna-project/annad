package common

import (
	"github.com/xh3b4sd/anna/spec"
)

// network

// Note that when adding a new object type to this struct, the two definitions
// below need to be kept in sync as well.
type networkObjectType struct {
	StrategyNetwork spec.ObjectType
}

var (
	// Note that this struct needs to be in sync with the below list. E.g. the
	// logger checks for valid object types using this list.
	NetworkObjectType = networkObjectType{
		StrategyNetwork: spec.ObjectType("strategy-network"),
	}

	// Note that this list needs to be in sync with the above struct. E.g. the
	// logger checks for valid object types using this list.
	NetworkObjectTypes = []spec.ObjectType{
		NetworkObjectType.StrategyNetwork,
	}
)

// neuron

// Note that when adding a new object type to this struct, the two definitions
// below need to be kept in sync as well.
type neuronObjectType struct {
}

var (
	// Note that this struct needs to be in sync with the below list. E.g. the
	// logger checks for valid object types using this list.
	NeuronObjectType = neuronObjectType{}

	// Note that this list needs to be in sync with the above struct. E.g. the
	// logger checks for valid object types using this list.
	NeuronObjectTypes = []spec.ObjectType{}
)

// storage

// Note that when adding a new object type to this struct, the two definitions
// below need to be kept in sync as well.
type storageObjectType struct {
	MemoryStorage spec.ObjectType
	RedisStorage  spec.ObjectType
}

var (
	// Note that this struct needs to be in sync with the below list. E.g. the
	// logger checks for valid object types using this list.
	StorageObjectType = storageObjectType{
		MemoryStorage: spec.ObjectType("memory-storage"),
		RedisStorage:  spec.ObjectType("redis-storage"),
	}

	// Note that this list needs to be in sync with the above struct. E.g. the
	// logger checks for valid object types using this list.
	StorageObjectTypes = []spec.ObjectType{
		StorageObjectType.MemoryStorage,
		StorageObjectType.RedisStorage,
	}
)

// all

// Note that when adding a new object type to this struct, the two definitions
// below need to be kept in sync as well.
type objectType struct {
	Anna          spec.ObjectType
	Core          spec.ObjectType
	Gateway       spec.ObjectType
	Impulse       spec.ObjectType
	FactoryClient spec.ObjectType
	FactoryServer spec.ObjectType
	Log           spec.ObjectType
	LogControl    spec.ObjectType

	networkObjectType

	neuronObjectType

	None   spec.ObjectType
	Server spec.ObjectType
	State  spec.ObjectType

	storageObjectType

	TextInterface spec.ObjectType
}

var (
	// Note that this struct needs to be in sync with the below list. E.g. the
	// logger checks for valid object types using this list.
	ObjectType = objectType{
		Anna:          spec.ObjectType("anna"),
		Core:          spec.ObjectType("core"),
		Gateway:       spec.ObjectType("gateway"),
		Impulse:       spec.ObjectType("impulse"),
		FactoryClient: spec.ObjectType("factory-client"),
		FactoryServer: spec.ObjectType("factory-server"),
		Log:           spec.ObjectType("log"),
		LogControl:    spec.ObjectType("log-control"),

		networkObjectType: NetworkObjectType,

		neuronObjectType: NeuronObjectType,

		None:   spec.ObjectType("none"),
		Server: spec.ObjectType("server"),
		State:  spec.ObjectType("state"),

		storageObjectType: StorageObjectType,

		TextInterface: spec.ObjectType("text-interface"),
	}

	// Note that this list needs to be in sync with the above struct. E.g. the
	// logger checks for valid object types using this list.
	ObjectTypes = []spec.ObjectType{
		ObjectType.Anna,
		ObjectType.Core,
		ObjectType.Gateway,
		ObjectType.Impulse,
		ObjectType.FactoryClient,
		ObjectType.FactoryServer,
		ObjectType.Log,
		ObjectType.LogControl,

		// Network.
		ObjectType.StrategyNetwork,

		ObjectType.None,
		ObjectType.Server,
		ObjectType.State,

		// Storage.
		ObjectType.MemoryStorage,
		ObjectType.RedisStorage,

		ObjectType.TextInterface,
	}
)
