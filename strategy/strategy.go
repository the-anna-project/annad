// Package strategy implements spec.Strategy to provide manageable action
// sequences.
package strategy

import (
	"crypto/rand"
	"math/big"
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

	// ID represents the strategy's ID. This is configurable because the strategy
	// object is a container for strategy information. So even the ID needs to be
	// configured when e.g. fetching strategy information from a database and
	// creating a new object to carry around the fetched data.
	ID spec.ObjectID
}

// DefaultConfig provides a default configuration to create a new strategy
// object by best effort. Note that the list of actions is empty and needs to
// be properly set before the strategy creation.
func DefaultConfig() Config {
	newConfig := Config{
		Actions: []spec.ObjectType{},
		ID:      id.NewObjectID(id.Hex128),
	}

	return newConfig
}

// NewStrategy creates a new configured strategy object.
func NewStrategy(config Config) (spec.Strategy, error) {
	newStrategy := &strategy{
		Config: config,
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeStrategy,
	}

	if len(newStrategy.Actions) == 0 {
		return nil, maskAnyf(invalidActionsError, "must not be empty")
	}

	newStrategy.Actions = randomizeActions(newStrategy.Actions)

	return newStrategy, nil
}

type strategy struct {
	Config
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (s *strategy) GetActions() []spec.ObjectType {
	return s.Actions
}

func (s *strategy) String() string {
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
