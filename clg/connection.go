package clg

import (
	"github.com/spf13/cast"

	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/worker-pool"
)

// clg

// TODO
func (c *Collection) FindConnections(stage int, inputs []reflect.Value) ([]spec.Strategy, error) {
	if len(inputs) == 0 {
		return nil, maskAnyf(invalidCLGExecutionError, "inputs must not be empty")
	}

	// We need all string representations of the given inputs. These information
	// are used as peer IDs to lookup connections with other peers. The stage
	// number is a peer. Each input type is a peer. Each input value is a peer.
	// Note that peerIDs is of type chan string.
	peerIDs, err := getPeerIDs(stage, inputs)
	if err != nil {
		return nil, maskAny(err)
	}

	// TODO
	TODO, err := findPeers(c.Storage, peerIDs)
	if err != nil {
		return nil, maskAny(err)
	}

	// simply return highest weighted top ten?

	return nil, nil
}

// helper

func getPeerIDs(stage int, inputs []reflect.Value) (chan string, error) {
	cap := 1 + 2*len(inputs)
	peerIDs := make(chan string, cap)

	// Convert stage value to string.
	peerIDs = append(peerIDs, cast.ToString(stage))

	// Convert input types to strings.
	for _, v := range inputs {
		peerIDs = append(peerIDs, v.Type().String())
	}

	// Convert input values to strings.
	for _, v := range inputs {
		casted, err := cast.ToString(v.Interface())
		if err != nil {
			// TODO In case the cast library fails, it means it cannot convert the
			// given type into string. This might happen. We need to monitor which
			// types this errors cause and improve the convertion if necessary.
			continue
		}
		peerIDs = append(peerIDs, casted)
	}

	return peerIDs, nil
}

func findPeers(storage spec.Storage, peerIDs chan string) (TODO, error) {
	// Collect all peers that are might be connected to the current scope.
	workerFunc := func(canceler <-chan struct{}) error {
		// get peers pointing to
		//   - stage related strategies
		//   - input type related strategies
		//   - input value related strategies
	}

	// Create a new worker pool configured with the new worker function.
	newWorkerPoolConfig := workerpool.DefaultConfig()
	newWorkerPoolConfig.WorkerFunc = workerFunc
	newWorkerPool, err := workerpool.New(newWorkerPoolConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	// Execute the worker pool and wait for it to be finished.
	errors := newWorkerPool.Execute()
	newWorkerPool.Wait()

	// Handle all errors occured during the worker pool execution.
	for err := range errors {
		if IsWorkerCanceled(err) {
			continue
		}
		return nil, maskAny(err)
	}
}
