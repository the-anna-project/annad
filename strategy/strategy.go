// Package strategy implements spec.Strategy to provide manageable action
// sequences.
package strategy

import (
	"crypto/rand"
	"math/big"
	"strings"
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeStrategy represents the object type of the strategy object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeStrategy = "strategy"
)

// Config represents the configuration used to create a new strategy object.
type Config struct {
	// Actions represents a list of ordered action items, that are object types.
	Actions []spec.ObjectType

	// HashMap provides a way to create a new strategy object out of a given hash
	// map containing bare strategy data. If this nil or empty, a completely new
	// strategy is created. Otherwise it is tried to create a new strategy using
	// the information of the given hash map.
	HashMap map[string]string

	// Requestor represents the object requesting a strategy. E.g. when the
	// character network requests a strategy to act on the given input, it will
	// instruct an impulse to go through the strategy network while being
	// configured with information about the character network. Here the
	// requestor would hold the object tyope of the character network.
	Requestor spec.ObjectType
}

// DefaultConfig provides a default configuration to create a new strategy
// object by best effort. Note that the list of actions is empty and needs to
// be properly set before the strategy creation.
func DefaultConfig() Config {
	newConfig := Config{
		Actions:   []spec.ObjectType{},
		HashMap:   nil,
		Requestor: "",
	}

	return newConfig
}

// NewStrategy creates a new configured strategy object.
func NewStrategy(config Config) (spec.Strategy, error) {
	var newStrategy *strategy

	if config.HashMap != nil {
		newStrategy = &strategy{}

		for k, v := range config.HashMap {
			if k == "actions" {
				var newActions []spec.ObjectType
				for _, a := range strings.Split(v, ",") {
					newActions = append(newActions, spec.ObjectType(a))
				}
				newStrategy.Actions = newActions
			}
			if k == "id" {
				newStrategy.ID = spec.ObjectID(v)
			}
			if k == "requestor" {
				newStrategy.Requestor = spec.ObjectType(v)
			}
		}
	} else {
		newStrategy = &strategy{
			Config: config,
			ID:     id.NewObjectID(id.Hex128),
		}
	}

	newStrategy.Type = ObjectTypeStrategy

	if len(newStrategy.Actions) == 0 {
		return nil, maskAnyf(invalidConfigError, "actions must not be empty")
	}
	if newStrategy.ID == "" {
		return nil, maskAnyf(invalidConfigError, "ID must not be empty")
	}
	if newStrategy.Requestor == "" {
		return nil, maskAnyf(invalidConfigError, "requestor not be empty")
	}

	newStrategy.Actions = randomizeActions(newStrategy.Actions)

	return newStrategy, nil
}

type strategy struct {
	Config

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (s *strategy) ActionsToString() string {
	str := ""
	actions := s.GetActions()

	for i, action := range actions {
		str += string(action)

		// When length of actions is 4, and in the last iteration i is 3, there
		// will be no more item to append. Thus we don't want to further append a
		// comma. So 3+1 is higher than 4-1, and we are save.
		if i+1 <= len(actions)-1 {
			str += ","
		}
	}

	return str
}

func (s *strategy) GetActions() []spec.ObjectType {
	return s.Actions
}

func (s *strategy) GetHashMap() map[string]string {
	hashMap := map[string]string{
		"actions":   s.ActionsToString(),
		"id":        string(s.GetID()),
		"requestor": string(s.GetRequestor()),
	}

	return hashMap
}

func (s *strategy) GetRequestor() spec.ObjectType {
	return s.Requestor
}

const (
	// objectTypeNone is simply a dummy object type injected during randomization
	// of the action list.
	objectTypeNone spec.ObjectType = "none"
)

// randomizeActions generates a random sequence using the given action items.
// Note that randomizing a strategy's action items MUST only be done when
// creating a new strategy. Further randomizations of existing strategies will
// cause the algorhythms the strategy network implements to fail.
//
// The following algorhythm is implemented as follows. Consider this given list
// of available action items.
//
//   a,b,c,d,e
//
// This are some possible combinations resulting out of the randomization.
//
//   c,e
//   b,b,d
//   a,b,a
//   d
//
func randomizeActions(actions []spec.ObjectType) []spec.ObjectType {
	newActions := []spec.ObjectType{}
	// The trick to randomize the given set of actions is to inject a well known
	// item that can be chosen and then ignored.
	options := append([]spec.ObjectType{objectTypeNone}, actions...)

	for {
		for range actions {
			max := big.NewInt(int64(len(options)))
			i, err := rand.Int(rand.Reader, max)
			if err != nil {
				panic(err)
			}
			newOption := options[i.Int64()]

			if newOption == objectTypeNone {
				// There was a random index that chose the item we want to ignore. Thus
				// we do so. This results in combinations not necessarily having the same
				// length as the original given list of actions.
				continue
			}

			newActions = append(newActions, newOption)
		}

		if len(newActions) == 0 {
			continue
		}

		break
	}

	return newActions
}
