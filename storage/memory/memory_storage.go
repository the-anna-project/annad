// Package memorystorage implements spec.Storage and provides functionality to
// persist data in memory.
package memorystorage

import (
	"sort"
	"strconv"
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeMemoryStorage represents the object type of the memory storage
	// object. This is used e.g. to register itself to the logger.
	ObjectTypeMemoryStorage spec.ObjectType = "memory-storage"
)

type scoredElements struct {
	// ElementScores represents the mapping between elements and their scores.
	// This is a 1:1 relation. Key is element. Value is score.
	ElementScores map[string]float64

	// ScoreElements represents the mapping between scores and associated
	// elements. This is a 1:n relation. Key is Score. Value is a list of
	// elements.
	ScoreElements map[float64][]string

	// Scores holds ordered scores. Lowest score first. Highest score last.
	Scores []float64
}

// Config represents the configuration used to create a new memory storage
// object.
type Config struct {
	KeyValue map[string]string
	Log      spec.Log
	Weighted map[string]scoredElements
}

// DefaultConfig provides a default configuration to create a new memory
// storage object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		KeyValue: map[string]string{},
		Log:      log.NewLog(log.DefaultConfig()),
		Weighted: map[string]scoredElements{},
	}

	return newConfig
}

// NewMemoryStorage creates a new configured memory storage object.
func NewMemoryStorage(config Config) spec.Storage {
	newStorage := &storage{
		Config: config,

		ID:    id.NewObjectID(id.Hex128),
		Mutex: sync.Mutex{},
		Type:  ObjectTypeMemoryStorage,
	}

	newStorage.Log.Register(newStorage.GetType())

	return newStorage
}

type storage struct {
	Config

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (s *storage) Get(key string) (string, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if value, ok := s.KeyValue[key]; ok {
		return value, nil
	}

	return "", nil
}

func (s *storage) GetElementsByScore(key string, score float64, maxElements int) ([]string, error) {
	if _, ok := s.Weighted[key]; !ok {
		return nil, nil
	}

	if elements, ok := s.Weighted[key].ScoreElements[score]; ok {
		n := maxElements
		if n > len(elements) {
			n = len(elements)
		}

		return elements[:n], nil
	}

	return nil, nil
}

func (s *storage) GetHighestScoredElements(key string, maxElements int) ([]string, error) {
	if _, ok := s.Weighted[key]; !ok {
		return nil, nil
	}

	var t int
	var scoredElements []string
	orig := s.Weighted[key].Scores
	l := len(orig)

	for i := 1; i <= l; i++ {
		score := orig[l-i]
		formatted := strconv.FormatFloat(score, 'f', 5, 64)

		elements, err := s.GetElementsByScore(key, score, maxElements)
		if err != nil {
			return nil, maskAny(err)
		}

		for _, e := range elements {
			scoredElements = append(scoredElements, e)
			scoredElements = append(scoredElements, formatted)

			t++
			if t == maxElements {
				return scoredElements, nil
			}
		}
	}

	return scoredElements, nil
}

func (s *storage) Set(key, value string) error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.KeyValue[key] = value

	return nil
}

func (s *storage) SetElementByScore(key, element string, score float64) error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	formatted := strconv.FormatFloat(score, 'f', 5, 64)
	score, err := strconv.ParseFloat(formatted, 64)
	if err != nil {
		return maskAny(err)
	}

	// Initialize the weighted list.
	if _, ok := s.Weighted[key]; !ok {
		// The weighted list for key does not yet exist.
		s.Weighted[key] = scoredElements{
			ElementScores: map[string]float64{},
			ScoreElements: map[float64][]string{},
			Scores:        []float64{},
		}
	}

	// Add and sort
	wl := s.Weighted[key]

	wl.ElementScores[element] = score

	wl.ScoreElements[score] = append(wl.ScoreElements[score], element)
	sort.Strings(wl.ScoreElements[score])

	var found bool
	for _, item := range wl.Scores {
		if item == score {
			found = true
			break
		}
	}
	if !found {
		wl.Scores = append(wl.Scores, score)
		sort.Float64s(wl.Scores)
	}

	s.Weighted[key] = wl

	return nil
}
