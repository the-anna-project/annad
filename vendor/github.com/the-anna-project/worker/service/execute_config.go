package service

import objectspec "github.com/the-anna-project/spec/object"

func newExecuteConfig() objectspec.WorkerExecuteConfig {
	return &executeConfig{
		// Settings.

		actions:       []func(canceler <-chan struct{}) error{},
		canceler:      nil,
		cancelOnError: true,
		numWorkers:    1,
	}
}

// executeConfig represents the configuration used to execute a new worker pool.
type executeConfig struct {
	// Settings.

	// actions represents the function executed by workers.
	actions []func(canceler <-chan struct{}) error
	// canceler can be used to end the worker pool's processes pro-actively. The
	// signal received here will be redirected to the canceler provided to the
	// worker functions.
	canceler chan struct{}
	// cancelOnError defines whether to signal cancelation of worker processes in
	// case one worker of the pool throws an error.
	cancelOnError bool
	// numWorkers represents the number of workers to be registered to run
	// concurrently within the pool.
	numWorkers int
}

func (ec *executeConfig) Actions() []func(canceler <-chan struct{}) error {
	return ec.actions
}

func (ec *executeConfig) Canceler() chan struct{} {
	return ec.canceler
}

func (ec *executeConfig) CancelOnError() bool {
	return ec.cancelOnError
}

func (ec *executeConfig) NumWorkers() int {
	return ec.numWorkers
}

func (ec *executeConfig) SetActions(actions []func(canceler <-chan struct{}) error) {
	ec.actions = actions
}

func (ec *executeConfig) SetCanceler(canceler chan struct{}) {
	ec.canceler = canceler
}

func (ec *executeConfig) SetCancelOnError(cancelOnError bool) {
	ec.cancelOnError = cancelOnError
}

func (ec *executeConfig) SetNumWorkers(numWorkers int) {
	ec.numWorkers = numWorkers
}
