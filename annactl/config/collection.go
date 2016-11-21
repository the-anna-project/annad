package config

import (
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/xh3b4sd/anna/annactl/config/config"
	"github.com/xh3b4sd/anna/annactl/config/endpoint"
	"github.com/xh3b4sd/anna/annactl/config/endpoint/text"
	"github.com/xh3b4sd/anna/annactl/config/session"
)

// NewCollection creates a new config collection. It provides configuration for
// the whole neural network.
func NewCollection() *Collection {
	collection := &Collection{}

	collection.SetConfig(config.New())
	collection.SetEndpointCollection(endpoint.NewCollection())
	collection.SetSession(session.New())

	collection.Endpoint().SetText(text.New())

	return collection
}

// Collection represents the config collection.
type Collection struct {
	// Settings.

	configObject       *config.Object
	endpointCollection *endpoint.Collection
	sessionObject      *session.Object
}

// Config returns the config file config of the config collection.
func (c *Collection) Config() *config.Object {
	return c.configObject
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

// Session returns the session file config of the config collection.
func (c *Collection) Session() *session.Object {
	return c.sessionObject
}

// SetConfig sets the config file config for the config collection.
func (c *Collection) SetConfig(configObject *config.Object) {
	c.configObject = configObject
}

// SetEndpointCollection sets the endpoint collection for the config collection.
func (c *Collection) SetEndpointCollection(endpointCollection *endpoint.Collection) {
	c.endpointCollection = endpointCollection
}

// SetSession sets the session file config for the config collection.
func (c *Collection) SetSession(sessionObject *session.Object) {
	c.sessionObject = sessionObject
}
