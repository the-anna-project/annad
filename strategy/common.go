package strategy

import (
	"crypto/rand"
	"math/big"
)

const (
	// clgNameDummy is simply a dummy CLG name injected during randomization
	// of the action list. See documentations below for more information.
	clgNameDummy = "dummy"
)

// randomizeCLGNames generates a random sequence using the given CLG names.
// Note that randomizing a strategy's action items MUST only be done when
// creating a new strategy. Further randomizations of existing strategies will
// cause the algorhythms the strategy network implements to fail.
//
// The following algorhythm is implemented as follows. Consider this given list
// of available action items.
//
//   a,b,c,d,e
//
// This are some possible combinations resulting out of the randomization.
//
//   c,e
//   b,b,d
//   a,b,a
//   d
//
func randomizeCLGNames(clgNames []string) []string {
	var newCLGNames []string

	if len(clgNames) == 0 {
		// In case there is no useful input given we simply return an empty list.
		// This also prevents a dead lock in the loops below.
		return newCLGNames
	}

	// The trick to randomize the given set of CLG names is to inject a well
	// known item that can be chosen and then ignored.
	options := append([]string{clgNameDummy}, clgNames...)

	for {
		for range clgNames {
			max := big.NewInt(int64(len(options)))
			i, err := rand.Int(rand.Reader, max)
			if err != nil {
				panic(err)
			}
			newOption := options[i.Int64()]

			if newOption == clgNameDummy {
				// There was a random index that chose the item we want to ignore. Thus
				// we do so. This results in combinations not necessarily having the same
				// length as the original given list of CLG names.
				continue
			}

			newCLGNames = append(newCLGNames, newOption)
		}

		if len(newCLGNames) == 0 {
			continue
		}

		break
	}

	return newCLGNames
}
