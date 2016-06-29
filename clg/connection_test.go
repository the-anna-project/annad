package clg

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func Test_CLG_FindConnection(t *testing.T) {
	newCollection := testMaybeNewCollection(t)

	stage := 2
	inputs := []reflect.Value{
		reflect.ValueOf("foo"),
		reflect.ValueOf("foo"),
	}

	connections, err := newCollection.FindConnections(stage, inputs)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}
