package hash

import (
	"math/rand"
	"time"
)

var (
	src         = rand.NewSource(time.Now().UnixNano())
	hashChars   = "abcdef0123456789"
	hashIdxBits = 6                        // 6 bits to represent a char index
	hashIdxMask = 1<<uint(hashIdxBits) - 1 // All 1-bits, as many as hashIdxBits
	hashIdxMax  = 63 / hashIdxBits         // # of char indices fitting in 63 bits
)

// See http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func randWithLen(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for hashIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), hashIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), hashIdxMax
		}
		if idx := int(cache & int64(hashIdxMask)); idx < len(hashChars) {
			b[i] = hashChars[idx]
			i--
		}
		cache >>= uint(hashIdxBits)
		remain--
	}

	return string(b)
}

// Hex4096 creates a new hexa decimal encoded, pseudo random, 4096 bit hash.
func Hex4096() string {
	return randWithLen(512)
}

// Hex2048 creates a new hexa decimal encoded, pseudo random, 2048 bit hash.
func Hex2048() string {
	return randWithLen(256)
}

// Hex128 creates a new hexa decimal encoded, pseudo random, 128 bit hash.
func Hex128() string {
	return randWithLen(16)
}
