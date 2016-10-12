package memory

import (
	"testing"
)

func Test_SetStorage_PushGetAll(t *testing.T) {
	newStorage := MustNew()
	defer newStorage.Shutdown()

	var err error
	err = newStorage.PushToSet("key", "element1")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = newStorage.PushToSet("key", "element2")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = newStorage.PushToSet("key", "element3")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	elements, err := newStorage.GetAllFromSet("key")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if len(elements) != 3 {
		t.Fatal("expected", 3, "got", len(elements))
	}
	if elements[0] != "element1" {
		t.Fatal("expected", "element1", "got", elements[0])
	}
	if elements[1] != "element2" {
		t.Fatal("expected", "element2", "got", elements[1])
	}
	if elements[2] != "element3" {
		t.Fatal("expected", "element3", "got", elements[2])
	}

	// Fetching all elements from a set does not remove the fetched elements from
	// the set. Multiple calls to GetAllFromSet always must return the same
	// elements.
	elements, err = newStorage.GetAllFromSet("key")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if len(elements) != 3 {
		t.Fatal("expected", 3, "got", len(elements))
	}
}

func Test_SetStorage_WalkPushRemove(t *testing.T) {
	newStorage := MustNew()
	defer newStorage.Shutdown()

	// Check the set is empty by default
	var element1 string
	err := newStorage.WalkSet("test-key", nil, func(element string) error {
		element1 = element
		return nil
	})
	if element1 != "" {
		t.Fatal("expected", "", "got", element1)
	}

	// Check an element can be pushed to a set.
	err = newStorage.PushToSet("test-key", "test-value")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	var element2 string
	err = newStorage.WalkSet("test-key", nil, func(element string) error {
		element2 = element
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element2 != "test-value" {
		t.Fatal("expected", "test-value", "got", element2)
	}

	// Check an element can be removed from a set.
	err = newStorage.RemoveFromSet("test-key", "test-value")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	var element3 string
	err = newStorage.WalkSet("test-key", nil, func(element string) error {
		element3 = element
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element3 != "" {
		t.Fatal("expected", "", "got", element3)
	}
}

func Test_SetStorage_WalkSet_Closer(t *testing.T) {
	newStorage := MustNew()
	defer newStorage.Shutdown()

	err := newStorage.PushToSet("test-key", "test-value")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Immediately close the walk.
	closer := make(chan struct{}, 1)
	closer <- struct{}{}

	// Check that the walk does not happen, because we already ended it.
	var element1 string
	err = newStorage.WalkSet("test-key", closer, func(element string) error {
		element1 = element
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element1 != "" {
		t.Fatal("expected", "", "got", element1)
	}
}
