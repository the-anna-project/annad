package id

import (
	"sync"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func testMaybeNewIDFactory(t *testing.T) spec.IDFactory {
	newConfig := DefaultFactoryConfig()
	newFactory, err := NewFactory(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newFactory
}

// Test_ID_001 checks that a generated ID is still unique after a certain
// number of generations.
func Test_IDFactory_WithType(t *testing.T) {
	newFactory := testMaybeNewIDFactory(t)

	alreadySeen := map[spec.ObjectID]struct{}{}

	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newObjectID, err := newFactory.WithType(Hex128)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}

			mutex.Lock()
			if _, ok := alreadySeen[newObjectID]; ok {
				t.Fatal("id.NewObjectID returned the same ID twice")
			}
			alreadySeen[newObjectID] = struct{}{}
			mutex.Unlock()
		}()
	}
	wg.Wait()
}
