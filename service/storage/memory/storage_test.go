package memory

import (
	"fmt"
	"os"
	"reflect"
	"sync"
	"testing"
	"time"

	kitlog "github.com/go-kit/kit/log"

	"github.com/the-anna-project/collection"
	"github.com/the-anna-project/id"
	memoryinstrumentor "github.com/the-anna-project/instrumentor/memory"
	"github.com/the-anna-project/log"
	"github.com/the-anna-project/random"
	servicespec "github.com/the-anna-project/spec/service"
)

func testNewStorage() servicespec.StorageService {
	idService := id.New()
	instrumentorService := memoryinstrumentor.New()
	logService := log.New()
	logService.SetRootLogger(kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr)))
	randomService := random.New()
	storageService := New()

	serviceCollection := collection.New()
	serviceCollection.SetIDService(idService)
	serviceCollection.SetInstrumentorService(instrumentorService)
	serviceCollection.SetLogService(logService)
	serviceCollection.SetRandomService(randomService)

	idService.SetServiceCollection(serviceCollection)
	instrumentorService.SetServiceCollection(serviceCollection)
	logService.SetServiceCollection(serviceCollection)
	randomService.SetServiceCollection(serviceCollection)
	storageService.SetServiceCollection(serviceCollection)

	storageService.Boot()

	return storageService
}

func Test_ListStorage_PushToListPopFromList(t *testing.T) {
	newStorage := testNewStorage()
	defer newStorage.Shutdown()

	var err error
	err = newStorage.PushToList("key", "element1")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = newStorage.PushToList("key", "element2")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = newStorage.PushToList("key", "element3")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	var element string
	element, err = newStorage.PopFromList("key")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element != "element1" {
		t.Fatal("expected", "element1", "got", element)
	}
	element, err = newStorage.PopFromList("key")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element != "element2" {
		t.Fatal("expected", "element2", "got", element)
	}
	element, err = newStorage.PopFromList("key")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element != "element3" {
		t.Fatal("expected", "element3", "got", element)
	}

	fail := make(chan error, 1)
	go func() {
		// Fetching elements from a list removes the fetched elements from the list.
		// After all elements are fetched from the list, the list must be empty.
		element, err = newStorage.PopFromList("key")
		if err != nil {
			fail <- maskAny(err)
			return
		}
		if element != "" {
			fail <- maskAny(fmt.Errorf("test failed"))
			return
		}
		fail <- maskAny(fmt.Errorf("test failed"))
	}()

	select {
	case <-time.After(100 * time.Millisecond):
		// The test succeeded.
	case err := <-fail:
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_ScoredSetStorage_GetElementsByScore(t *testing.T) {
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
			ErrorMatcher: nil,
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
			// Note the order because of the descending lexicographical order.
			Expected: []string{
				"zero.eight.two",
				"zero.eight.three",
				"zero.eight.one",
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
			ErrorMatcher: nil,
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
			// Note the order because of the descending lexicographical order.
			Expected: []string{
				"zero.eight.two",
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
		newStorage := testNewStorage()
		defer newStorage.Shutdown()

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

			for j, e := range testCase.Expected {
				if output[j] != e {
					t.Fatal("case", i+1, "expected", e, "got", output[j])
				}
			}
		}

		newStorage.Shutdown()
	}
}

func Test_ScoredSetStorage_GetHighestScoredElements(t *testing.T) {
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
			// Note the order because of the descending lexicographical order.
			Expected: []string{
				"zero.eight.two",
				"0.8",
				"zero.eight.three",
				"0.8",
				"zero.eight.one",
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
			Expected: []string{
				"zero.eight.two",
				"0.8",
			},
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
			// Note the order because of the descending lexicographical order.
			Expected: []string{
				"zero.eight.two",
				"0.8",
				"zero.eight.three",
				"0.8",
				"zero.eight.one",
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
				"zero.eight.two",
				"0.8",
				"zero.eight.three",
				"0.8",
			},
		},
	}

	for i, testCase := range testCases {
		// Setup
		newStorage := testNewStorage()
		defer newStorage.Shutdown()

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

		for j, e := range testCase.Expected {
			if output[j] != e {
				t.Fatal("case", i+1, "expected", e, "got", output[j])
			}
		}

		newStorage.Shutdown()
	}
}

func Test_ScoredSetStorage_RemoveScoredElement(t *testing.T) {
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
			ErrorMatcher:  nil,
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
			ErrorMatcher: nil,
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
			// Note the order because of the descending lexicographical order.
			Expected: []string{
				"zero.eight.two",
				"0.8",
				"zero.eight.three",
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
		newStorage := testNewStorage()
		defer newStorage.Shutdown()

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

			for j, e := range testCase.Expected {
				if values[j] != e {
					t.Fatal("case", i+1, "expected", e, "got", values[j])
				}
			}
		}

		newStorage.Shutdown()
	}
}

