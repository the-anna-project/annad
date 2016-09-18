package memory

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func Test_Memory_GetElementsByScore_Success(t *testing.T) {
	testCases := []struct {
		Key          string
		Score        float64
		MaxElements  int
		Elements     map[string]float64
		Expected     []string
		ErrorMatcher func(err error) bool
	}{
		{
			Key:         "mykey",
			Score:       0.5,
			MaxElements: -1,
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected: []string{
				"zero.five",
			},
			ErrorMatcher: nil,
		},
		{
			Key:         "mykey",
			Score:       0.5,
			MaxElements: 5,
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected: []string{
				"zero.five",
			},
			ErrorMatcher: nil,
		},
		{
			Key:         "mykey",
			Score:       0.5,
			MaxElements: 1,
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected: []string{
				"zero.five",
			},
			ErrorMatcher: nil,
		},
		{
			Key:         "mykey",
			Score:       3.4,
			MaxElements: 1,
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected:     []string{},
			ErrorMatcher: IsNotFound,
		},
		{
			Key:         "mykey",
			Score:       0.8,
			MaxElements: 5,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			// Note the order because of the lexicographical order.
			Expected: []string{
				"zero.eight.one",
				"zero.eight.three",
				"zero.eight.two",
			},
			ErrorMatcher: nil,
		},
		{
			Key:         "mykey",
			Score:       3.4,
			MaxElements: 5,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			Expected:     []string{},
			ErrorMatcher: IsNotFound,
		},
		{
			Key:         "mykey",
			Score:       0.8,
			MaxElements: 2,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			// Note the order because of the lexicographical order.
			Expected: []string{
				"zero.eight.one",
				"zero.eight.three",
			},
			ErrorMatcher: nil,
		},
		{
			Key:   "mykey",
			Score: 0.8,
			// Note we set MaxElements to zero, so nothing should be returned.
			MaxElements: 0,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			Expected:     []string{},
			ErrorMatcher: nil,
		},
		{
			Key:         "mykey",
			Score:       0.5,
			MaxElements: 1,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			Expected: []string{
				"zero.five",
			},
			ErrorMatcher: nil,
		},
	}

	for i, testCase := range testCases {
		// Setup
		newStorage := MustNew()

		for e, s := range testCase.Elements {
			err := newStorage.SetElementByScore(testCase.Key, e, s)
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
		}

		// Test
		output, err := newStorage.GetElementsByScore(testCase.Key, testCase.Score, testCase.MaxElements)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}

		// Assert
		if testCase.ErrorMatcher == nil {
			if len(output) != len(testCase.Expected) {
				t.Fatal("case", i+1, "expected", len(testCase.Expected), "got", len(output))
			}

			for i, e := range testCase.Expected {
				if output[i] != e {
					t.Fatal("case", i+1, "expected", e, "got", output[i])
				}
			}
		}
	}
}

