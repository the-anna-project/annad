// Package clg implementes fundamental actions used to create strategies that
// allow to discover new behavior for problem solving.
//
// Note that this package defines a go generate statement to compile its own
// source code into the AssetFileSystem implementation. That way the function
// bodies of the implemented CLGs are available for inspection and hashing.
// Hashes of CLGs are used to check whether they changed. A change of a CLG
// affects its functionality, its profile and probably even its use case. Thus
// changes of the CLGs function bodies need to be detected to trigger profile
// updates.
//
//go:generate loader generate -p clg
//
package clg

import (
	"reflect"
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/worker-pool"
)

const (
	// ObjectTypeCLGIndex represents the object type of the CLG index object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeCLGIndex spec.ObjectType = "clg-index"
)

// Config represents the configuration used to create a new CLG index object.
type Config struct {
	// Dependencies.
	Collection spec.CLGCollection
	Log        spec.Log

	// Settings.
	NumCLGProfileWorkers int
}

// DefaultConfig provides a default configuration to create a new CLG index
// object by best effort.
func DefaultConfig() Config {
	newCLGCollection, err := NewCLGCollection(DefaultCLGCollectionConfig())
	if err != nil {
		panic(err)
	}

	newConfig := Config{
		// Dependencies.
		Collection: newCLGCollection,
		Log:        log.NewLog(log.DefaultConfig()),

		// Settings.
		NumCLGProfileWorkers: 10,
	}

	return newConfig
}

// NewCLGIndex creates a new configured CLG index object.
func NewCLGIndex(config Config) (spec.CLGIndex, error) {
	newCLGIndex := &clgIndex{
		Config: config,

		BootOnce:     sync.Once{},
		Closer:       make(chan struct{}, 1),
		ID:           id.NewObjectID(id.Hex128),
		Mutex:        sync.Mutex{},
		Type:         ObjectTypeCLGIndex,
		ShutdownOnce: sync.Once{},
	}

	if newCLGIndex.NumCLGProfileWorkers < 1 {
		return nil, maskAnyf(invalidConfigError, "number of CLG profile workers must be greater than 0")
	}

	newCLGIndex.Log.Register(newCLGIndex.GetType())

	return newCLGIndex, nil
}

type clgIndex struct {
	Config

	BootOnce     sync.Once
	Closer       chan struct{}
	ID           spec.ObjectID
	Mutex        sync.Mutex
	Type         spec.ObjectType
	ShutdownOnce sync.Once
}

func (i *clgIndex) Boot() {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call Boot")

	i.BootOnce.Do(func() {
		go func() {
			err := i.CreateCLGProfiles(i.GetCLGCollection())
			if err != nil {
				i.Log.WithTags(spec.Tags{L: "E", O: i, T: nil, V: 4}, "%#v", maskAny(err))
			}
		}()
	})
}

func (i *clgIndex) CreateCLGProfile(clgCollection spec.CLGCollection, clgName string, canceler <-chan struct{}) (spec.CLGProfile, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call CreateCLGProfile")

	// Fill a queue.
	args, err := clgCollection.GetNamesMethod()
	if err != nil {
		return nil, maskAny(err)
	}
	clgNames, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	queue := make(chan string, len(clgNames))
	for _, clgName := range clgNames {
		queue <- clgName
	}

	// Initialize the profile creation.
	for {
		select {
		case <-canceler:
			return nil, maskAny(workerCanceledError)
		case clgName := <-queue:
			methodValue := reflect.ValueOf(clgCollection).MethodByName(clgName)
			if !i.isMethodValue(methodValue) {
				return nil, maskAnyf(invalidCLGError, clgName)
			}

			var err error
			newCLGProfileConfig := DefaultCLGProfileConfig()
			newCLGProfileConfig.MethodName = clgName
			newCLGProfileConfig.MethodHash, err = i.getCLGMethodHash(methodValue)
			if err != nil {
				return nil, maskAny(err)
			}
			newCLGProfileConfig.InputTypes, err = i.getCLGInputTypes(methodValue)
			if err != nil {
				return nil, maskAny(err)
			}
			newCLGProfileConfig.InputExamples, err = i.getCLGInputExamples(methodValue)
			if err != nil {
				return nil, maskAny(err)
			}
			newCLGProfileConfig.RightSideNeighbours, err = i.getCLGRightSideNeighbours(clgCollection, clgName, methodValue, canceler)
			if err != nil {
				return nil, maskAny(err)
			}
			newCLGProfile, err := NewCLGProfile(newCLGProfileConfig)
			if err != nil {
				return nil, maskAny(err)
			}

			return newCLGProfile, nil
		}
	}
}

func (i *clgIndex) CreateCLGProfiles(clgCollection spec.CLGCollection) error {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call CreateCLGProfiles")

	// Fill a queue.
	args, err := clgCollection.GetNamesMethod()
	if err != nil {
		return maskAny(err)
	}
	clgNames, err := ArgToStringSlice(args, 0)
	if err != nil {
		return maskAny(err)
	}
	queue := make(chan string, len(clgNames))
	for _, clgName := range clgNames {
		queue <- clgName
	}
	// We can close the queue channel immediately, because it only provides a one
	// way ticket. As soon as a CLG name was fetched from the queue it is
	// considered WIP. A CLG name must never be requeued.
	close(queue)

	// Start N worker goroutines.
	workerFunc := func(canceler <-chan struct{}) error {
		for {
			select {
			case <-canceler:
				return maskAny(workerCanceledError)
			case clgName := <-queue:
				// Try to fetch the CLG profile in advance.
				currentCLGProfile, err := i.GetCLGProfileByName(clgName)
				if IsCLGProfileNotFound(err) {
					// In case the CLG profile cannot be found, we are going ahead to create
					// one.
				} else if err != nil {
					return maskAny(err)
				}

				newCLGProfile, err := i.CreateCLGProfile(clgCollection, clgName, canceler)
				if err != nil {
					return maskAny(err)
				}

				if currentCLGProfile.Equals(newCLGProfile) {
					// The CLG profile has not changed. Thus nothing to do here.
					continue
				}

				err = i.StoreCLGProfile(newCLGProfile)
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

func (i *clgIndex) GetCLGCollection() spec.CLGCollection {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call GetCLGCollection")

	return i.Collection
}

// TODO
func (i *clgIndex) GetCLGProfileByName(clgName string) (spec.CLGProfile, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call StoreCLGProfile")

	return nil, nil
}

func (i *clgIndex) Shutdown() {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call Shutdown")

	i.ShutdownOnce.Do(func() {
		// Simply closing the closer will broadcast the signal to each listener.
		close(i.Closer)
	})
}

// TODO
func (i *clgIndex) StoreCLGProfile(clgProfile spec.CLGProfile) error {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call StoreCLGProfile")

	return nil
}
