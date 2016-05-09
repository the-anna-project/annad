// Package clg implementes fundamental actions used to create strategies that
// allow to discover new behavior for problem solving.
package clg

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/index/clg/profile"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/worker-pool"
)

const (
	// ObjectTypeCLGIndex represents the object type of the CLG index object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeCLGIndex spec.ObjectType = "clg-index"
)

// IndexConfig represents the configuration used to create a new CLG index
// object.
type IndexConfig struct {
	// Dependencies.
	Generator spec.CLGProfileGenerator
	Log       spec.Log

	// Settings.
	NumGeneratorWorkers int
}

// DefaultIndexConfig provides a default configuration to create a new CLG
// index object by best effort.
func DefaultIndexConfig() IndexConfig {
	newGenerator, err := profile.NewGenerator(profile.DefaultGeneratorConfig())
	if err != nil {
		panic(err)
	}

	newConfig := IndexConfig{
		// Dependencies.
		Generator: newGenerator,
		Log:       log.NewLog(log.DefaultConfig()),

		// Settings.
		NumGeneratorWorkers: 10,
	}

	return newConfig
}

// NewIndex creates a new configured CLG index object.
func NewIndex(config IndexConfig) (spec.CLGIndex, error) {
	newIndex := &index{
		IndexConfig: config,

		BootOnce:     sync.Once{},
		Closer:       make(chan struct{}, 1),
		ID:           id.NewObjectID(id.Hex128),
		Mutex:        sync.Mutex{},
		Type:         ObjectTypeCLGIndex,
		ShutdownOnce: sync.Once{},
	}

	if newIndex.NumGeneratorWorkers < 1 {
		return nil, maskAnyf(invalidConfigError, "number of CLG profile generator workers must be greater than 0")
	}
	if newIndex.Generator == nil {
		return nil, maskAnyf(invalidConfigError, "CLG profile generator must not be empty")
	}

	newIndex.Log.Register(newIndex.GetType())

	return newIndex, nil
}

type index struct {
	IndexConfig

	BootOnce     sync.Once
	Closer       chan struct{}
	ID           spec.ObjectID
	Mutex        sync.Mutex
	Type         spec.ObjectType
	ShutdownOnce sync.Once
}

func (i *index) Boot() {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call Boot")

	i.BootOnce.Do(func() {
		// Create and/or update CLG profiles.
		go func() {
			err := i.CreateProfiles(i.GetGenerator())
			if err != nil {
				i.Log.WithTags(spec.Tags{L: "E", O: i, T: nil, V: 4}, "%#v", maskAny(err))
			}
		}()
	})
}

func (i *index) CreateProfiles(generator spec.CLGProfileGenerator) error {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call CreateProfiles")

	// Initialize queues we read from and write to.
	profileNames, err := generator.GetProfileNames()
	if err != nil {
		return maskAny(err)
	}
	profileNameQueue := make(chan string, len(profileNames))
	for _, pn := range profileNames {
		profileNameQueue <- pn
	}
	// We can close the queue channel immediately, because it only provides a one
	// way ticket and all writing is already done. As soon as a CLG name was
	// fetched from the queue it is considered WIP. A CLG name must never be
	// requeued.
	close(profileNameQueue)

	// Create worker function executed concurrently by all workers within the
	// worker pool we are going to create.
	workerFunc := func(canceler <-chan struct{}) error {
		for {
			select {
			case <-canceler:
				return maskAny(workerCanceledError)
			case pn := <-profileNameQueue:
				newProfile, hashChanged, err := generator.CreateProfile(pn)
				if err != nil {
					return maskAny(err)
				}
				if !hashChanged {
					// The created CLG profile is already known and did not change yet.
					// No need to update the stored version of it. Go ahead to create the
					// next one.
					continue
				}

				err = generator.StoreProfile(newProfile)
				if err != nil {
					return maskAny(err)
				}
			}
		}
	}

	// Prepare the worker pool.
	newWorkerPoolConfig := workerpool.DefaultConfig()
	newWorkerPoolConfig.WorkerFunc = workerFunc
	newWorkerPoolConfig.Canceler = i.Closer
	newWorkerPool, err := workerpool.NewWorkerPool(newWorkerPoolConfig)
	if err != nil {
		return maskAny(err)
	}

	// Execute the worker pool and block until all work is done.
	err = i.maybeReturnAndLogErrors(newWorkerPool.Execute())
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (i *index) GetGenerator() spec.CLGProfileGenerator {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call GetGenerator")

	return i.Generator
}

func (i *index) Shutdown() {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call Shutdown")

	i.ShutdownOnce.Do(func() {
		// Simply closing the closer will broadcast the signal to each listener.
		close(i.Closer)
	})
}
