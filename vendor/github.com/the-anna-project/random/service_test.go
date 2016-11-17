package random

import (
	"crypto/rand"
	"io"
	"math/big"
	"testing"
	"time"
)

func Test_RandomService_CreateNBetween_Error_RandFactory(t *testing.T) {
	newService := New()
	newService.SetTimeout(10 * time.Millisecond)
	newService.SetRandFactory(func(randReader io.Reader, max *big.Int) (n *big.Int, err error) {
		return nil, maskAny(timeoutError)
	})

	n := 5
	max := 10

	_, err := newService.CreateNMax(n, max)
	if !IsTimeout(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RandomService_CreateNBetween_Error_Timeout(t *testing.T) {
	newService := New()
	newService.SetTimeout(20 * time.Millisecond)
	newService.SetRandFactory(func(randReader io.Reader, max *big.Int) (n *big.Int, err error) {
		time.Sleep(200 * time.Millisecond)
		return rand.Int(randReader, max)
	})

	n := 5
	max := 10

	_, err := newService.CreateNMax(n, max)
	if !IsTimeout(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_RandomService_CreateNBetween_GenerateNNumbers(t *testing.T) {
	newService := New()

	n := 5
	max := 10

	newRandomNumbers, err := newService.CreateNMax(n, max)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if len(newRandomNumbers) != 5 {
		t.Fatal("expected", 5, "got", len(newRandomNumbers))
	}
}

func Test_RandomService_CreateNBetween_GenerateRandomNumbers(t *testing.T) {
	newService := New()

	n := 100
	max := 10

	newRandomNumbers, err := newService.CreateNMax(n, max)
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
