package strategy

import (
	"math/rand"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

// Test_StrategyNetwork_NewStrategy_001 ensures that a specific combination of
// a strategy is created on a given random seed. This test also checks that the
// string representation works as expected.
func Test_StrategyNetwork_NewStrategy_001(t *testing.T) {
	testCases := []struct {
		Seed     int64
		Strategy string
	}{
		{
			Seed:     0,
			Strategy: "two,two",
		},
		{
			Seed:     1,
			Strategy: "two",
		},
		{
			Seed:     2,
			Strategy: "one",
		},
		{
			Seed:     99,
			Strategy: "one,one",
		},
	}

	var (
		one spec.ObjectType = "one"
		two spec.ObjectType = "two"
	)

	for i, testCase := range testCases {
		rand.Seed(testCase.Seed)

		newConfig := DefaultConfig()
		newConfig.Actions = []spec.ObjectType{
			one,
			two,
		}
		newStrategy := NewStrategy(newConfig)

		s := newStrategy.String()
		if s != testCase.Strategy {
			t.Fatal("test", i+1, "expected", testCase.Strategy, "got", s)
		}
	}
}
