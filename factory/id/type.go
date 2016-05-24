package id

import (
	"github.com/xh3b4sd/anna/spec"
)

const (
	// Hex128 creates a new hexa decimal encoded, pseudo random, 128 bit hash.
	Hex128 spec.IDType = 16

	// Hex512 creates a new hexa decimal encoded, pseudo random, 512 bit hash.
	Hex512 spec.IDType = 64

	// Hex1024 creates a new hexa decimal encoded, pseudo random, 1024 bit hash.
	Hex1024 spec.IDType = 128

	// Hex2048 creates a new hexa decimal encoded, pseudo random, 2048 bit hash.
	Hex2048 spec.IDType = 256

	// Hex4096 creates a new hexa decimal encoded, pseudo random, 4096 bit hash.
	Hex4096 spec.IDType = 512
)
