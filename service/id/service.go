// Package id provides a simple ID generating service using pseudo random
// strings.
package id

import "github.com/xh3b4sd/anna/service/spec"

const (
	// Hex128 creates a new hexa decimal encoded, pseudo random, 128 bit hash.
	Hex128 int = 16

	// Hex512 creates a new hexa decimal encoded, pseudo random, 512 bit hash.
	Hex512 int = 64

	// Hex1024 creates a new hexa decimal encoded, pseudo random, 1024 bit hash.
	Hex1024 int = 128

	// Hex2048 creates a new hexa decimal encoded, pseudo random, 2048 bit hash.
	Hex2048 int = 256

	// Hex4096 creates a new hexa decimal encoded, pseudo random, 4096 bit hash.
	Hex4096 int = 512
)

// New creates a new ID service.
func New() spec.ID {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection spec.Collection

	// Settings.

	// hashChars represents the characters used to create hashes.
	hashChars string
	idType    int
	metadata  map[string]string
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "id",
		"type": "service",
	}

	s.hashChars = "abcdef0123456789"
	s.idType = Hex128
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) New() (string, error) {
	ID, err := s.WithType(s.idType)
	if err != nil {
		return "", maskAny(err)
	}

	return ID, nil
}

func (s *service) Service() spec.Collection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(sc spec.Collection) {
	s.serviceCollection = sc
}

func (s *service) WithType(idType int) (string, error) {
	n := int(idType)
	max := len(s.hashChars)

	newRandomNumbers, err := s.Service().Random().CreateNMax(n, max)
	if err != nil {
		return "", maskAny(err)
	}

	b := make([]byte, n)

	for i, r := range newRandomNumbers {
		b[i] = s.hashChars[r]
	}

	return string(b), nil
}
