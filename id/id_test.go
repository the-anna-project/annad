package id_test

import (
	"sync"
	"testing"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/spec"
)

// Test_ID_001 checks that a generated ID is still unique after a certain
// number of generations.
func Test_ID_001(t *testing.T) {
	alreadySeen := map[spec.ObjectID]struct{}{}

	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newObjectID := id.NewObjectID(id.Hex128)

			if _, ok := alreadySeen[newObjectID]; ok {
				t.Fatal("id.NewObjectID returned the same ID twice")
			}
			mutex.Lock()
			alreadySeen[newObjectID] = struct{}{}
			mutex.Unlock()
		}()
	}
	wg.Wait()
}