func Test_Memory_GetElementsByScore_NotFound(t *testing.T) {
	newStorage := MustNew()

	_, err := newStorage.GetElementsByScore("not-found", 0.8, 3)
	if !IsNotFound(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_Memory_GetHighestScoredElements_Success(t *testing.T) {
	testCases := []struct {
		Key         string
		MaxElements int
		Elements    map[string]float64
		Expected    []string
	}{
		{
			Key:         "mykey",
			MaxElements: -1,
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected: []string{
				"zero.five",
				"0.5",
			},
		},
		{
			Key:         "mykey",
			MaxElements: 5,
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected: []string{
				"zero.five",
				"0.5",
			},
		},
		{
			Key:         "mykey",
			MaxElements: 1,
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected: []string{
				"zero.five",
				"0.5",
			},
		},
		{
			Key:         "mykey",
			MaxElements: 5,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			// Note the order because of the lexicographical order.
			Expected: []string{
				"zero.eight.one",
				"0.8",
				"zero.eight.three",
				"0.8",
				"zero.eight.two",
				"0.8",
				"zero.five",
				"0.5",
				"zero.one",
				"0.1",
			},
		},
		{
			Key:         "mykey",
			MaxElements: 0,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			Expected: []string{},
		},
		{
			Key:         "mykey",
			MaxElements: 2,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			// Note the order because of the lexicographical order.
			Expected: []string{
				"zero.eight.one",
				"0.8",
				"zero.eight.three",
				"0.8",
			},
		},
		{
			Key:         "mykey",
			MaxElements: 1,
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			Expected: []string{
				"zero.eight.one",
				"0.8",
			},
		},
	}

	for i, testCase := range testCases {
		// Setup
		newStorage := MustNew()

		for e, s := range testCase.Elements {
			err := newStorage.SetElementByScore(testCase.Key, e, s)
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
		}

		// Test
		output, err := newStorage.GetHighestScoredElements(testCase.Key, testCase.MaxElements)
		if err != nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}

		// Assert
		if len(output) != len(testCase.Expected) {
			t.Fatal("case", i+1, "expected", len(testCase.Expected), "got", len(output))
		}

		for i, e := range testCase.Expected {
			if output[i] != e {
				t.Fatal("case", i+1, "expected", e, "got", output[i])
			}
		}
	}
}

func Test_Memory_GetHighestScoredElements_NotFound(t *testing.T) {
	newStorage := MustNew()

	_, err := newStorage.GetHighestScoredElements("not-found", 3)
	if !IsNotFound(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_Memory_GetRandomKey(t *testing.T) {
	newStorage := MustNew()

	// We store 3 keys in each map of the memory storage to verify that we fetch a
	// random key across all stored keys.
	newStorage.Set("SetKey1", "SetValue1")
	newStorage.Set("SetKey2", "SetValue2")
	newStorage.Set("SetKey3", "SetValue3")
	newStorage.PushToSet("PushToSetKey1", "PushToSetElement1")
	newStorage.PushToSet("PushToSetKey2", "PushToSetElement2")
	newStorage.PushToSet("PushToSetKey3", "PushToSetElement3")
	newStorage.SetElementByScore("SetElementByScoreKey1", "SetElementByScoreElement1", 0)
	newStorage.SetElementByScore("SetElementByScoreKey2", "SetElementByScoreElement2", 0)
	newStorage.SetElementByScore("SetElementByScoreKey3", "SetElementByScoreElement3", 0)
	newStorage.SetStringMap("SetStringMapKey1", map[string]string{})
	newStorage.SetStringMap("SetStringMapKey2", map[string]string{})
	newStorage.SetStringMap("SetStringMapKey3", map[string]string{})

	alreadySeen := map[string]struct{}{}

	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newKey, err := newStorage.GetRandomKey()
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}

			mutex.Lock()
			alreadySeen[newKey] = struct{}{}
			mutex.Unlock()
		}()
	}
	wg.Wait()

	l := len(alreadySeen)
	if l != 12 {
		t.Fatal("expected", 12, "got", l)
	}
}

func Test_Memory_KeyValue(t *testing.T) {
	newStorage := MustNew()

	value, err := newStorage.Get("foo")
	if !IsNotFound(err) {
		t.Fatal("expected", nil, "got", err)
	}

	err = newStorage.Set("foo", "bar")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	value, err = newStorage.Get("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if value != "bar" {
		t.Fatal("expected", "bar", "got", value)
	}
}

