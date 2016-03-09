// Package id provides simple ID generation using pseudo randomly generated
// strings.
package id

import (
	"crypto/rand"
	"math/big"

	"github.com/xh3b4sd/anna/spec"
)

// Type represents some kind of configuration for ID creation.
type Type int

const (
	// Hex128 creates a new hexa decimal encoded, pseudo random, 128 bit hash.
	Hex128 Type = 16

	// Hex512 creates a new hexa decimal encoded, pseudo random, 512 bit hash.
	Hex512 Type = 64

	// Hex1024 creates a new hexa decimal encoded, pseudo random, 1024 bit hash.
	Hex1024 Type = 128

	// Hex2048 creates a new hexa decimal encoded, pseudo random, 2048 bit hash.
	Hex2048 Type = 256

	// Hex4096 creates a new hexa decimal encoded, pseudo random, 4096 bit hash.
	Hex4096 Type = 512
)

var (
	hashChars = "abcdef0123456789"
)

// NewObjectID creates a new object ID for the given type.
func NewObjectID(t Type) spec.ObjectID {
	b := make([]byte, int(t))

	for i := range b {
		max := big.NewInt(int64(len(hashChars)))
		j, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(err)
		}

		b[i] = hashChars[int(j.Int64())]
	}

	return spec.ObjectID(b)
}
