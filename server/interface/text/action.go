package text

import (
	"time"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/gateway"
)

// ReadCoreRequestAction represents the action of a scheduler job being
// executed to process ReadPlainWithInput requests asynchronously. args is
// supposed to be of type api.CoreRequest and represents the arguments passed
// to this action method. closer represents a notification channel signaling
// the cancelation of the current job. Thus it informs the action to stop.
func (i *tinterface) ReadCoreRequestAction(args interface{}, closer <-chan struct{}) (string, error) {
	// Create a new signal to send it to the neural network.
	newConfig := gateway.DefaultSignalConfig()
	newConfig.Input = args.(api.CoreRequest)
	newSignal := gateway.NewSignal(newConfig)

	done := make(chan string, 1)
	fail := make(chan error, 1)

	go func() {
		// Start processing the input. We want to send a signal with the given input
		// in all cases to the neural network, regardless any cancelations through
		// the closer. The closer is allowed to end the work being done here in case
		// the input was processed by the neural network at least one time.
		newSignal, err := i.TextGateway.Send(newSignal, nil)
		if err != nil {
			fail <- maskAny(err)
		}

		done <- newSignal.GetOutput().(string)
	}()

	for {
		select {
		case <-closer:
			// This action was closed by the scheduler itself. This happens e.g.
			// when the job's final status was manually set, or another job for the
			// same session ID was scheduled.
			return "", nil
		case output := <-done:
			return output, nil
		case err := <-fail:
			return "", maskAny(err)
		default:
			// We did not yet receive the signal. Wait a little bit and go ahead with
			// the next iteration.
			time.Sleep(1 * time.Second)
		}
	}
}
