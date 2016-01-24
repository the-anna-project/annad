package id

import (
	"math/rand"
	"time"

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

	// Hex4096 creates a new hexa decimal encoded, pseudo random, 4096 bit hash.
	Hex4096 IDType = 512
)

var (
	randSrc   = rand.NewSource(time.Now().UnixNano())
	hashChars = "abcdef0123456789"
)

func randWithLen(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = hashChars[randSrc.Int63()%int64(len(hashChars))]
	}
	return string(b)
}

func NewID(idType IDType) spec.ObjectID {
	return spec.ObjectID(randWithLen(int(idType)))
}
