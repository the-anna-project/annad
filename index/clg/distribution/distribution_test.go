package distribution

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func Test_NewDistribution_Success(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.Name = "test"
	newConfig.StaticChannels = []float64{50, 100}
	newConfig.Vectors = [][]float64{{0}, {0}, {0}}
	_, err := NewDistribution(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_NewDistribution_Error_Name(t *testing.T) {
	newConfig := DefaultConfig()
	// Name configuration is missing.
	newConfig.StaticChannels = []float64{50, 100}
	newConfig.Vectors = [][]float64{{0}, {0}, {0}}
	_, err := NewDistribution(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_NewDistribution_Error_Vectors_Empty(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.Name = "test"
	newConfig.StaticChannels = []float64{50, 100}
	// Vectors configuration is missing.
	_, err := NewDistribution(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_NewDistribution_Error_Vectors_Dimension(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.Name = "test"
	newConfig.StaticChannels = []float64{50, 100}
	// Vectors configuration is invalid.
	newConfig.Vectors = [][]float64{{0}, {0, 1}, {0}}
	_, err := NewDistribution(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_NewDistribution_Error_StaticChannels_Empty(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.Name = "test"
	// StaticChannels configuration is missing.
	newConfig.StaticChannels = []float64{}
	newConfig.Vectors = [][]float64{{0}, {0}, {0}}
	_, err := NewDistribution(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_NewDistribution_Error_StaticChannels_Duplicate(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.Name = "test"
	// StaticChannels configuration is invalid.
	newConfig.StaticChannels = []float64{25, 25}
	newConfig.Vectors = [][]float64{{0}, {0}, {0}}
	_, err := NewDistribution(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Distribution_Calculate(t *testing.T) {
	testCases := []struct {
		StaticChannels []float64
		Vectors        [][]float64
		Expected       []float64
	}{
		{
			StaticChannels: []float64{20, 40, 60, 80, 100},
			Vectors: [][]float64{
				{0, 1},
				{2, 3},
				{4, 5},
				{6, 7},
				{8, 9},
			},
			Expected: []float64{5, 0, 0, 0, 0},
		},
		{
			StaticChannels: []float64{50, 100},
			Vectors: [][]float64{
				{0, 1},
				{2, 3},
				{4, 5},
				{6, 7},
				{8, 9},
			},
			Expected: []float64{5, 0},
		},
		{
			StaticChannels: []float64{2, 4, 6, 8, 10},
			Vectors: [][]float64{
				{0, 1},
				{2, 3},
				{4, 5},
				{6, 7},
				{8, 9},
			},
			Expected: []float64{1.5, 1, 1, 1, 0.5},
		},
		{
			StaticChannels: []float64{10, 20, 30, 40, 50},
			Vectors: [][]float64{
				{3, 4},
				{6, 7},
				{15, 16},
				// Note this two vectors are located outside the respected channel range.
				{61, 62},
				{92, 93},
			},
			Expected: []float64{2, 1, 0, 0, 0},
		},
		{
			StaticChannels: []float64{20, 40, 60, 80, 100},
			Vectors: [][]float64{
				// Note this is the only vector covering all given channels. Thus its
				// weight is divided across them.
				{0, 100},
			},
			Expected: []float64{0.2, 0.2, 0.2, 0.2, 0.2},
		},
	}

	for i, testCase := range testCases {
		newConfig := DefaultConfig()
		newConfig.Name = "test"
		newConfig.StaticChannels = testCase.StaticChannels
		newConfig.Vectors = testCase.Vectors
		newDistribution, err := NewDistribution(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		output := newDistribution.Calculate()

		if len(output) != len(testCase.Expected) {
			t.Fatal("case", i+1, "expected", len(testCase.Expected), "got", len(output))
		}
		for j, e := range testCase.Expected {
			if output[j] != e {
				t.Fatal("case", j+1, "of", i+1, "expected", e, "got", output[j])
			}
		}
	}
}

func Test_Distribution_Difference(t *testing.T) {
	testCases := []struct {
		StaticChannels1 []float64
		StaticChannels2 []float64
		Vectors1        [][]float64
		Vectors2        [][]float64
		Expected        []float64
		ErrorMatcher    func(err error) bool
	}{
		{
			StaticChannels1: []float64{50, 100},
			StaticChannels2: []float64{50, 100},
			Vectors1:        [][]float64{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}},
			Vectors2:        [][]float64{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}},
			Expected:        []float64{0, 0},
		},
		{
			StaticChannels1: []float64{20, 40, 60, 80, 100},
			StaticChannels2: []float64{20, 40, 60, 80, 100},
			Vectors1:        [][]float64{{0, 1}, {2, 3}},
			Vectors2:        [][]float64{{0, 1}, {2, 3}},
			Expected:        []float64{0, 0, 0, 0, 0},
		},
		{
			StaticChannels1: []float64{20, 40, 60, 80, 100},
			StaticChannels2: []float64{20, 40, 60, 80, 100},
			Vectors1:        [][]float64{{0, 1}, {8, 9}},
			Vectors2:        [][]float64{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}},
			Expected:        []float64{3, 0, 0, 0, 0},
		},
		{
			StaticChannels1: []float64{20, 40, 60, 80, 100},
			StaticChannels2: []float64{20, 40, 60, 80, 100},
			Vectors1:        [][]float64{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}},
			Vectors2:        [][]float64{{0, 1}, {8, 9}},
			Expected:        []float64{-3, 0, 0, 0, 0},
		},
		{
			StaticChannels1: []float64{20, 100},
			StaticChannels2: []float64{20, 40, 60, 80, 100},
			Vectors1:        [][]float64{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}},
			Vectors2:        [][]float64{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}},
			ErrorMatcher:    IsChannelsDiffer,
		},
		{
			StaticChannels1: []float64{20, 40, 60, 80, 100},
			StaticChannels2: []float64{20, 100},
			Vectors1:        [][]float64{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}},
			Vectors2:        [][]float64{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}},
			ErrorMatcher:    IsChannelsDiffer,
		},
	}

	for i, testCase := range testCases {
		newConfig1 := DefaultConfig()
		newConfig1.Name = "test"
		newConfig1.StaticChannels = testCase.StaticChannels1
		newConfig1.Vectors = testCase.Vectors1
		newDistribution1, err := NewDistribution(newConfig1)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		newConfig2 := DefaultConfig()
		newConfig2.Name = "test"
		newConfig2.StaticChannels = testCase.StaticChannels2
		newConfig2.Vectors = testCase.Vectors2
		newDistribution2, err := NewDistribution(newConfig2)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		output, err := newDistribution1.Difference(newDistribution2)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			if len(output) != len(testCase.Expected) {
				t.Fatal("case", i+1, "expected", len(testCase.Expected), "got", len(output))
			}
			for j, e := range testCase.Expected {
				if output[j] != e {
					t.Fatal("case", j+1, "of", i+1, "expected", e, "got", output[j])
				}
			}
		}
	}
}

func Test_Distribution_GetDimensions(t *testing.T) {
	testCases := []struct {
		StaticChannels []float64
		Vectors        [][]float64
		Expected       int
	}{
		{
			StaticChannels: []float64{50, 100},
			Vectors: [][]float64{
				{0, 1},
				{2, 3},
			},
			Expected: 2,
		},
		{
			StaticChannels: []float64{50, 100},
			Vectors: [][]float64{
				{0, 1, 2},
				{3, 4, 5},
			},
			Expected: 3,
		},
		{
			StaticChannels: []float64{50, 100},
			Vectors: [][]float64{
				{5, 6, 7, 8, 9},
				{3, 4, 5, 6, 7},
			},
			Expected: 5,
		},
	}

	for i, testCase := range testCases {
		newConfig := DefaultConfig()
		newConfig.Name = "test"
		newConfig.StaticChannels = testCase.StaticChannels
		newConfig.Vectors = testCase.Vectors
		newDistribution, err := NewDistribution(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		output := newDistribution.GetDimensions()

		if output != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
		}
	}
}

func Test_Distribution_GetStringMap_Success(t *testing.T) {
	newStringMap := map[string]string{
		"name":            "name",
		"id":              "id",
		"static-channels": "25,50,100",
		"vectors":         "2,3|14,15|38,49",
	}

	newConfig := DefaultConfig()
	newConfig.StringMap = newStringMap
	newDistribution, err := NewDistribution(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if newDistribution.GetName() != "name" {
		t.Fatal("expected", "name", "got", newDistribution.GetName())
	}
	if newDistribution.GetID() != spec.ObjectID("id") {
		t.Fatal("expected", spec.ObjectID("id"), "got", newDistribution.GetID())
	}
	if !reflect.DeepEqual(newDistribution.GetStaticChannels(), []float64{25, 50, 100}) {
		t.Fatal("expected", []float64{25, 50, 100}, "got", newDistribution.GetStaticChannels())
	}
	if !reflect.DeepEqual(newDistribution.GetVectors(), [][]float64{{2, 3}, {14, 15}, {38, 49}}) {
		t.Fatal("expected", [][]float64{{2, 3}, {14, 15}, {38, 49}}, "got", newDistribution.GetVectors())
	}

	output := newDistribution.GetStringMap()

	if len(output) != len(newStringMap) {
		t.Fatal("expected", len(newStringMap), "got", len(output))
	}
	if output["name"] != "name" {
		t.Fatal("expected", "name", "got", output["name"])
	}
	if output["id"] != "id" {
		t.Fatal("expected", "id", "got", output["id"])
	}
	if output["static-channels"] != "25,50,100" {
		t.Fatal("expected", "25,50,100", "got", output["static-channels"])
	}
	if output["vectors"] != "2,3|14,15|38,49" {
		t.Fatal("expected", "2,3|14,15|38,49", "got", output["vectors"])
	}
}

func Test_Distribution_GetStringMap_Error(t *testing.T) {
	newStringMap := map[string]string{
		"name":            "name",
		"id":              "id",
		"static-channels": "25,invalid,100",
		"vectors":         "2,3|14,15|38,49",
	}
	newConfig := DefaultConfig()
	newConfig.StringMap = newStringMap
	_, err := NewDistribution(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}

	newStringMap = map[string]string{
		"name":            "name",
		"id":              "id",
		"static-channels": "25,50,100",
		"vectors":         "2,3|invalid|38,49",
	}
	newConfig = DefaultConfig()
	newConfig.StringMap = newStringMap
	_, err = NewDistribution(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Distribution_GetName(t *testing.T) {
	testCases := []struct {
		Name           string
		StaticChannels []float64
		Vectors        [][]float64
		Expected       string
	}{
		{
			Name:           "foo",
			StaticChannels: []float64{50, 100},
			Vectors:        [][]float64{{0, 1}, {2, 3}},
			Expected:       "foo",
		},
		{
			Name:           "test",
			StaticChannels: []float64{50, 100},
			Vectors:        [][]float64{{0, 1}, {2, 3}},
			Expected:       "test",
		},
		{
			Name:           "name",
			StaticChannels: []float64{50, 100},
			Vectors:        [][]float64{{0, 1}, {2, 3}},
			Expected:       "name",
		},
	}

	for i, testCase := range testCases {
		newConfig := DefaultConfig()
		newConfig.Name = testCase.Name
		newConfig.StaticChannels = testCase.StaticChannels
		newConfig.Vectors = testCase.Vectors
		newDistribution, err := NewDistribution(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		output := newDistribution.GetName()

		if output != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
		}
	}
}

func Test_Distribution_GetStaticChannels(t *testing.T) {
	testCases := []struct {
		StaticChannels []float64
		Vectors        [][]float64
		Expected       []float64
	}{
		{
			StaticChannels: []float64{50, 100},
			Vectors:        [][]float64{{0, 1}, {2, 3}},
			Expected:       []float64{50, 100},
		},
		{
			StaticChannels: []float64{20, 50, 80},
			Vectors:        [][]float64{{0, 1}, {2, 3}},
			Expected:       []float64{20, 50, 80},
		},
		{
			StaticChannels: []float64{10, 33, 90},
			Vectors:        [][]float64{{0, 1}, {2, 3}},
			Expected:       []float64{10, 33, 90},
		},
	}

	for i, testCase := range testCases {
		newConfig := DefaultConfig()
		newConfig.Name = "test"
		newConfig.StaticChannels = testCase.StaticChannels
		newConfig.Vectors = testCase.Vectors
		newDistribution, err := NewDistribution(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		output := newDistribution.GetStaticChannels()

		if len(output) != len(testCase.Expected) {
			t.Fatal("case", i+1, "expected", len(testCase.Expected), "got", len(output))
		}
		for j, e := range testCase.Expected {
			if output[j] != e {
				t.Fatal("case", j+1, "of", i+1, "expected", e, "got", output[j])
			}
		}
	}
}

func Test_Distribution_GetVectors(t *testing.T) {
	testCases := []struct {
		StaticChannels []float64
		Vectors        [][]float64
		Expected       [][]float64
	}{
		{
			StaticChannels: []float64{50, 100},
			Vectors:        [][]float64{{0, 1}, {2, 3}},
			Expected:       [][]float64{{0, 1}, {2, 3}},
		},
		{
			StaticChannels: []float64{20, 50, 80},
			Vectors:        [][]float64{{0, 1}, {2, 3}},
			Expected:       [][]float64{{0, 1}, {2, 3}},
		},
		{
			StaticChannels: []float64{10, 33, 90},
			Vectors:        [][]float64{{0, 1}, {2, 3}},
			Expected:       [][]float64{{0, 1}, {2, 3}},
		},
	}

	for i, testCase := range testCases {
		newConfig := DefaultConfig()
		newConfig.Name = "test"
		newConfig.StaticChannels = testCase.StaticChannels
		newConfig.Vectors = testCase.Vectors
		newDistribution, err := NewDistribution(newConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		output := newDistribution.GetVectors()

		if len(output) != len(testCase.Expected) {
			t.Fatal("case", i+1, "expected", len(testCase.Expected), "got", len(output))
		}
		for j, e := range testCase.Expected {
			if !reflect.DeepEqual(output[j], e) {
				t.Fatal("case", j+1, "of", i+1, "expected", e, "got", output[j])
			}
		}
	}
}
