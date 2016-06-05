package knowledge

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectTypeKnowledgeNetwork represents the object type of the knowledge
	// network object. This is used e.g. to register itself to the logger.
	ObjectTypeKnowledgeNetwork spec.ObjectType = "knowledge-network"
)

// NetworkConfig represents the configuration used to create a new knowledge
// network object.
type NetworkConfig struct {
	// Dependencies.
	Log     spec.Log
	Storage spec.Storage
}

// DefaultNetworkConfig provides a default configuration to create a new
// knowledge network object by best effort.
func DefaultNetworkConfig() NetworkConfig {
	newStorage, err := memory.NewStorage(memory.DefaultStorageConfig())
	if err != nil {
		panic(err)
	}

	newConfig := NetworkConfig{
		Log:     log.NewLog(log.DefaultConfig()),
		Storage: newStorage,
	}

	return newConfig
}

// NewNetwork creates a new configured knowledge network object.
func NewNetwork(config NetworkConfig) (spec.Network, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		return nil, maskAny(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		return nil, maskAny(err)
	}

	newNetwork := &network{
		NetworkConfig: config,

		BootOnce:     sync.Once{},
		ID:           newID,
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         ObjectTypeKnowledgeNetwork,
	}

	newNetwork.Log.Register(newNetwork.GetType())

	return newNetwork, nil
}

type network struct {
	NetworkConfig

	BootOnce     sync.Once
	ID           spec.ObjectID
	Mutex        sync.Mutex
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (n *network) Boot() {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Boot")

	n.BootOnce.Do(func() {
	})
}

func (n *network) Shutdown() {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Shutdown")

	n.ShutdownOnce.Do(func() {
	})
}

func (n *network) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call Trigger")

	return nil, nil
}
