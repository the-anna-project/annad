package feature

import (
	"strings"
)

func seqCombinations(sequence, separator string, minLength, maxLength int) []string {
	var combs []string
	splitted := strings.Split(sequence, separator)

	if matchesLength(sequence, minLength, maxLength) {
		combs = append(combs, sequence)
	}

	for i := range splitted {
		comb := splitted[i]
		if !containsString(combs, comb) && matchesLength(comb, minLength, maxLength) {
			combs = append(combs, comb)
		}

		j := i
		for range splitted {
			j++

			if j > len(splitted) {
				break
			}

			comb := strings.Join(splitted[i:j], separator)
			if !containsString(combs, comb) && matchesLength(comb, minLength, maxLength) {
				combs = append(combs, comb)
			}
		}
	}

	return combs
}

func matchesLength(seq string, min, max int) bool {
	if min != -1 && len(seq) < min {
		return false
	}
	if max != -1 && len(seq) > max {
		return false
	}
	return true
}

func seqPositions(sequence, seq string) [][]float64 {
	var positions [][]float64
	ll := float64(len(sequence))
	rdsf := 100 / ll
	rdef := float64(len(seq)) * 100 / ll

	for i := range sequence {
		if strings.HasPrefix(sequence[i:], seq) {
			// Position must be represented by its relative dimensions. When trying
			// to find out if a period is always at the end of a sentence, the
			// feature position needs to reflect that. E.g. there are sentences that
			// are 10 or 100 characters long. The feature detected at the end of such
			// sequences need to be represented by a position indicating 100 percent.
			rds := float64(i) * rdsf
			rde := rds + rdef
			positions = append(positions, []float64{rds, rde})
		}
	}

	return positions
}

func containsString(combs []string, comb string) bool {
	for _, c := range combs {
		if c == comb {
			return true
		}
	}

	return false
}
