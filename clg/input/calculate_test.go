package input

import (
	"reflect"
	"testing"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

func Test_CLG_Divide(t *testing.T) {
	testCases := []struct {
		A        string
		Expected string
	}{
		{
			A:        "3.5",
			Expected: "3.5",
		},
		{
			A:        "foo",
			Expected: "foo",
		},
		{
			A:        "error",
			Expected: "error",
		},
		{
			A:        "this is a test input",
			Expected: "this is a test input",
		},
	}

	newCLG := MustNew()
	ctx := context.Background()

	for i, testCase := range testCases {
		newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
		newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(testCase.A)}
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
		if len(args) != 1 {
			t.Fatal("case", i+1, "expected", 1, "got", len(args))
		}
		result := args[0].String()

		if result != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", result)
		}
	}
}
