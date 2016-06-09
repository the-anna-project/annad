package strategy

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func Test_Strategy_JSON_overwrite(t *testing.T) {
	firstStrategy := testMaybeNewStrategy(t)

	b, err := json.Marshal(firstStrategy)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newConfig := DefaultConfig()
	newConfig.Root = spec.CLG("other")
	secondStrategy, err := New(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = json.Unmarshal(b, secondStrategy)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if !reflect.DeepEqual(firstStrategy, secondStrategy) {
		t.Fatal("expected", false, "got", true)
	}
}

func Test_Strategy_JSON_empty(t *testing.T) {
	firstStrategy := testMaybeNewStrategy(t)

	b, err := json.Marshal(firstStrategy)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	secondStrategy := NewEmpty()
	err = json.Unmarshal(b, secondStrategy)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if !reflect.DeepEqual(firstStrategy, secondStrategy) {
		t.Fatal("expected", false, "got", true)
	}
}
