package memory

import (
	"sort"
	"strconv"
	"sync"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/factory/random"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectType represents the object type of the memory storage object. This
	// is used e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "memory-storage"
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

// StorageConfig represents the configuration used to create a new memory storage
// object.
type StorageConfig struct {
	KeyValue      map[string]string
	Log           spec.Log
	MathSet       map[string]map[string]struct{}
	RandomFactory spec.RandomFactory
	StringMap     map[string]map[string]string
	Weighted      map[string]scoredElements
}

// DefaultStorageConfig provides a default configuration to create a new memory
// storage object by best effort.
func DefaultStorageConfig() StorageConfig {
	newRandomFactory, err := random.NewFactory(random.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}

	newConfig := StorageConfig{
		KeyValue:      map[string]string{},
		Log:           log.New(log.DefaultConfig()),
		MathSet:       map[string]map[string]struct{}{},
		RandomFactory: newRandomFactory,
		StringMap:     map[string]map[string]string{},
		Weighted:      map[string]scoredElements{},
	}

	return newConfig
}

// NewStorage creates a new configured memory storage object.
func NewStorage(config StorageConfig) (spec.Storage, error) {
	newStorage := &storage{
		StorageConfig: config,

		ID:    id.MustNew(),
		Mutex: sync.Mutex{},
		Type:  ObjectType,
	}

	newStorage.Log.Register(newStorage.GetType())

	return newStorage, nil
}

// MustNew creates either a new default configured storage object, or panics.
func MustNew() spec.Storage {
	newStorage, err := NewStorage(DefaultStorageConfig())
	if err != nil {
		panic(err)
	}

	return newStorage
}

type storage struct {
	StorageConfig

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

	return "", notFoundError
}

func (s *storage) GetElementsByScore(key string, score float64, maxElements int) ([]string, error) {
	s.Mutex.Lock()
	weighted, ok := s.Weighted[key]
	s.Mutex.Unlock()
	if !ok {
		return nil, notFoundError
	}

	if elements, ok := weighted.ScoreElements[score]; ok {
		n := maxElements
		if n > len(elements) || n < 0 {
			n = len(elements)
		}

		return elements[:n], nil
	}

	return nil, notFoundError
}

func (s *storage) GetHighestScoredElements(key string, maxElements int) ([]string, error) {
	s.Mutex.Lock()
	weighted, ok := s.Weighted[key]
	s.Mutex.Unlock()
	if !ok {
		return nil, notFoundError
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

func (s *storage) GetRandomKey() (string, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	// Here we create a random number to chose a random map, of which we have 4.
	// The random numbers starts at 0. So the maximum random number we want to
	// have is 3. Because the max parameter of CreateNMax is exclusive, we set max
	// to 4.
	mapIDs, err := s.RandomFactory.CreateNMax(1, 4)
	if err != nil {
		return "", maskAny(err)
	}

	switch mapIDs[0] {
	case 0:
		for k := range s.KeyValue {
			return k, nil
		}
	case 1:
		for k := range s.MathSet {
			return k, nil
		}
	case 2:
		for k := range s.StringMap {
			return k, nil
		}
	case 3:
		for k := range s.Weighted {
			return k, nil
		}
	}

	return "", notFoundError
}

func (s *storage) GetStringMap(key string) (map[string]string, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if value, ok := s.StringMap[key]; ok {
		return value, nil
	}

	return nil, notFoundError
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
		return notFoundError
	}
	delete(set, element)

	return nil
}

func (s *storage) RemoveScoredElement(key string, element string) error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	weighted, ok := s.Weighted[key]
	if !ok {
		return notFoundError
	}

	score, ok := weighted.ElementScores[element]
	if !ok {
		return notFoundError
	}
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

func (s *storage) SetStringMap(key string, stringMap map[string]string) error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.StringMap[key] = stringMap

	return nil
}

func (s *storage) WalkScoredElements(key string, closer <-chan struct{}, cb func(element string, score float64) error) error {
	s.Mutex.Lock()
	weighted, ok := s.Weighted[key]
	s.Mutex.Unlock()
	if !ok {
		return notFoundError
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
		return notFoundError
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
