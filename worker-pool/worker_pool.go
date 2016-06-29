// Package workerpool provides functionality to manage a pool of workers
// processing actions in parallel. This includes several features like
// signaling to cancel a workers process in case one worker throws an error.
package workerpool

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeWorkerPool represents the object type of the worker pool object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeWorkerPool spec.ObjectType = "worker-pool"
)

// Config represents the configuration used to create a new worker pool object.
type Config struct {
	// Dependencies.
	Log spec.Log

	// Settings.

	// Canceler can be used to end the worker pool's processes pro-actively. The
	// signal received here will be redirected to the canceler provided to the
	// worker functions.
	Canceler chan struct{}

	// CancelOnError defines whether to signal cancelation of worker processes in
	// case one worker of the pool throws an error.
	CancelOnError bool

	// NumWorkers represents the number of workers to be registered to run
	// concurrently within the pool.
	NumWorkers int

	// WorkerFunc represents the function executed by workers.
	WorkerFunc func(canceler <-chan struct{}) error
}

// DefaultConfig provides a default configuration to create a new worker pool
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		Log: log.NewLog(log.DefaultConfig()),

		// Settings.
		Canceler:      nil,
		CancelOnError: true,
		NumWorkers:    10,
		WorkerFunc:    nil,
	}

	return newConfig
}

// New creates a new configured worker pool object.
func New(config Config) (spec.WorkerPool, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		panic(err)
	}

	newWorkerPool := &workerPool{
		Config:      config,
		Errors:      make(chan error, 1),
		ExecuteOnce: sync.Once{},
		Drained:     make(chan struct{}, 1),
		ID:          newID,
		Mutex:       sync.Mutex{},
		Type:        ObjectTypeWorkerPool,
	}

	if newWorkerPool.NumWorkers < 1 {
		return nil, maskAnyf(invalidConfigError, "number of workers must be greater than 0")
	}
	if newWorkerPool.WorkerFunc == nil {
		return nil, maskAnyf(invalidConfigError, "worker function must not be empty")
	}

	newWorkerPool.Log.Register(newWorkerPool.GetType())

	return newWorkerPool, nil
}

type workerPool struct {
	Config

	Errors      chan error
	ExecuteOnce sync.Once

	// Drained can is used to wait until the worker pool is drained and no work
	// is to be done anymore.
	Drained chan struct{}

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (wp *workerPool) Execute() chan error {
	var errors chan error

	wp.ExecuteOnce.Do(func() {
		errors = wp.executeOnce()
	})

	return errors
}

func (wp *workerPool) Wait() {
	<-wp.Drained
}
