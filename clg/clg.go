// Package clg implementes fundamental actions used to create strategies that
// allow to discover new behavior for problem solving.
package clg

import (
	"reflect"
	"strings"
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeCLGIndex represents the object type of the CLG index object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeCLGIndex spec.ObjectType = "clg-index"
)

// Config represents the configuration used to create a new CLG index object.
type Config struct {
	// Dependencies.
	Log spec.Log
}

// DefaultConfig provides a default configuration to create a new CLG index
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		Log: log.NewLog(log.DefaultConfig()),
	}

	return newConfig
}

// NewCLGIndex creates a new configured CLG index object.
func NewCLGIndex(config Config) (spec.CLGIndex, error) {
	newCLGIndex := &clgIndex{
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeCLGIndex,
	}

	newCLGIndex.Log.Register(newCLGIndex.GetType())

	return newCLGIndex, nil
}

type clgIndex struct {
	Config

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (i *clgIndex) CallCLGByName(args ...interface{}) ([]interface{}, error) {
	methodName, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}

	inputValues := ArgsToValues(args[1:])
	methodValue := reflect.ValueOf(i).MethodByName(methodName)
	if !methodValue.IsValid() {
		return nil, maskAnyf(methodNotFoundError, methodName)
	}

	outputValues := methodValue.Call(inputValues)
	results, err := ValuesToArgs(outputValues)
	if err != nil {
		return nil, maskAny(err)
	}

	return results, nil
}

func (i *clgIndex) GetCLGNames(args ...interface{}) ([]interface{}, error) {
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}
	var pattern string
	if len(args) == 1 {
		var err error
		pattern, err = ArgToString(args, 0)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	var allCLGNames []string

	t := reflect.TypeOf(i)
	for i := 0; i < t.NumMethod(); i++ {
		methodName := t.Method(i).Name
		if pattern != "" && !strings.Contains(methodName, pattern) {
			continue
		}
		allCLGNames = append(allCLGNames, methodName)
	}

	return []interface{}{allCLGNames}, nil
}
