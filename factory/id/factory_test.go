package id

import (
	"crypto/rand"
	"io"
	"math/big"
	"sync"
	"testing"
	"time"

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
	newConfig.RandFactory = nil

	_, err := NewFactory(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_IDFactory_NewFactory_Error_RandReader(t *testing.T) {
	newConfig := DefaultFactoryConfig()
	newConfig.RandReader = nil

	_, err := NewFactory(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_IDFactory_NewFactory_Error_Timeout(t *testing.T) {
	newConfig := DefaultFactoryConfig()
	newConfig.Timeout = 0

	_, err := NewFactory(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_IDFactory_WithType_Error_RandReader(t *testing.T) {
	newConfig := DefaultFactoryConfig()
	newConfig.Timeout = 10 * time.Millisecond

	newConfig.RandFactory = func(randReader io.Reader, max *big.Int) (n *big.Int, err error) {
		return nil, maskAny(invalidConfigError)
	}

	newFactory, err := NewFactory(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	_, err = newFactory.WithType(Hex128)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_IDFactory_WithType_Error_Timeout(t *testing.T) {
	newConfig := DefaultFactoryConfig()
	newConfig.Timeout = 20 * time.Millisecond

	newConfig.RandFactory = func(randReader io.Reader, max *big.Int) (n *big.Int, err error) {
		time.Sleep(200 * time.Millisecond)
		return rand.Int(randReader, max)
	}

	newFactory, err := NewFactory(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	_, err = newFactory.WithType(Hex128)
	if !IsTimeout(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

// Test_IDFactory_WithType checks that a generated ID is still unique after a
// certain number of concurrent generations.
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