func Test_ScoredSetStorage_SetWalkRemove(t *testing.T) {
	newStorage := testNewStorage()
	defer newStorage.Shutdown()

	err := newStorage.SetElementByScore("test-key", "test-value-1", 0.8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	var element2 []string
	var score2 []float64
	err = newStorage.WalkScoredSet("test-key", nil, func(element string, score float64) error {
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
	err = newStorage.WalkScoredSet("test-key", nil, func(element string, score float64) error {
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
	err = newStorage.WalkScoredSet("test-key", nil, func(element string, score float64) error {
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

func Test_ScoredSetStorage_WalkScoredSet(t *testing.T) {
	newStorage := testNewStorage()
	defer newStorage.Shutdown()

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
	err = newStorage.WalkScoredSet("test-key", closer, func(element string, score float64) error {
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

func Test_SetStorage_PushGetAll(t *testing.T) {
	newStorage := testNewStorage()
	defer newStorage.Shutdown()

	var err error
	err = newStorage.PushToSet("key", "element1")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = newStorage.PushToSet("key", "element2")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	err = newStorage.PushToSet("key", "element3")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	elements, err := newStorage.GetAllFromSet("key")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if len(elements) != 3 {
		t.Fatal("expected", 3, "got", len(elements))
	}
	if elements[0] != "element1" {
		t.Fatal("expected", "element1", "got", elements[0])
	}
	if elements[1] != "element2" {
		t.Fatal("expected", "element2", "got", elements[1])
	}
	if elements[2] != "element3" {
		t.Fatal("expected", "element3", "got", elements[2])
	}

	// Fetching all elements from a set does not remove the fetched elements from
	// the set. Multiple calls to GetAllFromSet always must return the same
	// elements.
	elements, err = newStorage.GetAllFromSet("key")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if len(elements) != 3 {
		t.Fatal("expected", 3, "got", len(elements))
	}
}

func Test_SetStorage_WalkPushRemove(t *testing.T) {
	newStorage := testNewStorage()
	defer newStorage.Shutdown()

	// Check the set is empty by default
	var element1 string
	err := newStorage.WalkSet("test-key", nil, func(element string) error {
		element1 = element
		return nil
	})
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

func Test_SetStorage_WalkSet_Closer(t *testing.T) {
	newStorage := testNewStorage()
	defer newStorage.Shutdown()

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

func Test_Storage_Shutdown(t *testing.T) {
	newStorage := testNewStorage()

	newStorage.Shutdown()
	newStorage.Shutdown()
	newStorage.Shutdown()
}

func Test_StringMapStorage_GetSetGet(t *testing.T) {
	newStorage := testNewStorage()
	defer newStorage.Shutdown()

	value, err := newStorage.GetStringMap("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if len(value) != 0 {
		t.Fatal("expected", 0, "got", len(value))
	}

	err = newStorage.SetStringMap("foo", map[string]string{"bar": "baz"})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	value, err = newStorage.GetStringMap("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if len(value) != 1 {
		t.Fatal("expected", 1, "got", len(value))
	}
	v, ok := value["bar"]
	if !ok {
		t.Fatal("expected", true, "got", false)
	}
	if v != "baz" {
		t.Fatal("expected", "baz", "got", v)
	}
}

func Test_StringStorage_GetRandom(t *testing.T) {
	newStorage := testNewStorage()
	defer newStorage.Shutdown()

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
	newStorage.SetStringMap("SetStringMapKey1", map[string]string{"foo": "bar"})
	newStorage.SetStringMap("SetStringMapKey2", map[string]string{"foo": "bar"})
	newStorage.SetStringMap("SetStringMapKey3", map[string]string{"foo": "bar"})

	alreadySeen := map[string]struct{}{}

	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newKey, err := newStorage.GetRandom()
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

func Test_StringStorage_GetSetGet(t *testing.T) {
	newStorage := testNewStorage()
	defer newStorage.Shutdown()

	_, err := newStorage.Get("foo")
	if !IsNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}

	err = newStorage.Set("foo", "bar")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	value, err := newStorage.Get("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if value != "bar" {
		t.Fatal("expected", "bar", "got", value)
	}
}

func Test_StringStorage_SetGetRemoveGet(t *testing.T) {
	newStorage := testNewStorage()
	defer newStorage.Shutdown()

	err := newStorage.Set("foo", "bar")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	value, err := newStorage.Get("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if value != "bar" {
		t.Fatal("expected", "bar", "got", value)
	}

	err = newStorage.Remove("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	_, err = newStorage.Get("foo")
	if !IsNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringStorage_WalkSetRemove(t *testing.T) {
	newStorage := testNewStorage()
	defer newStorage.Shutdown()

	// Verify the key space is empty by default.
	var count1 int
	err := newStorage.WalkKeys("*", nil, func(element string) error {
		count1++
		return nil
	})
	if count1 != 0 {
		t.Fatal("expected", 0, "got", count1)
	}

	// Set a new key.
	err = newStorage.Set("test-key", "test-value")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	var count2 int
	var element2 string
	err = newStorage.WalkKeys("*", nil, func(element string) error {
		count2++
		element2 = element
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if count2 != 1 {
		t.Fatal("expected", 1, "got", count2)
	}
	if element2 != "prefix:test-key" {
		t.Fatal("expected", "prefix:test-key", "got", element2)
	}

	// Remove one key.
	err = newStorage.Remove("test-key")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Verify there is now no key anymore.
	var count3 int
	err = newStorage.WalkKeys("*", nil, func(element string) error {
		count3++
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if count3 != 0 {
		t.Fatal("expected", 0, "got", count3)
	}
}

func Test_StringStorage_WalkKeys_Closer(t *testing.T) {
	newStorage := testNewStorage()
	defer newStorage.Shutdown()

	err := newStorage.PushToSet("test-key", "test-value")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	// Immediately close the walk.
	closer := make(chan struct{}, 1)
	closer <- struct{}{}

	// Check that the walk does not happen, because we already ended it.
	var element1 string
	err = newStorage.WalkKeys("test-key", closer, func(element string) error {
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
