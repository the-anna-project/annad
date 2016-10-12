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
