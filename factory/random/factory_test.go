package random

import (
	"crypto/rand"
	"io"
	"math/big"
	"testing"
	"time"

	"github.com/xh3b4sd/anna/spec"
)

func testMaybeNewRandomFactory(t *testing.T) spec.RandomFactory {
	newConfig := DefaultFactoryConfig()
	newFactory, err := NewFactory(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newFactory
}

func Test_RandomFactory_NewFactory_Error_RandFactory(t *testing.T) {
	newConfig := DefaultFactoryConfig()
	newConfig.RandFactory = nil

	_, err := NewFactory(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RandomFactory_NewFactory_Error_RandReader(t *testing.T) {
	newConfig := DefaultFactoryConfig()
	newConfig.RandReader = nil

	_, err := NewFactory(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RandomFactory_NewFactory_Error_Timeout(t *testing.T) {
	newConfig := DefaultFactoryConfig()
	newConfig.Timeout = 0

	_, err := NewFactory(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RandomFactory_CreateNBetween_Error_RandReader(t *testing.T) {
	newConfig := DefaultFactoryConfig()
	newConfig.Timeout = 10 * time.Millisecond

	newConfig.RandFactory = func(randReader io.Reader, max *big.Int) (n *big.Int, err error) {
		return nil, maskAny(invalidConfigError)
	}

	newFactory, err := NewFactory(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	n := 5
	max := 10

	_, err = newFactory.CreateNMax(n, max)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RandomFactory_CreateNBetween_Error_Timeout(t *testing.T) {
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

	n := 5
	max := 10

	_, err = newFactory.CreateNMax(n, max)
	if !IsTimeout(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RandomFactory_CreateNBetween_GenerateNNumbers(t *testing.T) {
	newConfig := DefaultFactoryConfig()
	newFactory, err := NewFactory(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	n := 5
	max := 10
	newRandomNumbers, err := newFactory.CreateNMax(n, max)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if len(newRandomNumbers) != 5 {
		t.Fatal("expected", 5, "got", len(newRandomNumbers))
	}
}

func Test_RandomFactory_CreateNBetween_GenerateRandomNumbers(t *testing.T) {
	newConfig := DefaultFactoryConfig()
	newFactory, err := NewFactory(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	n := 100
	max := 10
	newRandomNumbers, err := newFactory.CreateNMax(n, max)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	alreadySeen := map[int]struct{}{}

	for _, r := range newRandomNumbers {
		alreadySeen[r] = struct{}{}
	}

	l := len(alreadySeen)
	if l != 10 {
		t.Fatal("expected", 10, "got", l)
	}
}
