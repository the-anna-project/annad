package id

import (
	"io"
	"math/big"
	"sync"
	"testing"

	"github.com/xh3b4sd/anna/factory/random"
	"github.com/xh3b4sd/anna/spec"
)

func Test_IDFactory_NewFactory_Error_HashChars(t *testing.T) {
	newConfig := DefaultFactoryConfig()
	newConfig.HashChars = ""

	_, err := NewFactory(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_IDFactory_NewFactory_Error_RandFactory(t *testing.T) {
	newConfig := DefaultFactoryConfig()
	newConfig.RandomFactory = nil

	_, err := NewFactory(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_IDFactory_WithType_Error(t *testing.T) {
	// Create custom random factory with timeout config.
	newRandomFactoryConfig := random.DefaultFactoryConfig()
	newRandomFactoryConfig.RandFactory = func(randReader io.Reader, max *big.Int) (n *big.Int, err error) {
		return nil, maskAny(invalidConfigError)
	}
	newRandomFactory, err := random.NewFactory(newRandomFactoryConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newConfig := DefaultFactoryConfig()
	newConfig.RandomFactory = newRandomFactory
	newFactory, err := NewFactory(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	_, err = newFactory.WithType(Hex128)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

// Test_IDFactory_New checks that a generated ID is still unique after a
// certain number of concurrent generations.
func Test_IDFactory_New(t *testing.T) {
	newFactory := MustNewFactory()

	alreadySeen := map[spec.ObjectID]struct{}{}

	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newObjectID, err := newFactory.New()
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

// Test_IDFactory_WithType checks that a generated ID is still unique after a
// certain number of concurrent generations.
func Test_IDFactory_WithType(t *testing.T) {
	newFactory := MustNewFactory()

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
