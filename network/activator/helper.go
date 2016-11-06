package activator

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/xh3b4sd/anna/object/networkpayload"
	"github.com/xh3b4sd/anna/spec"
)

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, e := range a {
		if e != b[i] {
			return false
		}
	}

	return true
}

// getInputTypes returns a list of reflect types that represent the output
// arguments of a function. Therefore f must be a function. Otherwhise
// getInputTypes panics if f's reflect.Kind is not reflect.Func.
func getInputTypes(f interface{}) []reflect.Type {
	t := reflect.TypeOf(f)

	var inputType []reflect.Type

	for i := 0; i < t.NumIn(); i++ {
		inputType = append(inputType, t.In(i))
	}

	return inputType
}

func mergeNetworkPayloads(networkPayloads []spec.NetworkPayload) (spec.NetworkPayload, error) {
	if len(networkPayloads) == 0 {
		return nil, maskAny(networkPayloadNotFoundError)
	}

	var args []reflect.Value
	var sources []spec.ObjectID
	for _, np := range networkPayloads {
		for _, v := range np.GetArgs() {
			args = append(args, v)
		}

		sources = append(sources, np.GetSources()...)
	}

	ctx := networkPayloads[0].GetContext()
	behaviourID, ok := ctx.GetBehaviourID()
	if !ok {
		return nil, maskAnyf(invalidBehaviourIDError, "must not be empty")
	}

	networkPayloadConfig := networkpayload.DefaultConfig()
	networkPayloadConfig.Args = args
	networkPayloadConfig.Context = ctx
	networkPayloadConfig.Destination = spec.ObjectID(behaviourID)
	networkPayloadConfig.Sources = sources
	networkPayload, err := networkpayload.New(networkPayloadConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return networkPayload, nil
}

func queueToValues(queue []spec.NetworkPayload) []interface{} {
	var values []interface{}

	for _, p := range queue {
		values = append(values, p)
	}

	return values
}

func stringToQueue(s string) ([]spec.NetworkPayload, error) {
	var queue []spec.NetworkPayload

	for _, s := range strings.Split(s, ",") {
		np := networkpayload.MustNew()
		err := json.Unmarshal([]byte(s), &np)
		if err != nil {
			return nil, maskAny(err)
		}
		queue = append(queue, np)
	}

	return queue, nil
}

func typesToStrings(types []reflect.Type) []string {
	var strings []string

	for _, t := range types {
		strings = append(strings, t.String())
	}

	return strings
}

// valuesToQueue parses permutation values to network payloads. The underlying
// type of each network payload must be spec.NetworkPayload. Otherwhise
// valuesToQueue panics.
func valuesToQueue(values []interface{}) []spec.NetworkPayload {
	var queue []spec.NetworkPayload

	for _, v := range values {
		queue = append(queue, v.(spec.NetworkPayload))
	}

	return queue
}

// valuesToTypes parses permutation values to reflect types. The underlying type
// of each permutation value must be spec.NetworkPayload. The list of returned
// types will represent the reflect types of the given network payload's
// arguments. Thus, if the underlying type of the given values is not
// spec.NetworkPayload, valuesToTypes panics.
func valuesToTypes(values []interface{}) []reflect.Type {
	var types []reflect.Type

	for _, v := range values {
		for _, arg := range v.(spec.NetworkPayload).GetArgs() {
			types = append(types, arg.Type())
		}
	}

	return types
}
