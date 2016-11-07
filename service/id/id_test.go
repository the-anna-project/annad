package id

import (
	"io"
	"math/big"
	"sync"
	"testing"

	"github.com/xh3b4sd/anna/service/random"
	"github.com/xh3b4sd/anna/service/spec"
)

func Test_IDService_NewService_Error_HashChars(t *testing.T) {
	newConfig := DefaultServiceConfig()
	newConfig.HashChars = ""

	_, err := NewService(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_IDService_NewService_Error_RandService(t *testing.T) {
	newConfig := DefaultServiceConfig()
	newConfig.RandomService = nil

	_, err := NewService(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_IDService_NewService_Error_Type(t *testing.T) {
	newConfig := DefaultServiceConfig()
	newConfig.Type = spec.IDType(0)

	_, err := NewService(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_IDService_New_Error(t *testing.T) {
	// Create custom random service with timeout config.
	newRandomServiceConfig := random.DefaultServiceConfig()
	newRandomServiceConfig.RandFactory = func(randReader io.Reader, max *big.Int) (n *big.Int, err error) {
		return nil, maskAny(invalidConfigError)
	}
	newRandomService, err := random.NewService(newRandomServiceConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newConfig := DefaultServiceConfig()
	newConfig.RandomService = newRandomService
	newService, err := NewService(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	_, err = newService.New()
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_IDService_WithType_Error(t *testing.T) {
	// Create custom random service with timeout config.
	newRandomServiceConfig := random.DefaultServiceConfig()
	newRandomServiceConfig.RandFactory = func(randReader io.Reader, max *big.Int) (n *big.Int, err error) {
		return nil, maskAny(invalidConfigError)
	}
	newRandomService, err := random.NewService(newRandomServiceConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newConfig := DefaultServiceConfig()
	newConfig.RandomService = newRandomService
	newService, err := NewService(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	_, err = newService.WithType(Hex128)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

// Test_IDService_New checks that a generated ID is still unique after a
// certain number of concurrent generations.
func Test_IDService_New(t *testing.T) {
	newService := MustNewService()

	alreadySeen := map[string]struct{}{}

	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newObjectID, err := newService.New()
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}

			mutex.Lock()
			defer mutex.Unlock()
			if _, ok := alreadySeen[newObjectID]; ok {
				t.Fatal("id.NewObjectID returned the same ID twice")
			}
			alreadySeen[newObjectID] = struct{}{}
		}()
	}
	wg.Wait()
}

// Test_IDService_WithType checks that a generated ID is still unique after a
// certain number of concurrent generations.
func Test_IDService_WithType(t *testing.T) {
	newService := MustNewService()

	alreadySeen := map[string]struct{}{}

	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newObjectID, err := newService.WithType(Hex128)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}

			mutex.Lock()
			defer mutex.Unlock()
			if _, ok := alreadySeen[newObjectID]; ok {
				t.Fatal("id.NewObjectID returned the same ID twice")
			}
			alreadySeen[newObjectID] = struct{}{}
		}()
	}
	wg.Wait()
}

func Test_IDService_MustNewID(t *testing.T) {
	if MustNew() == MustNew() {
		t.Fatal("expected", false, "got", true)
	}
}
