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

	activator          spec.Activator
	config             spec.Config
	connection         spec.Connection
	endpointCollection spec.EndpointCollection
	feature            spec.Feature
	forwarder          spec.Forwarder
	fs                 spec.FS
	id                 spec.ID
	instrumentor       spec.Instrumentor
	log                spec.Log
	network            spec.Network
	permutation        spec.Permutation
	random             spec.Random
	storageCollection  spec.StorageCollection
	textInput          spec.TextInput
	textOutput         spec.TextOutput
	tracker            spec.Tracker

	// Settings.

	shutdownOnce sync.Once
}

func (c *collection) Activator() spec.Activator {
	return c.activator
}

func (c *collection) Boot() {
	go c.Activator().Boot()
	go c.Config().Boot()
	go c.Connection().Boot()
	go c.Endpoint().Boot()
	go c.Feature().Boot()
	go c.Forwarder().Boot()
	go c.FS().Boot()
	go c.ID().Boot()
	go c.Instrumentor().Boot()
	go c.Log().Boot()
	go c.Network().Boot()
	go c.Permutation().Boot()
	go c.Random().Boot()
	go c.Storage().Boot()
	go c.TextInput().Boot()
	go c.TextOutput().Boot()
	go c.Tracker().Boot()
}

func (c *collection) Config() spec.Config {
	return c.config
}

func (c *collection) Connection() spec.Connection {
	return c.connection
}

func (c *collection) Endpoint() spec.EndpointCollection {
	return c.endpointCollection
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

func (c *collection) Network() spec.Network {
	return c.network
}

func (c *collection) Permutation() spec.Permutation {
	return c.permutation
}

func (c *collection) Random() spec.Random {
	return c.random
}

func (c *collection) SetActivator(a spec.Activator) {
	c.activator = a
}

func (c *collection) SetConfig(config spec.Config) {
	c.config = config
}

func (c *collection) SetConnection(conn spec.Connection) {
	c.connection = conn
}

func (c *collection) SetEndpointCollection(endpointCollection spec.EndpointCollection) {
	c.endpointCollection = endpointCollection
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
			c.Service().Endpoint().Shutdown()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			c.Network().Shutdown()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			c.Storage().Shutdown()
			wg.Done()
		}()

		wg.Wait()
	})
}

func (c *collection) Storage() spec.StorageCollection {
	return c.storageCollection
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
