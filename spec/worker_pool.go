package spec

// WorkerPool represents a pool of workers, that are goroutines, doing work as
// configured. This eases the management and implementation of concurrent
// processing.
type WorkerPool interface {
	// Execute runs all workers configured for the current pool concurrently. The
	// call to Execute blocks until all workers within the pool have finished
	// their work.
	//
	// The returned error channel is never nil. In case of an error it will
	// always at least contain one error, depending on the worker pool's
	// configuration.
	//
	// Note that only the first call to execute blocks and executes workers due
	// to sync.Once.
	Execute() chan error

	Object
}
