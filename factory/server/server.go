// FactoryServer implements spec.Factory and provides a centralized service for
// general object creation accessable via a gateway.
package factoryserver

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/file-system/memory"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/impulse"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	ObjectTypeFactoryServer spec.ObjectType = "factory-server"
)

type Config struct {
	// Dependencies.
	FactoryClient  spec.Factory
	FactoryGateway spec.Gateway
	FileSystem     spec.FileSystem
	Log            spec.Log
	TextGateway    spec.Gateway

	// Settings.
	RedisAddr string
}

func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		FactoryClient:  factoryclient.NewFactory(factoryclient.DefaultConfig()),
		FactoryGateway: gateway.NewGateway(gateway.DefaultConfig()),
		FileSystem:     memoryfilesystem.NewFileSystem(),
		Log:            log.NewLog(log.DefaultConfig()),
		TextGateway:    gateway.NewGateway(gateway.DefaultConfig()),

		// Settings.
		RedisAddr: "127.0.0.1:6379",
	}

	return newConfig
}

func NewFactory(config Config) spec.Factory {
	newFactory := &factoryServer{
		Booted: false,
		Closed: false,
		Closer: make(chan struct{}, 1),
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeFactoryServer,
	}

	newFactory.Log.Register(newFactory.GetType())

	return newFactory
}

type factoryServer struct {
	Config

	Booted bool
	Closed bool
	Closer chan struct{}
	ID     spec.ObjectID
	Mutex  sync.Mutex
	Type   spec.ObjectType
}

func (fs *factoryServer) Boot() {
	fs.Mutex.Lock()
	defer fs.Mutex.Unlock()

	if fs.Booted {
		return
	}
	fs.Booted = true

	fs.Log.WithTags(spec.Tags{L: "D", O: fs, T: nil, V: 15}, "call Boot")

	go fs.FactoryGateway.Listen(fs.gatewayListener, fs.Closer)
}

func (fs *factoryServer) NewImpulse() (spec.Impulse, error) {
	fs.Log.WithTags(spec.Tags{L: "D", O: fs, T: nil, V: 15}, "call NewImpulse")

	newConfig := impulse.DefaultConfig()
	newConfig.Log = fs.Log
	newImpulse, err := impulse.NewImpulse(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newImpulse, nil
}

func (fs *factoryServer) Shutdown() {
	fs.Log.WithTags(spec.Tags{L: "D", O: fs, T: nil, V: 15}, "call Shutdown")

	fs.Mutex.Lock()
	defer fs.Mutex.Unlock()

	fs.TextGateway.Close()

	if !fs.Closed {
		fs.Closer <- struct{}{}
		fs.Closed = true
	}
}
