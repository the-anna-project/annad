package impulse

import (
	"sync"

	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

const (
	ObjectType spec.ObjectType = "impulse"
)

type Config struct {
	State spec.State `json:"state,omitempty"`
}

func DefaultConfig() Config {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = ObjectType

	newConfig := Config{
		State: state.NewState(newStateConfig),
	}

	return newConfig
}

func NewImpulse(config Config) spec.Impulse {
	newImpulse := &impulse{
		Config: config,
		Mutex:  sync.Mutex{},
	}

	return newImpulse
}

type impulse struct {
	Config

	Mutex sync.Mutex `json:"mutex,omitempty"`
}

func (i *impulse) GetObjectID() spec.ObjectID {
	return i.GetState().GetObjectID()
}

func (i *impulse) GetObjectType() spec.ObjectType {
	return i.GetState().GetObjectType()
}

func (i *impulse) GetState() spec.State {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	return i.State
}

func (i *impulse) SetState(state spec.State) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	i.State = state
}

func (i *impulse) WalkThrough(neu spec.Neuron) (spec.Impulse, spec.Neuron, error) {
	var err error
	var imp spec.Impulse = i

	// Process the further walk through of the impulse dynamically. The loop is a
	// good alternative for real recursion.
	for {
		imp, neu, err = neu.Trigger(imp)
		if err != nil {
			return nil, nil, maskAny(err)
		}

		if neu == nil {
			// As soon as a neuron has decided to not forward an impulse to any
			// further neuron, the impulse went its way all through the whole
			// network. So we break here to return the impulse.
			break
		}
	}

	return imp, neu, nil
}
