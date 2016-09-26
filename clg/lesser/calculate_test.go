package lesser

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

func Test_CLG_Lesser(t *testing.T) {
	testCases := []struct {
		A        float64
		B        float64
		Expected float64
	}{
		{
			A:        3.5,
			B:        3.5,
			Expected: 3.5,
		},
		{
			A:        12.5,
			B:        3.5,
			Expected: 3.5,
		},
		{
			A:        14.5,
			B:        35.5,
			Expected: 14.5,
		},
		{
			A:        7.5,
			B:        -3.5,
			Expected: -3.5,
		},
		{
			A:        4.5,
			B:        12.5,
			Expected: 4.5,
		},
		{
			A:        65,
			B:        17,
			Expected: 17,
		},
		{
			A:        17,
			B:        65,
			Expected: 17,
		},
	}

	newCLG := MustNew()

	for i, testCase := range testCases {
		newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
		newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(testCase.A), reflect.ValueOf(testCase.B)}
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
