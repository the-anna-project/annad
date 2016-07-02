package findconnections

import (
	"encoding/json"

	"github.com/spf13/cast"

	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/worker-pool"
)

// clg

func (c *clg) Calculate(imp spec.Impulse, stage int, inputs []reflect.Value) (spec.Impulse, []spec.Strategy, error) {
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
	// storage
	// asynch
	// contextual scope
	// stage, input type, input value
	peers, err := findPeers(c, peerIDs)
	if err != nil {
		return nil, maskAny(err)
	}

	// simply return highest weighted top ten?

	return imp, peers, nil
}

// helper

func getPeerIDs(stage int, inputs []reflect.Value) (chan string, error) {
	cap := 1 + 2*len(inputs)
	peerIDs := make(chan string, cap)

	// Convert stage value to string.
	peerIDs = append(peerIDs, "stage:"+cast.ToString(stage))

	// Convert input types to strings.
	for _, v := range inputs {
		peerIDs = append(peerIDs, "input-type:"+v.Type().String())
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
		peerIDs = append(peerIDs, "input-value:"+casted)
	}

	return peerIDs, nil
}

func findPeers(c *Collection, peerIDs chan string) ([]spec.Strategy, error) {
	pipeline := make(chan spec.Strategy, cap(peerIDs))
	defer close(pipeline)

	// Synchronize the asynchronous pipeline. The peers list will be filled with
	// all peers coming through the pipeline.
	var peers []spec.Strategy
	go func() {
		for {
			select {
			case p := <-pipeline:
				peers = append(peers, p)
			}
		}
	}()

	// Collect all peers that are might be connected to the current scope.
	workerFunc := func(canceler <-chan struct{}) error {
		for {
			select {
			case <-canceler:
				return nil
			case peerID := <-peerIDs:
				// Here we lookup the stored peers. "o" represents the object used in
				// the key for identification and scoping. This initiating instance is
				// the CLG FindConnections.
				o := "FindConnections"
				value, err := c.Storage.Get(key.NewCLGKey(o, peerID))
				if err != nil {
					return maskAny(err)
				}

				var peer spec.Strategy
				err := json.Unmarshal([]byte(value), &peer)
				if err != nil {
					return maskAny(err)
				}

				pipeline <- peer
			}
		}
	}

	// Create a new worker pool configured with the new worker function.
	newWorkerPoolConfig := workerpool.DefaultConfig()
	newWorkerPoolConfig.WorkerFunc = workerFunc
	newWorkerPool, err := workerpool.New(newWorkerPoolConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	// Execute the worker pool and wait for it to be finished.
	err := maybeReturnAndLogErrors(c, newWorkerPool.Execute())
	if err != nil {
		return nil, maskAny(err)
	}

	return peers, nil
}

func maybeReturnAndLogErrors(c *Collection, errors chan error) error {
	var executeErr error

	for err := range errors {
		if executeErr == nil {
			// Only return the first error.
			executeErr = err
		} else {
			// Log all errors but the first one
			c.Log.WithTags(spec.Tags{L: "E", O: c, T: nil, V: 4}, "%#v", maskAny(err))
		}
	}

	if executeErr != nil {
		return maskAny(executeErr)
	}

	return nil
}
