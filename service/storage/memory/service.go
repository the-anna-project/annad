package memory

import (
	"sync"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/cenk/backoff"

	objectspec "github.com/the-anna-project/spec/object"
	servicespec "github.com/the-anna-project/spec/service"
	"github.com/xh3b4sd/anna/service/storage/redis"
)

// New creates a new memory storage service. Therefore it manages an in-memory
// redis instance which can be shut down using the configured closer. This is
// used for local development.
func New() servicespec.StorageService {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	closer       chan struct{}
	metadata     map[string]string
	redisStorage servicespec.StorageService
	shutdownOnce sync.Once
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"kind": "memory",
		"name": "storage",
		"type": "service",
	}

	addrChan := make(chan string, 1)
	closer := make(chan struct{}, 1)
	redisAddr := ""

	go func() {
		s, err := miniredis.Run()
		if err != nil {
			panic(err)
		}
		addrChan <- s.Addr()

		<-closer
		s.Close()
	}()
	select {
	case <-time.After(1 * time.Second):
		panic("starting miniredis timed out")
	case addr := <-addrChan:
		redisAddr = addr
	}

	newRedisStorage := redis.New()
	newRedisStorage.SetAddress(redisAddr)
	newRedisStorage.SetBackoffFactory(func() objectspec.Backoff {
		return backoff.NewExponentialBackOff()
	})

	s.closer = closer
	s.redisStorage = newRedisStorage
	s.shutdownOnce = sync.Once{}
}

func (s *service) Get(key string) (string, error) {
	s.Service().Log().Line("func", "Get")

	result, err := s.redisStorage.Get(key)
	if err != nil {
		return "", maskAny(err)
	}

	return result, nil
}

func (s *service) GetAllFromSet(key string) ([]string, error) {
	s.Service().Log().Line("func", "GetAllFromSet")

	result, err := s.redisStorage.GetAllFromSet(key)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *service) GetElementsByScore(key string, score float64, maxElements int) ([]string, error) {
	s.Service().Log().Line("func", "GetElementsByScore")

	result, err := s.redisStorage.GetElementsByScore(key, score, maxElements)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *service) GetHighestScoredElements(key string, maxElements int) ([]string, error) {
	s.Service().Log().Line("func", "GetHighestScoredElements")

	result, err := s.redisStorage.GetHighestScoredElements(key, maxElements)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *service) GetRandom() (string, error) {
	s.Service().Log().Line("func", "GetRandom")

	result, err := s.redisStorage.GetRandom()
	if err != nil {
		return "", maskAny(err)
	}

	return result, nil
}

func (s *service) GetStringMap(key string) (map[string]string, error) {
	s.Service().Log().Line("func", "GetStringMap")

	result, err := s.redisStorage.GetStringMap(key)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) PopFromList(key string) (string, error) {
	s.Service().Log().Line("func", "PopFromList")

	result, err := s.redisStorage.PopFromList(key)
	if err != nil {
		return "", maskAny(err)
	}

	return result, nil
}

func (s *service) PushToList(key string, element string) error {
	s.Service().Log().Line("func", "PushToList")

	err := s.redisStorage.PushToList(key, element)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) PushToSet(key string, element string) error {
	s.Service().Log().Line("func", "PushToSet")

	err := s.redisStorage.PushToSet(key, element)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) Remove(key string) error {
	s.Service().Log().Line("func", "Remove")

	err := s.redisStorage.Remove(key)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) RemoveFromSet(key string, element string) error {
	s.Service().Log().Line("func", "RemoveFromSet")

	err := s.redisStorage.RemoveFromSet(key, element)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) RemoveScoredElement(key string, element string) error {
	s.Service().Log().Line("func", "RemoveScoredElement")

	err := s.redisStorage.RemoveScoredElement(key, element)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) Set(key, value string) error {
	s.Service().Log().Line("func", "Set")

	err := s.redisStorage.Set(key, value)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) SetAddress(a string) {
}

func (s *service) SetBackoffFactory(bf func() objectspec.Backoff) {
}

func (s *service) SetElementByScore(key, element string, score float64) error {
	s.Service().Log().Line("func", "SetElementByScore")

	err := s.redisStorage.SetElementByScore(key, element, score)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) SetPrefix(p string) {
}

func (s *service) SetServiceCollection(sc servicespec.ServiceCollection) {
	s.serviceCollection = sc
}

func (s *service) SetStringMap(key string, stringMap map[string]string) error {
	s.Service().Log().Line("func", "SetStringMap")

	err := s.redisStorage.SetStringMap(key, stringMap)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) Shutdown() {
	s.Service().Log().Line("func", "Shutdown")

	s.shutdownOnce.Do(func() {
		close(s.closer)
	})
}

func (s *service) WalkKeys(glob string, closer <-chan struct{}, cb func(key string) error) error {
	s.Service().Log().Line("func", "WalkKeys")

	err := s.redisStorage.WalkKeys(glob, closer, cb)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) WalkScoredSet(key string, closer <-chan struct{}, cb func(element string, score float64) error) error {
	s.Service().Log().Line("func", "WalkScoredSet")

	err := s.redisStorage.WalkScoredSet(key, closer, cb)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) WalkSet(key string, closer <-chan struct{}, cb func(element string) error) error {
	s.Service().Log().Line("func", "WalkSet")

	err := s.redisStorage.WalkSet(key, closer, cb)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
