// FactoryServer implements spec.Factory and provides a centralized service for
// general object creation accessable via a gateway.
package factoryserver

import (
	"sync"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/core"
	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/file-system/fake"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/impulse"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/network/strategy"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
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
		FileSystem:     filesystemfake.NewFileSystem(),
		Log:            log.NewLog(log.DefaultConfig()),
		TextGateway:    gateway.NewGateway(gateway.DefaultConfig()),

		// Settings.
		RedisAddr: "127.0.0.1:6379",
	}

	return newConfig
}

func NewFactory(config Config) spec.Factory {
	newFactory := &factoryServer{
		Closed: false,
		Closer: make(chan struct{}, 1),
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   common.ObjectType.FactoryServer,
	}

	return newFactory
}

type factoryServer struct {
	Closed bool
	Closer chan struct{}

	Config

	ID spec.ObjectID

	Mutex sync.Mutex

	Type spec.ObjectType
}

func (fs *factoryServer) Boot() {
	fs.Log.WithTags(spec.Tags{L: "D", O: fs, T: nil, V: 15}, "call Boot")

	go fs.FactoryGateway.Listen(fs.gatewayListener, fs.Closer)
}

func (fs *factoryServer) NewCore() (spec.Core, error) {
	fs.Log.WithTags(spec.Tags{L: "D", O: fs, T: nil, V: 15}, "call NewCore")

	newConfig := core.DefaultConfig()
	newConfig.FactoryClient = fs.FactoryClient
	newConfig.Log = fs.Log
	newConfig.TextGateway = fs.TextGateway
	newCore := core.NewCore(newConfig)

	return newCore, nil
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

func (fs *factoryServer) NewRedisStorage() (spec.Storage, error) {
	fs.Log.WithTags(spec.Tags{L: "D", O: fs, T: nil, V: 15}, "call NewRedisStorage")

	newDialConfig := storage.DefaultRedisDialConfig()
	newDialConfig.Addr = fs.RedisAddr

	newPoolConfig := storage.DefaultRedisPoolConfig()
	newPoolConfig.Dial = storage.NewRedisDial(newDialConfig)
	newPool := storage.NewRedisPool(newPoolConfig)

	newStorageConfig := storage.DefaultRedisStorageConfig()
	newStorageConfig.Log = fs.Log
	newStorageConfig.Pool = newPool

	newStorage := storage.NewRedisStorage(newStorageConfig)

	return newStorage, nil
}

func (fs *factoryServer) NewStrategyNetwork() (spec.Network, error) {
	fs.Log.WithTags(spec.Tags{L: "D", O: fs, T: nil, V: 15}, "call NewStrategyNetwork")

	newConfig := strategynetwork.DefaultNetworkConfig()
	newConfig.Log = fs.Log
	newNetwork, err := strategynetwork.NewNetwork(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newNetwork, nil
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
