package distribution

import (
	"testing"
)

func Test_Distribution_GetType(t *testing.T) {
	newConfig := DefaultDistributionConfig()
	newConfig.Name = "test"
	newConfig.StaticChannels = []float64{50, 100}
	newConfig.Vectors = [][]float64{{0}, {0}, {0}}
	newDistribution, err := NewDistribution(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if newDistribution.GetType() != ObjectTypeDistribution {
		t.Fatal("expected", ObjectTypeDistribution, "got", newDistribution.GetType())
	}
}

func Test_Distribution_GetID(t *testing.T) {
	newFirstConfig := DefaultDistributionConfig()
	newFirstConfig.Name = "test"
	newFirstConfig.StaticChannels = []float64{50, 100}
	newFirstConfig.Vectors = [][]float64{{0}, {0}, {0}}
	newFirstDistribution, err := NewDistribution(newFirstConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newSecondConfig := DefaultDistributionConfig()
	newSecondConfig.Name = "test"
	newSecondConfig.StaticChannels = []float64{50, 100}
	newSecondConfig.Vectors = [][]float64{{0}, {0}, {0}}
	newSecondDistribution, err := NewDistribution(newSecondConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if newFirstDistribution.GetID() == newSecondDistribution.GetID() {
		t.Fatalf("IDs of schedulers are equal")
	}
}
