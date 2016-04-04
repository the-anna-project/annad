// Package clg implementes fundamental actions used to create strategies that
// allow to discover new behavior for problem solving.
package clg

import (
	"reflect"
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeIndex represents the object type of the index object. This is
	// used e.g. to register itself to the logger.
	ObjectTypeIndex spec.ObjectType = "index"
)

// Config represents the configuration used to create a new index object.
type Config struct {
	// Dependencies.
	Log spec.Log
}

// DefaultConfig provides a default configuration to create a new index object
// by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		Log: log.NewLog(log.DefaultConfig()),
	}

	return newConfig
}

// NewIndex creates a new configured index object.
func NewIndex(config Config) (spec.Index, error) {
	newIndex := &index{
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeIndex,
	}

	newIndex.Log.Register(newIndex.GetType())

	return newIndex, nil
}

type index struct {
	Config

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (i *index) Call(methodName string, args ...interface{}) ([]interface{}, error) {
	inputValues := ArgsToValues(args)
	outputValues := reflect.ValueOf(i).MethodByName(methodName).Call(inputValues)

	results, err := ValuesToArgs(outputValues)
	if err != nil {
		return nil, maskAny(err)
	}

	return results, nil
}