func Test_Memory_Push_WalkScoredElements_Remove(t *testing.T) {
	newStorage := MustNew()

	err := newStorage.SetElementByScore("test-key", "test-value-1", 0.8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	var element2 []string
	var score2 []float64
	err = newStorage.WalkScoredElements("test-key", nil, func(element string, score float64) error {
		element2 = append(element2, element)
		score2 = append(score2, score)
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if !reflect.DeepEqual(element2, []string{"test-value-1"}) {
		t.Fatal("expected", []string{"test-value-1"}, "got", element2)
	}
	if !reflect.DeepEqual(score2, []float64{0.8}) {
		t.Fatal("expected", []float64{0.8}, "got", score2)
	}
	// Add second element.
	err = newStorage.SetElementByScore("test-key", "test-value-2", 0.8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	element2 = []string{}
	score2 = []float64{}
	err = newStorage.WalkScoredElements("test-key", nil, func(element string, score float64) error {
		element2 = append(element2, element)
		score2 = append(score2, score)
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if !reflect.DeepEqual(element2, []string{"test-value-1", "test-value-2"}) {
		t.Fatal("expected", []string{"test-value-1", "test-value-2"}, "got", element2)
	}
	if !reflect.DeepEqual(score2, []float64{0.8, 0.8}) {
		t.Fatal("expected", []float64{0.8, 0.8}, "got", score2)
	}

	// Check an element can be removed from a set.
	err = newStorage.RemoveScoredElement("test-key", "test-value-1")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = newStorage.RemoveScoredElement("test-key", "test-value-2")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	var element3 string
	var score3 float64
	err = newStorage.WalkScoredElements("test-key", nil, func(element string, score float64) error {
		element3 = element
		score3 = score
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element3 != "" {
		t.Fatal("expected", "", "got", element3)
	}
	if score3 != 0 {
		t.Fatal("expected", "", "got", score3)
	}
}

func Test_Memory_Push_WalkSet_Remove(t *testing.T) {
	newStorage := MustNew()

	// Check the set is empty by default
	var element1 string
	err := newStorage.WalkSet("test-key", nil, func(element string) error {
		element1 = element
		return nil
	})
	if !IsNotFound(err) {
		t.Fatal("expected", nil, "got", err)
	}
	if element1 != "" {
		t.Fatal("expected", "", "got", element1)
	}

	// Check an element can be pushed to a set.
	err = newStorage.PushToSet("test-key", "test-value")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	var element2 string
	err = newStorage.WalkSet("test-key", nil, func(element string) error {
		element2 = element
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element2 != "test-value" {
		t.Fatal("expected", "test-value", "got", element2)
	}

	// Check an element can be removed from a set.
	err = newStorage.RemoveFromSet("test-key", "test-value")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	var element3 string
	err = newStorage.WalkSet("test-key", nil, func(element string) error {
		element3 = element
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element3 != "" {
		t.Fatal("expected", "", "got", element3)
	}
}

func Test_Memory_RemoveFromSet_Empty(t *testing.T) {
	newStorage := MustNew()

	err := newStorage.RemoveFromSet("test-key", "test-value")
	if !IsNotFound(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_Memory_RemoveScoredElement(t *testing.T) {
	testCases := []struct {
		Elements      map[string]float64
		RemoveElement string
		Expected      []string
		ErrorMatcher  func(err error) bool
	}{
		{
			RemoveElement: "zero.five",
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected:     []string{},
			ErrorMatcher: nil,
		},
		{
			RemoveElement: "zero.five",
			Elements: map[string]float64{
				"zero.five":  0.5,
				"three.four": 3.4,
			},
			Expected: []string{
				"three.four",
				"3.4",
			},
			ErrorMatcher: nil,
		},
		{
			RemoveElement: "zero.five",
			Elements:      map[string]float64{},
			Expected:      []string{},
			ErrorMatcher:  IsNotFound,
		},
		{
			RemoveElement: "invalid",
			Elements: map[string]float64{
				"zero.five": 0.5,
			},
			Expected: []string{
				"zero.five",
				"0.5",
			},
			ErrorMatcher: IsNotFound,
		},
		{
			RemoveElement: "zero.eight.one",
			Elements: map[string]float64{
				"zero.five":        0.5,
				"zero.eight.one":   0.8,
				"zero.eight.two":   0.8,
				"zero.eight.three": 0.8,
				"zero.one":         0.1,
			},
			// Note the order because of the lexicographical order.
			Expected: []string{
				"zero.eight.three",
				"0.8",
				"zero.eight.two",
				"0.8",
				"zero.five",
				"0.5",
				"zero.one",
				"0.1",
			},
			ErrorMatcher: nil,
		},
	}

	for i, testCase := range testCases {
		// Setup
		newStorage := MustNew()

		for e, s := range testCase.Elements {
			err := newStorage.SetElementByScore("test-key", e, s)
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}
		}

		// Test
		err := newStorage.RemoveScoredElement("test-key", testCase.RemoveElement)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}

		// Assert
		if testCase.ErrorMatcher == nil {
			values, err := newStorage.GetHighestScoredElements("test-key", -1)
			if err != nil {
				t.Fatal("case", i+1, "expected", nil, "got", err)
			}

			if len(values) != len(testCase.Expected) {
				t.Fatal("case", i+1, "expected", len(testCase.Expected), "got", len(values))
			}

			for i, e := range testCase.Expected {
				if values[i] != e {
					t.Fatal("case", i+1, "expected", e, "got", values[i])
				}
			}
		}
	}
}

func Test_Memory_StringMap(t *testing.T) {
	newStorage := MustNew()

	value, err := newStorage.GetStringMap("foo")
	if !IsNotFound(err) {
		t.Fatal("expected", nil, "got", err)
	}

	err = newStorage.SetStringMap("foo", map[string]string{"bar": "baz"})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	value, err = newStorage.GetStringMap("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if !reflect.DeepEqual(value, map[string]string{"bar": "baz"}) {
		t.Fatal("expected", map[string]string{"bar": "baz"}, "got", value)
	}
}

func Test_Memory_WalkSet_Closer(t *testing.T) {
	newStorage := MustNew()

	err := newStorage.PushToSet("test-key", "test-value")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Immediately close the walk.
	closer := make(chan struct{}, 1)
	closer <- struct{}{}

	// Check that the walk does not happen, because we already ended it.
	var element1 string
	err = newStorage.WalkSet("test-key", closer, func(element string) error {
		element1 = element
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element1 != "" {
		t.Fatal("expected", "", "got", element1)
	}
}

func Test_Memory_WalkSet_Error(t *testing.T) {
	newStorage := MustNew()

	err := newStorage.PushToSet("test-key", "test-value")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Check that the walk does not happen, because we already ended it.
	err = newStorage.WalkSet("test-key", nil, func(element string) error {
		return fmt.Errorf("test-error")
	})
	if err == nil || err.Error() != "test-error" {
		t.Fatal("expected", "test-error", "got", err)
	}
}

func Test_Memory_WalkScoredElements_Closer(t *testing.T) {
	newStorage := MustNew()

	err := newStorage.SetElementByScore("test-key", "test-value", 0.8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Immediately close the walk.
	closer := make(chan struct{}, 1)
	closer <- struct{}{}

	// Check that the walk does not happen, because we already ended it.
	var element1 string
	var score1 float64
	err = newStorage.WalkScoredElements("test-key", closer, func(element string, score float64) error {
		element1 = element
		score1 = score
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element1 != "" {
		t.Fatal("expected", "", "got", element1)
	}
	if score1 != 0 {
		t.Fatal("expected", "", "got", score1)
	}
}

func Test_Memory_WalkScoredElements_Empty(t *testing.T) {
	newStorage := MustNew()

	// Check the set is empty by default
	var element1 string
	var score1 float64
	err := newStorage.WalkScoredElements("test-key", nil, func(element string, score float64) error {
		element1 = element
		score1 = score
		return nil
	})
	if !IsNotFound(err) {
		t.Fatal("expected", nil, "got", err)
	}
	if element1 != "" {
		t.Fatal("expected", "", "got", element1)
	}
	if score1 != 0 {
		t.Fatal("expected", "", "got", score1)
	}
}

func Test_Memory_WalkScoredElements_Error(t *testing.T) {
	newStorage := MustNew()

	err := newStorage.SetElementByScore("test-key", "test-value", 0.8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Check that the walk does not happen, because we already ended it.
	err = newStorage.WalkScoredElements("test-key", nil, func(element string, score float64) error {
		return fmt.Errorf("test-error")
	})
	if err == nil || err.Error() != "test-error" {
		t.Fatal("expected", "test-error", "got", err)
	}
}
