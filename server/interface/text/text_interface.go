// Package textinterface implements spec.TextInterface and provides a way to
// feed neural networks with text input. To make Anna consume text, there is
// the text interface implemented through the network API.
package textinterface

import (
	"sync"

	"time"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/scheduler"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeTextInterface represents the object type of the text interface
	// object. This is used e.g. to register itself to the logger.
	ObjectTypeTextInterface spec.ObjectType = "text-interface"
)

// Config represents the configuration used to create a new text interface
// object.
type Config struct {
	Log         spec.Log
	Scheduler   spec.Scheduler
	TextGateway spec.Gateway
}

// DefaultConfig provides a default configuration to create a new text
// interface object by best effort.
func DefaultConfig() Config {
	return Config{
		Log:         log.NewLog(log.DefaultConfig()),
		Scheduler:   nil,
		TextGateway: gateway.NewGateway(gateway.DefaultConfig()),
	}
}

// NewTextInterface creates a new configured text interface object.
func NewTextInterface(config Config) (spec.TextInterface, error) {
	newInterface := &textInterface{
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   spec.ObjectType(ObjectTypeTextInterface),
	}

	newInterface.Log.Register(newInterface.GetType())

	if newInterface.Scheduler == nil {
		return nil, maskAnyf(invalidConfigError, "scheduler must not be empty")
	}
	newInterface.Scheduler.Register("ReadPlainWithInputAction", newInterface.ReadPlainWithInputAction)

	return newInterface, nil
}

type textInterface struct {
	Config

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

// TODO this should actually fetch a url from the web
func (ti *textInterface) FetchURL(url string) ([]byte, error) {
	return nil, nil
}

// TODO this should actually read a file from file system
func (ti *textInterface) ReadFile(file string) ([]byte, error) {
	return nil, nil
}

// TODO this should actually be streamed
func (ti *textInterface) ReadStream(stream string) ([]byte, error) {
	return nil, nil
}

// return response
func (ti *textInterface) ReadPlainWithID(ctx context.Context, jobID string) (string, error) {
	ti.Log.WithTags(spec.Tags{L: "D", O: ti, T: nil, V: 13}, "call ReadPlainWithID")

	job, err := ti.Scheduler.WaitForFinalStatus(spec.ObjectID(jobID), ctx.Done())
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

// return jobID
func (ti *textInterface) ReadPlainWithInput(ctx context.Context, input, expected, sessionID string) (string, error) {
	ti.Log.WithTags(spec.Tags{L: "D", O: ti, T: nil, V: 13}, "call ReadPlainWithInput")

	newJobConfig := scheduler.DefaultJobConfig()
	newJobConfig.ActionID = "ReadPlainWithInputAction"
	newJobConfig.Args = readPlainWithInputArgs{
		Input:    input,
		Expected: expected,
	}
	newJobConfig.SessionID = sessionID
	newJob, err := scheduler.NewJob(newJobConfig)
	if err != nil {
		return "", maskAny(err)
	}

	err = ti.Scheduler.Execute(newJob)
	if err != nil {
		return "", maskAny(err)
	}

	return string(newJob.GetID()), nil
}

// readPlainWithInputArgs represents the arguments configured for and passed to
// ReadPlainWithInputAction.
type readPlainWithInputArgs struct {
	Input    string
	Expected string
}

// ReadPlainWithInputAction represents the action of a scheduler job being
// executed to process ReadPlainWithInput requests asynchronously. args is
// supposed to be of type readPlainWithInputArgs and represents the arguments
// passed to this action method. closer represents a notification channel
// signaling the cancelation of the current job. Thus it informs the action to
// stop.
func (ti *textInterface) ReadPlainWithInputAction(args interface{}, closer <-chan struct{}) (string, error) {
	input := args.(readPlainWithInputArgs).Input
	expected := args.(readPlainWithInputArgs).Expected

	newConfig := gateway.DefaultSignalConfig()
	newConfig.Input = input
	newSignal := gateway.NewSignal(newConfig)

	// Start processing the input. Note that we in all cases want to send a
	// signal with the given input at least ones to the neural networks,
	// regardless any cancelations through the closer. The closer is allowed to
	// end the work being done here in case the input was processed by the neural
	// networks at least one time.
	for {
		newSignal, err := ti.TextGateway.Send(newSignal, nil)
		if err != nil {
			return "", maskAny(err)
		}

		output := newSignal.GetOutput()
		o := output.(string)
		if expected == "" || (expected != "" && o == expected) {
			// When there is no expected output given, simply return what we got.
			// When there is expected output given and it matches what we got,
			// return it.
			return o, nil
		}

		select {
		case <-closer:
			// This action was closed by the scheduler itself. This happens e.g.
			// when the job's final status was manually set, or another job for the
			// same session ID was scheduled.
			return "", nil
		default:
			// We did not yet receive the signal to stop the work of this action. Go
			// ahead with the next iteration.
		}

		time.Sleep(1 * time.Second)
	}
}
