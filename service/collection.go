package service

import (
	"sync"

	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// NewCollection creates a new service collection.
func NewCollection() servicespec.Collection {
	return &collection{}
}

type collection struct {
	// Dependencies.

	activator   servicespec.Activator
	forwarder   servicespec.Forwarder
	fs          servicespec.FS
	id          servicespec.ID
	log         servicespec.Log
	network     servicespec.Network
	permutation servicespec.Permutation
	random      servicespec.Random
	textInput   servicespec.TextInput
	textOutput  servicespec.TextOutput
	tracker     servicespec.Tracker

	// Internals.

	metadata     map[string]string
	shutdownOnce sync.Once
}

func (c *collection) Configure() error {
	// Internals.

	id, err := c.ID().New()
	if err != nil {
		return maskAny(err)
	}
	c.metadata = map[string]string{
		"id":   id,
		"name": "collection",
		"type": "service",
	}

	c.shutdownOnce = sync.Once{}

	return nil
}

func (c *collection) Activator() servicespec.Activator {
	return c.activator
}

func (c *collection) Forwarder() servicespec.Forwarder {
	return c.forwarder
}

func (c *collection) FS() servicespec.FS {
	return c.fs
}

func (c *collection) ID() servicespec.ID {
	return c.id
}

func (c *collection) Log() servicespec.Log {
	return c.log
}

func (c *collection) Network() servicespec.Network {
	return c.network
}

func (c *collection) Metadata() map[string]string {
	return c.metadata
}

func (c *collection) Permutation() servicespec.Permutation {
	return c.permutation
}

func (c *collection) Random() servicespec.Random {
	return c.random
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

		wg.Wait()
	})
}

func (c *collection) SetActivator(a servicespec.Activator) {
	c.activator = a
}

func (c *collection) SetForwarder(f servicespec.Forwarder) {
	c.forwarder = f
}

func (c *collection) SetFS(fs servicespec.FS) {
	c.fs = fs
}

func (c *collection) SetID(id servicespec.ID) {
	c.id = id
}

func (c *collection) SetLog(l servicespec.Log) {
	c.log = l
}

func (c *collection) SetNetwork(n servicespec.Network) {
	c.network = n
}

func (c *collection) SetPermutation(p servicespec.Permutation) {
	c.permutation = p
}

func (c *collection) SetRandom(r servicespec.Random) {
	c.random = r
}

func (c *collection) SetTextInput(ti servicespec.TextInput) {
	c.textInput = ti
}

func (c *collection) SetTextOutput(to servicespec.TextOutput) {
	c.textOutput = to
}

func (c *collection) SetTracker(t servicespec.Tracker) {
	c.tracker = t
}

func (c *collection) TextInput() servicespec.TextInput {
	return c.textInput
}

func (c *collection) TextOutput() servicespec.TextOutput {
	return c.textOutput
}

func (c *collection) Tracker() servicespec.Tracker {
	return c.tracker
}

func (c *collection) Validate() error {
	// Dependencies.
	if c.activator == nil {
		return maskAnyf(invalidConfigError, "activator service must not be empty")
	}
	if c.forwarder == nil {
		return maskAnyf(invalidConfigError, "forwarder service must not be empty")
	}
	if c.id == nil {
		return maskAnyf(invalidConfigError, "ID service must not be empty")
	}
	if c.log == nil {
		return maskAnyf(invalidConfigError, "log service must not be empty")
	}
	if c.permutation == nil {
		return maskAnyf(invalidConfigError, "permutation service must not be empty")
	}
	if c.random == nil {
		return maskAnyf(invalidConfigError, "random service must not be empty")
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
