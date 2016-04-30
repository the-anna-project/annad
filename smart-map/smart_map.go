// Package smartmap provides functionality to assign typed input to a storable
// string map.
package smartmap

import (
	"sync"

	"github.com/xh3b4sd/anna/spec"
)

const (
	// StringPrefix represents the prefix used to label keys and values having
	// the type string within the interface map. This is done for later type
	// identification.
	StringPrefix = "string:"

	// StringSlicePrefix represents the prefix used to label keys and values
	// having the type []string within the interface map. This is done for later
	// type identification.
	StringSlicePrefix = "[]string:"
)

// Config represents the configuration used to create a new smart map object.
type Config struct {
	// StringMap provides a way to create a new smart map object out of a given
	// string map holding raw, as strings represented, information of arbitrary
	// objects.
	StringMap map[string]string
}

// DefaultConfig provides a default configuration to create a new smart map
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		StringMap: map[string]string{},
	}

	return newConfig
}

// NewSmartMap creates a new configured smart map object.
func NewSmartMap(config Config) (spec.SmartMap, error) {
	newSmartMap := &smartMap{
		Config: config,
		Mutex:  sync.Mutex{},
	}

	return newSmartMap, nil
}

type smartMap struct {
	Config

	Mutex sync.Mutex
}

func (sm *smartMap) GetStringMap() map[string]string {
	sm.Mutex.Lock()
	defer sm.Mutex.Unlock()

	return sm.StringMap
}

func (sm *smartMap) GetStringString(key string) (string, error) {
	sm.Mutex.Lock()
	defer sm.Mutex.Unlock()

	pk, err := withPrefix(key)
	if err != nil {
		return "", maskAny(err)
	}

	for k, v := range sm.StringMap {
		if k == pk {
			wo, err := withoutPrefix(v)
			if err != nil {
				return "", maskAny(err)
			}
			if s, ok := wo.(string); ok {
				return s, nil
			}
			return "", maskAnyf(wrongTypeError, "expected string got %T", wo)
		}
	}

	return "", maskAnyf(keyNotFoundError, key)
}

func (sm *smartMap) GetStringStringSlice(key string) ([]string, error) {
	sm.Mutex.Lock()
	defer sm.Mutex.Unlock()

	pk, err := withPrefix(key)
	if err != nil {
		return nil, maskAny(err)
	}

	for k, v := range sm.StringMap {
		if k == pk {
			wo, err := withoutPrefix(v)
			if err != nil {
				return nil, maskAny(err)
			}
			if ss, ok := wo.([]string); ok {
				return ss, nil
			}
			return nil, maskAnyf(wrongTypeError, "expected []string got %T", wo)
		}
	}

	return nil, maskAnyf(keyNotFoundError, key)
}

func (sm *smartMap) SetStringString(key, value string) error {
	sm.Mutex.Lock()
	defer sm.Mutex.Unlock()

	pk, err := withPrefix(key)
	if err != nil {
		return maskAny(err)
	}
	pv, err := withPrefix(value)
	if err != nil {
		return maskAny(err)
	}

	sm.StringMap[pk] = pv

	return nil
}

func (sm *smartMap) SetStringStringSlice(key string, value []string) error {
	sm.Mutex.Lock()
	defer sm.Mutex.Unlock()

	pk, err := withPrefix(key)
	if err != nil {
		return maskAny(err)
	}
	pv, err := withPrefix(value)
	if err != nil {
		return maskAny(err)
	}

	sm.StringMap[pk] = pv

	return nil
}
