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
	HashMap  map[string]map[string]string
	Log      spec.Log
	MathSet  map[string]map[string]struct{}
	Weighted map[string]scoredElements
}

// DefaultConfig provides a default configuration to create a new memory
// storage object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		KeyValue: map[string]string{},
		HashMap:  map[string]map[string]string{},
		Log:      log.NewLog(log.DefaultConfig()),
		MathSet:  map[string]map[string]struct{}{},
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
	s.Mutex.Lock()
	weighted, ok := s.Weighted[key]
	s.Mutex.Unlock()
	if !ok {
		return nil, nil
	}

	if elements, ok := weighted.ScoreElements[score]; ok {
		n := maxElements
		if n > len(elements) || n < 0 {
			n = len(elements)
		}

		return elements[:n], nil
	}

	return nil, nil
}

func (s *storage) GetHashMap(key string) (map[string]string, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if value, ok := s.HashMap[key]; ok {
		return value, nil
	}

	return nil, nil
}

func (s *storage) GetHighestScoredElements(key string, maxElements int) ([]string, error) {
	s.Mutex.Lock()
	weighted, ok := s.Weighted[key]
	s.Mutex.Unlock()
	if !ok {
		return nil, nil
	}

	var t int
	var scoredElements []string
	orig := weighted.Scores
	l := len(orig)

	for i := 1; i <= l; i++ {
		score := orig[l-i]

		elements, err := s.GetElementsByScore(key, score, maxElements)
		if err != nil {
			return nil, maskAny(err)
		}

		formatted := strconv.FormatFloat(score, 'f', -1, 64)
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

func (s *storage) PushToSet(key string, element string) error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	set, ok := s.MathSet[key]
	if !ok {
		set = map[string]struct{}{
			element: {},
		}
	}

	set[element] = struct{}{}
	s.MathSet[key] = set

	return nil
}

func (s *storage) RemoveFromSet(key string, element string) error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	set, ok := s.MathSet[key]
	if !ok {
		return nil
	}
	delete(set, element)

	return nil
}

func (s *storage) RemoveScoredElement(key string, element string) error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	weighted, ok := s.Weighted[key]
	if !ok {
		return nil
	}

	score := weighted.ElementScores[element]
	delete(weighted.ElementScores, element)

	elements := weighted.ScoreElements[score]
	if len(elements) == 1 {
		delete(weighted.ScoreElements, score)
	} else {
		var newElements []string
		for _, e := range elements {
			if e != element {
				newElements = append(newElements, e)
			}
		}
		weighted.ScoreElements[score] = newElements
	}

	if len(elements) == 1 {
		// In case there was only one element, and we removed it, we also need to
		// remove the score from the "global" list.
		var newScores []float64
		for _, es := range weighted.Scores {
			if es != score {
				newScores = append(newScores, es)
			}
		}
		weighted.Scores = newScores
	}

	s.Weighted[key] = weighted

	return nil
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

	var foundScoreElements bool
	for _, item := range wl.ScoreElements[score] {
		if item == element {
			foundScoreElements = true
			break
		}
	}
	if !foundScoreElements {
		wl.ScoreElements[score] = append(wl.ScoreElements[score], element)
		sort.Strings(wl.ScoreElements[score])
	}

	var foundScores bool
	for _, item := range wl.Scores {
		if item == score {
			foundScores = true
			break
		}
	}
	if !foundScores {
		wl.Scores = append(wl.Scores, score)
		sort.Float64s(wl.Scores)
	}

	s.Weighted[key] = wl

	return nil
}

func (s *storage) SetHashMap(key string, hashMap map[string]string) error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.HashMap[key] = hashMap

	return nil
}

func (s *storage) WalkScoredElements(key string, closer <-chan struct{}, cb func(element string, score float64) error) error {
	s.Mutex.Lock()
	weighted, ok := s.Weighted[key]
	s.Mutex.Unlock()
	if !ok {
		return nil
	}

	for _, score := range weighted.Scores {
		for _, element := range weighted.ScoreElements[score] {
			select {
			case <-closer:
				return nil
			default:
			}

			err := cb(element, score)
			if err != nil {
				return maskAny(err)
			}
		}
	}

	return nil
}

func (s *storage) WalkSet(key string, closer <-chan struct{}, cb func(element string) error) error {
	s.Mutex.Lock()
	set, ok := s.MathSet[key]
	s.Mutex.Unlock()
	if !ok {
		return nil
	}

	for element := range set {
		select {
		case <-closer:
			return nil
		default:
		}

		err := cb(element)
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}
