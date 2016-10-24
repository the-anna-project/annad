package memory

import (
	"sync"
	"testing"
)

func Test_StringStorage_GetRandom(t *testing.T) {
	newStorage := MustNew()
	defer newStorage.Shutdown()

	// We store 3 keys in each map of the memory storage to verify that we fetch a
	// random key across all stored keys.
	newStorage.Set("SetKey1", "SetValue1")
	newStorage.Set("SetKey2", "SetValue2")
	newStorage.Set("SetKey3", "SetValue3")
	newStorage.PushToSet("PushToSetKey1", "PushToSetElement1")
	newStorage.PushToSet("PushToSetKey2", "PushToSetElement2")
	newStorage.PushToSet("PushToSetKey3", "PushToSetElement3")
	newStorage.SetElementByScore("SetElementByScoreKey1", "SetElementByScoreElement1", 0)
	newStorage.SetElementByScore("SetElementByScoreKey2", "SetElementByScoreElement2", 0)
	newStorage.SetElementByScore("SetElementByScoreKey3", "SetElementByScoreElement3", 0)
	newStorage.SetStringMap("SetStringMapKey1", map[string]string{"foo": "bar"})
	newStorage.SetStringMap("SetStringMapKey2", map[string]string{"foo": "bar"})
	newStorage.SetStringMap("SetStringMapKey3", map[string]string{"foo": "bar"})

	alreadySeen := map[string]struct{}{}

	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newKey, err := newStorage.GetRandom()
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}

			mutex.Lock()
			alreadySeen[newKey] = struct{}{}
			mutex.Unlock()
		}()
	}
	wg.Wait()

	l := len(alreadySeen)
	if l != 12 {
		t.Fatal("expected", 12, "got", l)
	}
}

func Test_StringStorage_GetSetGet(t *testing.T) {
	newStorage := MustNew()
	defer newStorage.Shutdown()

	_, err := newStorage.Get("foo")
	if !IsNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}

	err = newStorage.Set("foo", "bar")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	value, err := newStorage.Get("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if value != "bar" {
		t.Fatal("expected", "bar", "got", value)
	}
}

func Test_StringStorage_Remove_NotFound(t *testing.T) {
	newStorage := MustNew()
	defer newStorage.Shutdown()

	err := newStorage.Remove("foo")
	if !IsNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringStorage_SetGetRemoveGet(t *testing.T) {
	newStorage := MustNew()
	defer newStorage.Shutdown()

	err := newStorage.Set("foo", "bar")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	value, err := newStorage.Get("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if value != "bar" {
		t.Fatal("expected", "bar", "got", value)
	}

	err = newStorage.Remove("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	_, err = newStorage.Get("foo")
	if !IsNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringStorage_WalkSetRemove(t *testing.T) {
	newStorage := MustNew()
	defer newStorage.Shutdown()

	// Verify the key space is empty by default.
	var count1 int
	err := newStorage.WalkKeys("*", nil, func(element string) error {
		count1++
		return nil
	})
	if count1 != 0 {
		t.Fatal("expected", 0, "got", count1)
	}

	// Set a new key.
	err = newStorage.Set("test-key", "test-value")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	var count2 int
	var element2 string
	err = newStorage.WalkKeys("*", nil, func(element string) error {
		count2++
		element2 = element
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if count2 != 1 {
		t.Fatal("expected", 1, "got", count2)
	}
	if element2 != "prefix:test-key" {
		t.Fatal("expected", "prefix:test-key", "got", element2)
	}

	// Remove one key.
	err = newStorage.Remove("test-key")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Verify there is now no key anymore.
	var count3 int
	err = newStorage.WalkKeys("*", nil, func(element string) error {
		count3++
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if count3 != 0 {
		t.Fatal("expected", 0, "got", count3)
	}
}

func Test_StringStorage_WalkKeys_Closer(t *testing.T) {
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
	err = newStorage.WalkKeys("test-key", closer, func(element string) error {
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
