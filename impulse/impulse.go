package impulse

import (
	"sync"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

const (
	ObjectType spec.ObjectType = "impulse"
)

type Config struct {
	Log spec.Log `json:"-"`

	States map[string]spec.State `json:"states,omitempty"`
}

func DefaultConfig() Config {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = ObjectType

	newConfig := Config{
		Log: log.NewLog(log.DefaultConfig()),
		States: map[string]spec.State{
			common.DefaultStateKey: state.NewState(newStateConfig),
		},
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

func (i *impulse) Copy() spec.Impulse {
	impulseCopy := *i

	for key, state := range impulseCopy.States {
		impulseCopy.States[key] = state.Copy()
	}

	return &impulseCopy
}

func (i *impulse) GetObjectID() spec.ObjectID {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	return i.States[common.DefaultStateKey].GetObjectID()
}

func (i *impulse) GetObjectType() spec.ObjectType {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	return i.States[common.DefaultStateKey].GetObjectType()
}

func (i *impulse) GetState(key string) (spec.State, error) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	if state, ok := i.States[key]; ok {
		return state, nil
	}

	return nil, maskAny(stateNotFoundError)
}

func (i *impulse) SetState(key string, state spec.State) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	i.States[key] = state
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
