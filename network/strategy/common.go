package strategynetwork

import (
	"crypto/rand"
	"math/big"
)

func randomMinMax(min, max int) int {
	bigMin := big.NewInt(int64(min))
	bigMax := big.NewInt(int64(max))

	r, err := rand.Int(rand.Reader, bigMax.Sub(bigMax, bigMin))
	if err != nil {
		panic(err)
	}

	return int(r.Add(r, bigMin).Int64())
}
