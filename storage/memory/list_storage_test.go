package memory

import (
	"testing"
	"time"
)

func Test_ListStorage_PushToListPopFromList(t *testing.T) {
	newStorage := MustNew()
	defer newStorage.Shutdown()

	var err error
	err = newStorage.PushToList("key", "element1")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = newStorage.PushToList("key", "element2")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = newStorage.PushToList("key", "element3")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	var element string
	element, err = newStorage.PopFromList("key")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element != "element1" {
		t.Fatal("expected", "element1", "got", element)
	}
	element, err = newStorage.PopFromList("key")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element != "element2" {
		t.Fatal("expected", "element2", "got", element)
	}
	element, err = newStorage.PopFromList("key")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element != "element3" {
		t.Fatal("expected", "element3", "got", element)
	}

	timeOut := make(chan struct{}, 1)
	go func() {
		// Fetching elements from a list removes the fetched elements from the list.
		// After all elements are fetched from the list, the list must be empty.
		element, err = newStorage.PopFromList("key")
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		if element != "" {
			t.Fatal("expected", "", "got", element)
		}
		timeOut <- struct{}{}
	}()

	select {
	case <-time.After(100 * time.Millisecond):
		// The test succeeded.
	case <-timeOut:
		t.Fatal("expected", "success", "got", "timeout")
	}
}
