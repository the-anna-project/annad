package config

import (
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/xh3b4sd/anna/object/config/config"
	"github.com/xh3b4sd/anna/object/config/endpoint"
	"github.com/xh3b4sd/anna/object/config/space"
	"github.com/xh3b4sd/anna/object/config/storage"
)

// NewCollection creates a new config collection. It provides configuration for
// the whole neural network.
func NewCollection() *Collection {
	return &Collection{}
}

// Collection represents the config collection.
type Collection struct {
	// Settings.

	endpointCollection *endpoint.Collection
	config             *config.Object
	spaceCollection    *space.Collection
	storageCollection  *storage.Collection
}

// Config returns the config file config of the config collection.
func (c *Collection) Config() *config.Object {
	return c.config
}

// Endpoint returns the endpoint collection of the config collection.
func (c *Collection) Endpoint() *endpoint.Collection {
	return c.endpointCollection
}

// Merge combines values of a flag-set with these of their corresponding
// environment and config file variables, in this order.
func (c *Collection) Merge(flagSet *pflag.FlagSet) error {
	v := viper.New()

	// Check the defined config file.
	v.AddConfigPath(c.Config().Dir())
	v.SetConfigName(c.Config().Name())
	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// In case there is no config file given we simply go ahead to check the
			// process environment.
		} else {
			return maskAny(err)
		}
	}

	// We merge the defined flags with their upper case counterparts from the
	// environment .
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.BindPFlags(flagSet)

	flagSet.VisitAll(func(flag *pflag.Flag) {
		if flag.Changed {
			// The current flag was set via the command line. We definitely want to use
			// the set value. Therefore we do not merge anything into it.
			return
		}
		if !v.IsSet(flag.Name) {
			// There is neither configuration in the provided config file nor in the
			// process environment. That means we cannot use it to merge it into any
			// defined flag.
			return
		}

		flag.Value.Set(v.GetString(flag.Name))
	})

	return nil
}

// SetConfig sets the config file config for the config collection.
func (c *Collection) SetConfig(config *config.Object) {
	c.config = config
}

// SetEndpointCollection sets the endpoint collection for the config collection.
func (c *Collection) SetEndpointCollection(endpointCollection *endpoint.Collection) {
	c.endpointCollection = endpointCollection
}

// SetSpaceCollection sets the space collection for the config collection.
func (c *Collection) SetSpaceCollection(spaceCollection *space.Collection) {
	c.spaceCollection = spaceCollection
}

// SetStorageCollection sets the storage collection for the config collection.
func (c *Collection) SetStorageCollection(storageCollection *storage.Collection) {
	c.storageCollection = storageCollection
}

// Space returns the space collection of the config collection.
func (c *Collection) Space() *space.Collection {
	return c.spaceCollection
}

// Storage returns the storage collection of the config collection.
func (c *Collection) Storage() *storage.Collection {
	return c.storageCollection
}
