package islesser

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

func Test_CLG_IsLesser(t *testing.T) {
	testCases := []struct {
		A        float64
		B        float64
		Expected bool
	}{
		{
			A:        3.5,
			B:        3.5,
			Expected: false,
		},
		{
			A:        12.5,
			B:        3.5,
			Expected: false,
		},
		{
			A:        14.5,
			B:        35.5,
			Expected: true,
		},
		{
			A:        7.5,
			B:        -3.5,
			Expected: false,
		},
		{
			A:        4.5,
			B:        12.5,
			Expected: true,
		},
		{
			A:        65,
			B:        17,
			Expected: false,
		},
		{
			A:        17,
			B:        65,
			Expected: true,
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
		result := args[1].Bool()

		if result != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", result)
		}
	}
}
