package service

import (
	"sync"

	"github.com/xh3b4sd/anna/service/spec"
)

// NewCollection creates a new service collection.
func NewCollection() spec.Collection {
	return &collection{}
}

type collection struct {
	// Dependencies.

	activator         spec.Activator
	feature           spec.Feature
	forwarder         spec.Forwarder
	fs                spec.FS
	id                spec.ID
	instrumentor      spec.Instrumentor
	log               spec.Log
	metricsEndpoint   spec.MetricsEndpoint
	network           spec.Network
	permutation       spec.Permutation
	random            spec.Random
	storageCollection spec.StorageCollection
	textEndpoint      spec.TextEndpoint
	textInput         spec.TextInput
	textOutput        spec.TextOutput
	tracker           spec.Tracker

	// Settings.

	metadata     map[string]string
	shutdownOnce sync.Once
}

func (c *collection) Configure() error {
	// Settings.

	id, err := c.ID().New()
	if err != nil {
		return maskAny(err)
	}
	c.metadata = map[string]string{
		"id":   id,
		"kind": "service",
		"name": "collection",
		"type": "service",
	}

	c.shutdownOnce = sync.Once{}

	return nil
}

func (c *collection) Activator() spec.Activator {
	return c.activator
}

func (c *collection) Feature() spec.Feature {
	return c.feature
}

func (c *collection) Forwarder() spec.Forwarder {
	return c.forwarder
}

func (c *collection) FS() spec.FS {
	return c.fs
}

func (c *collection) ID() spec.ID {
	return c.id
}

func (c *collection) Instrumentor() spec.Instrumentor {
	return c.instrumentor
}

func (c *collection) Log() spec.Log {
	return c.log
}

func (c *collection) Metadata() map[string]string {
	return c.metadata
}

func (c *collection) MetricsEndpoint() spec.MetricsEndpoint {
	return c.metricsEndpoint
}

func (c *collection) Network() spec.Network {
	return c.network
}

func (c *collection) Permutation() spec.Permutation {
	return c.permutation
}

func (c *collection) Random() spec.Random {
	return c.random
}

func (c *collection) Storage() spec.StorageCollection {
	return c.storageCollection
}

func (c *collection) SetActivator(a spec.Activator) {
	c.activator = a
}

func (c *collection) SetFeature(f spec.Feature) {
	c.feature = f
}

func (c *collection) SetForwarder(f spec.Forwarder) {
	c.forwarder = f
}

func (c *collection) SetFS(fs spec.FS) {
	c.fs = fs
}

func (c *collection) SetID(id spec.ID) {
	c.id = id
}

func (c *collection) SetInstrumentor(i spec.Instrumentor) {
	c.instrumentor = i
}

func (c *collection) SetLog(l spec.Log) {
	c.log = l
}

func (c *collection) SetMetricsEndpoint(s spec.MetricsEndpoint) {
	c.metricsEndpoint = s
}

func (c *collection) SetNetwork(n spec.Network) {
	c.network = n
}

func (c *collection) SetPermutation(p spec.Permutation) {
	c.permutation = p
}

func (c *collection) SetRandom(r spec.Random) {
	c.random = r
}

func (c *collection) SetStorageCollection(sc spec.StorageCollection) {
	c.storageCollection = sc
}

func (c *collection) SetTextEndpoint(te spec.TextEndpoint) {
	c.textEndpoint = te
}

func (c *collection) SetTextInput(ti spec.TextInput) {
	c.textInput = ti
}

func (c *collection) SetTextOutput(to spec.TextOutput) {
	c.textOutput = to
}

func (c *collection) SetTracker(t spec.Tracker) {
	c.tracker = t
}

func (c *collection) Shutdown() {
	c.shutdownOnce.Do(func() {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			c.Log().Line("msg", "shutting down network")
			c.Network().Shutdown()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			c.Log().Line("msg", "shutting down storage collection")
			c.Storage().Shutdown()
			wg.Done()
		}()

		wg.Wait()
	})
}

func (c *collection) TextEndpoint() spec.TextEndpoint {
	return c.textEndpoint
}

func (c *collection) TextInput() spec.TextInput {
	return c.textInput
}

func (c *collection) TextOutput() spec.TextOutput {
	return c.textOutput
}

func (c *collection) Tracker() spec.Tracker {
	return c.tracker
}

func (c *collection) Validate() error {
	// Dependencies.

	if c.activator == nil {
		return maskAnyf(invalidConfigError, "activator service must not be empty")
	}
	if c.feature == nil {
		return maskAnyf(invalidConfigError, "feature service must not be empty")
	}
	if c.forwarder == nil {
		return maskAnyf(invalidConfigError, "forwarder service must not be empty")
	}
	if c.id == nil {
		return maskAnyf(invalidConfigError, "ID service must not be empty")
	}
	if c.instrumentor == nil {
		return maskAnyf(invalidConfigError, "instrumentor service must not be empty")
	}
	if c.log == nil {
		return maskAnyf(invalidConfigError, "log service must not be empty")
	}
	if c.metricsEndpoint == nil {
		return maskAnyf(invalidConfigError, "metricsEndpoint service must not be empty")
	}
	if c.permutation == nil {
		return maskAnyf(invalidConfigError, "permutation service must not be empty")
	}
	if c.random == nil {
		return maskAnyf(invalidConfigError, "random service must not be empty")
	}
	if c.storageCollection == nil {
		return maskAnyf(invalidConfigError, "storage collection must not be empty")
	}
	if c.textEndpoint == nil {
		return maskAnyf(invalidConfigError, "text endpoint service must not be empty")
	}
	if c.textInput == nil {
		return maskAnyf(invalidConfigError, "text input service must not be empty")
	}
	if c.textOutput == nil {
		return maskAnyf(invalidConfigError, "text output service must not be empty")
	}
	if c.tracker == nil {
		return maskAnyf(invalidConfigError, "tracker service must not be empty")
	}

	return nil
}
