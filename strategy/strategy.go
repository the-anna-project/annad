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
	// CLGNames represents a list of ordered action items, that are CLG names.
	CLGNames []string

	// StringMap provides a way to create a new strategy object out of a given
	// hash map containing bare strategy data. If this is nil or empty, a
	// completely new strategy is created. Otherwise it is tried to create a new
	// strategy using the information of the given hash map.
	StringMap map[string]string

	// Requestor represents the object requesting a strategy. E.g. when the
	// character network requests a strategy to act on the given input, it will
	// instruct an impulse to go through the strategy network while being
	// configured with information about the character network. Here the
	// requestor would hold the object tyope of the character network.
	Requestor spec.ObjectType
}

// DefaultConfig provides a default configuration to create a new strategy
// object by best effort. Note that the list of CLG names is empty and needs to
// be properly set before the strategy creation.
func DefaultConfig() Config {
	newConfig := Config{
		CLGNames:  []string{},
		StringMap: nil,
		Requestor: "",
	}

	return newConfig
}

// NewStrategy creates a new configured strategy object.
func NewStrategy(config Config) (spec.Strategy, error) {
	var newStrategy *strategy
	var err error

	if config.StringMap != nil {
		newStrategy, err = stringMapToStrategy(config.StringMap)
		if err != nil {
			return nil, maskAnyf(invalidConfigError, err.Error())
		}
	} else {
		newStrategy = &strategy{
			Config: config,
			ID:     id.NewObjectID(id.Hex128),
			Type:   ObjectTypeStrategy,
		}
		// TODO test that CLG names are only randomized on new creation. Given
		// string maps MUST not be randomized again.
		newStrategy.CLGNames = randomizeCLGNames(newStrategy.CLGNames)
	}

	if len(newStrategy.CLGNames) == 0 {
		return nil, maskAnyf(invalidConfigError, "CLG names must not be empty")
	}
	if newStrategy.ID == "" {
		return nil, maskAnyf(invalidConfigError, "ID must not be empty")
	}
	if newStrategy.Requestor == "" {
		return nil, maskAnyf(invalidConfigError, "requestor not be empty")
	}

	return newStrategy, nil
}

type strategy struct {
	Config

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (s *strategy) GetCLGNames() []string {
	return s.CLGNames
}

func (s *strategy) GetRequestor() spec.ObjectType {
	return s.Requestor
}

func (s *strategy) GetStringMap() (map[string]string, error) {
	newStringMap, err := strategyToStringMap(s)
	if err != nil {
		return nil, maskAny(err)
	}

	return newStringMap, nil
}

const (
	// clgNameDummy is simply a dummy CLG name injected during randomization
	// of the action list. See documentations below for more information.
	clgNameDummy = "dummy"
)

// randomizeCLGNames generates a random sequence using the given CLG names.
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
func randomizeCLGNames(clgNames []string) []string {
	var newCLGNames []string
	if len(clgNames) == 0 {
		// In case there is no useful input given we simply return an empty list.
		// This also prevents a dead lock in the loops below.
		return newCLGNames
	}

	// The trick to randomize the given set of CLG names is to inject a well
	// known item that can be chosen and then ignored.
	options := append([]string{clgNameDummy}, clgNames...)

	for {
		for range clgNames {
			max := big.NewInt(int64(len(options)))
			i, err := rand.Int(rand.Reader, max)
			if err != nil {
				panic(err)
			}
			newOption := options[i.Int64()]

			if newOption == clgNameDummy {
				// There was a random index that chose the item we want to ignore. Thus
				// we do so. This results in combinations not necessarily having the same
				// length as the original given list of CLG names.
				continue
			}

			newCLGNames = append(newCLGNames, newOption)
		}

		if len(newCLGNames) == 0 {
			continue
		}

		break
	}

	return newCLGNames
}
