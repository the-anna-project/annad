// Package id provides simple ID generation using pseudo randomly generated
// strings.
package id

import (
	"crypto/rand"
	"math/big"

	"github.com/xh3b4sd/anna/spec"
)

type IDType int

const (
	// Hex128 creates a new hexa decimal encoded, pseudo random, 128 bit hash.
	Hex128 IDType = 16

	// Hex512 creates a new hexa decimal encoded, pseudo random, 512 bit hash.
	Hex512 IDType = 64

	// Hex1024 creates a new hexa decimal encoded, pseudo random, 1024 bit hash.
	Hex1024 IDType = 128

	// Hex2048 creates a new hexa decimal encoded, pseudo random, 2048 bit hash.
	Hex2048 IDType = 256

	// Hex4096Bit creates a new hexa decimal encoded, pseudo random, 4096 bit hash.
	Hex4096 IDType = 512
)

var (
	hashChars = "abcdef0123456789"
)

func NewObjectID(idType IDType) spec.ObjectID {
	b := make([]byte, int(idType))
	for i := range b {
		max := big.NewInt(int64(len(hashChars)))
		j, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(err)
		}
		b[i] = hashChars[j.Int64()]
	}

	return spec.ObjectID(b)
}
