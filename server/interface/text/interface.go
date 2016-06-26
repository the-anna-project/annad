package text

import (
	"sync"

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

// tinterface is not named interface because this is a reserved key in golang.
type tinterface struct {
	InterfaceConfig

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (i *tinterface) GetResponseForID(ctx context.Context, jobID string) (string, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call GetResponseForID")

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
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call ReadCoreRequest")

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
