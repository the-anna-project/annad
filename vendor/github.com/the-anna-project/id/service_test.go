package id

import (
	"io"
	"math/big"
	"sync"
	"testing"

	"github.com/the-anna-project/random"
	"github.com/the-anna-project/collection"
)

func Test_IDService_WithType_Error(t *testing.T) {
	serviceCollection := collection.New()
	serviceCollection.SetRandomService(random.New())
	serviceCollection.Random().SetRandFactory(func(randReader io.Reader, max *big.Int) (n *big.Int, err error) {
		return nil, maskAny(invalidConfigError)
	})

	idService := New()
	idService.SetServiceCollection(serviceCollection)

	_, err := idService.WithType(Hex128)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

// Test_IDService_New checks that a generated ID is still unique after a
// certain number of concurrent generations.
func Test_IDService_New(t *testing.T) {
	serviceCollection := collection.New()
	serviceCollection.SetRandomService(random.New())

	idService := New()
	idService.SetServiceCollection(serviceCollection)

	alreadySeen := map[string]struct{}{}

	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newObjectID, err := idService.New()
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}

			mutex.Lock()
			defer mutex.Unlock()
			if _, ok := alreadySeen[newObjectID]; ok {
				t.Fatal("idService.New returned the same ID twice")
			}
			alreadySeen[newObjectID] = struct{}{}
		}()
	}
	wg.Wait()
}

// Test_IDService_WithType checks that a generated ID is still unique after a
// certain number of concurrent generations.
func Test_IDService_WithType(t *testing.T) {
	serviceCollection := collection.New()
	serviceCollection.SetRandomService(random.New())

	idService := New()
	idService.SetServiceCollection(serviceCollection)

	alreadySeen := map[string]struct{}{}

	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newObjectID, err := idService.WithType(Hex128)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}

			mutex.Lock()
			defer mutex.Unlock()
			if _, ok := alreadySeen[newObjectID]; ok {
				t.Fatal("idService.New returned the same ID twice")
			}
			alreadySeen[newObjectID] = struct{}{}
		}()
	}
	wg.Wait()
}
