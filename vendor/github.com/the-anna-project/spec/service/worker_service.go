package service

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// WorkerService implements a service to process work concurrently.
type WorkerService interface {
	Boot()
	// Execute processes all configured actions concurrently. The call to Execute
	// blocks until all goroutines within the worker pool have finished their
	// work.
	Execute(config objectspec.WorkerExecuteConfig) error
	// ExecuteConfig provides a default configuration for Execute.
	ExecuteConfig() objectspec.WorkerExecuteConfig
	Metadata() map[string]string
	Service() ServiceCollection
	SetServiceCollection(serviceCollection ServiceCollection)
}
