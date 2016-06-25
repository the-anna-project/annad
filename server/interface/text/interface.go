package text

import (
	"sync"

	"time"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/scheduler"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeTextInterface represents the object type of the text interface
	// object. This is used e.g. to register itself to the logger.
	ObjectTypeTextInterface spec.ObjectType = "text-interface"
)

// InterfaceConfig represents the configuration used to create a new text
// interface object.
type InterfaceConfig struct {
	Log         spec.Log
	Scheduler   spec.Scheduler
	TextGateway spec.Gateway
}

// DefaultInterfaceConfig provides a default configuration to create a new text
// interface object by best effort.
func DefaultInterfaceConfig() InterfaceConfig {
	newScheduler, err := scheduler.NewScheduler(scheduler.DefaultConfig())
	if err != nil {
		panic(err)
	}

	newConfig := InterfaceConfig{
		Log:         log.NewLog(log.DefaultConfig()),
		Scheduler:   newScheduler,
		TextGateway: gateway.NewGateway(gateway.DefaultConfig()),
	}

	return newConfig
}

// NewInterface creates a new configured text interface object.
func NewInterface(config InterfaceConfig) (spec.TextInterface, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		return nil, maskAny(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		return nil, maskAny(err)
	}

	newInterface := &tinterface{
		InterfaceConfig: config,
		ID:              newID,
		Mutex:           sync.Mutex{},
		Type:            spec.ObjectType(ObjectTypeTextInterface),
	}

	newInterface.Log.Register(newInterface.GetType())

	if newInterface.Scheduler == nil {
		return nil, maskAnyf(invalidConfigError, "scheduler must not be empty")
	}
	newInterface.Scheduler.Register("ReadCoreRequestAction", newInterface.ReadCoreRequestAction)

	return newInterface, nil
}

type tinterface struct {
	InterfaceConfig

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (i *tinterface) GetResponseForID(ctx context.Context, jobID string) (string, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call ReadPlainWithID")

	job, err := i.Scheduler.WaitForFinalStatus(spec.ObjectID(jobID), ctx.Done())
	if err != nil {
		return "", maskAny(err)
	}

	if job == nil {
		// This should only happen in case the request was ended by ctx.Done().
		return "", nil
	}
	result := job.GetResult()

	return result, nil
}

func (i *tinterface) ReadCoreRequest(ctx context.Context, coreRequest api.CoreRequest, sessionID string) (string, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call ReadPlainWithInput")

	newJobConfig := scheduler.DefaultJobConfig()
	newJobConfig.ActionID = "ReadCoreRequestAction"
	newJobConfig.Args = coreRequest
	newJobConfig.SessionID = sessionID
	newJob, err := scheduler.NewJob(newJobConfig)
	if err != nil {
		return "", maskAny(err)
	}

	err = i.Scheduler.Execute(newJob)
	if err != nil {
		return "", maskAny(err)
	}

	return string(newJob.GetID()), nil
}

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
