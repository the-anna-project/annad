package strategy

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_Strategy_Static_JSON_overwrite(t *testing.T) {
	firstStrategy := testMaybeNewStatic(t, nil)

	b, err := json.Marshal(firstStrategy)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newConfig := DefaultStaticConfig()
	newConfig.Argument = 55
	secondStrategy, err := NewStatic(newConfig)
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

func Test_Strategy_Static_JSON_empty(t *testing.T) {
	firstStrategy := testMaybeNewStatic(t, nil)

	b, err := json.Marshal(firstStrategy)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	secondStrategy := NewEmptyStatic()
	err = json.Unmarshal(b, secondStrategy)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if !reflect.DeepEqual(firstStrategy, secondStrategy) {
		t.Fatal("expected", false, "got", true)
	}
}
