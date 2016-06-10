package clg

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func testMaybeNewCollection(t *testing.T) Collection {
	return collection
}

func Test_CLG_Execute(t *testing.T) {
	name := spec.CLG("Sum")
	inputs := []reflect.Value{reflect.ValueOf(3.27), reflect.ValueOf(5.112)}

	outputs, err := Execute(name, inputs)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if len(outputs) != 1 {
		t.Fatal("expected", 1, "got", len(outputs))
	}

	if outputs[0].Float() != 8.382 {
		t.Fatal("expected", 8.382, "got", outputs[0])
	}
}

func Test_CLG_Execute_Error_MethodNotFound(t *testing.T) {
	name := spec.CLG("not found")

	_, err := Execute(name, nil)
	if !IsMethodNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_CLG_Inputs(t *testing.T) {
	name := spec.CLG("Sum")
	// Note it doesn't matter what floats we provide here, we only need a list of
	// two float64 reflect.Type.
	expected := []reflect.Type{reflect.TypeOf(3.27), reflect.TypeOf(5.112)}

	inputs, err := Inputs(name)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if !reflect.DeepEqual(inputs, expected) {
		t.Fatal("expected", false, "got", true)
	}
}

func Test_CLG_Inputs_Error_MethodNotFound(t *testing.T) {
	name := spec.CLG("not found")

	_, err := Inputs(name)
	if !IsMethodNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_CLG_Outputs(t *testing.T) {
	name := spec.CLG("Sum")
	// Note it doesn't matter what floats we provide here, we only need a list of
	// two float64 reflect.Type.
	expected := []reflect.Type{reflect.TypeOf(3.27)}

	outputs, err := Outputs(name)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if !reflect.DeepEqual(outputs, expected) {
		t.Fatal("expected", false, "got", true)
	}
}

func Test_CLG_Outputs_Error_MethodNotFound(t *testing.T) {
	name := spec.CLG("not found")

	_, err := Outputs(name)
	if !IsMethodNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_CLG_Names(t *testing.T) {
	expected := []spec.CLG{
		"Sum",
	}

	names := Names()

	for _, e := range expected {
		var found bool

		for _, n := range names {
			if n == e {
				found = true
			}
		}

		if !found {
			t.Fatal("expected", true, "got", false)
		}
	}
}
