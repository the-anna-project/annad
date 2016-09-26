package isbetween

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

func Test_CLG_IsBetween(t *testing.T) {
	testCases := []struct {
		N        float64
		Min      float64
		Max      float64
		Expected bool
	}{
		{
			N:        1,
			Min:      2,
			Max:      4,
			Expected: false,
		},
		{
			N:        2,
			Min:      2,
			Max:      4,
			Expected: true,
		},
		{
			N:        3,
			Min:      2,
			Max:      4,
			Expected: true,
		},
		{
			N:        4,
			Min:      2,
			Max:      4,
			Expected: true,
		},
		{
			N:        5,
			Min:      2,
			Max:      4,
			Expected: false,
		},
		{
			N:        35,
			Min:      -13,
			Max:      518,
			Expected: true,
		},
		{
			N:        -87,
			Min:      -413,
			Max:      -18,
			Expected: true,
		},
		{
			N:        -7,
			Min:      -413,
			Max:      -18,
			Expected: false,
		},
		{
			N:        -987,
			Min:      -413,
			Max:      -18,
			Expected: false,
		},
		{
			N:        1.8,
			Min:      2.34,
			Max:      4.944,
			Expected: false,
		},
		{
			N:        2.334,
			Min:      2.2,
			Max:      4.1,
			Expected: true,
		},
		{
			N:        3.9,
			Min:      2.003,
			Max:      4,
			Expected: true,
		},
		{
			N:        4,
			Min:      2.22,
			Max:      4.83,
			Expected: true,
		},
	}

	newCLG := MustNew()

	for i, testCase := range testCases {
		newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
		newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(testCase.N), reflect.ValueOf(testCase.Min), reflect.ValueOf(testCase.Max)}
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
