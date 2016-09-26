package round

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

func Test_CLG_Round(t *testing.T) {
	testCases := []struct {
		Float     float64
		Precision int
		Expected  float64
	}{
		{
			Float:     3.5,
			Precision: 0,
			Expected:  4,
		},
		{
			Float:     3.4,
			Precision: 0,
			Expected:  3,
		},
		{
			Float:     3.4,
			Precision: 1,
			Expected:  3.4,
		},
		{
			Float:     3.4,
			Precision: 2,
			Expected:  3.4,
		},
		{
			Float:     3.476,
			Precision: 2,
			Expected:  3.48,
		},
		{
			Float:     -3.476,
			Precision: 2,
			Expected:  -3.48,
		},
		{
			Float:     3,
			Precision: 0,
			Expected:  3,
		},
		{
			Float:     3,
			Precision: 2,
			Expected:  3,
		},
	}

	newCLG := MustNew()

	for i, testCase := range testCases {
		newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
		newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(testCase.Float), reflect.ValueOf(testCase.Precision)}
		newNetworkPayloadConfig.Destination = "destination"
		newNetworkPayloadConfig.Sources = []spec.ObjectID{"source"}
		newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
		if err != nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}

		calculatedNetworkPayload, err := newCLG.Calculate(newNetworkPayload)
		if err != nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}
		args := calculatedNetworkPayload.GetArgs()
		if len(args) != 2 {
			t.Fatal("case", i+1, "expected", 2, "got", len(args))
		}
		result := args[1].Float()

		if result != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", result)
		}
	}
}
