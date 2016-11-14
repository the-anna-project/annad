package spec

import (
	"github.com/xh3b4sd/anna/object/config"
	"github.com/xh3b4sd/anna/object/config/endpoint"
	"github.com/xh3b4sd/anna/object/config/space"
	"github.com/xh3b4sd/anna/object/config/storage"
)

// Config represents a configuration service providing configuration for all
// services within the service collection.
type Config interface {
	Boot()
	Endpoint() *endpoint.Collection
	Metadata() map[string]string
	Service() Collection
	SetConfigCollection(configCollection *config.Collection)
	SetServiceCollection(serviceCollection Collection)
	Space() *space.Collection
	Storage() *storage.Collection
}
