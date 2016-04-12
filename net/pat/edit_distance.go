package patnet

// EditDistance implementes the Levenshtein distance to measure similarity
// between two strings. Here all edit operations are weighted with the cost 1.
// See http://en.wikipedia.org/wiki/Levenshtein_distance. The following code is
// a golang port of the optimized C version of
// http://en.wikibooks.org/wiki/Algorithm_implementation/Strings/Levenshtein_distance#C.
func EditDistance(s1, s2 string) int {
	var cost, lastdiag, olddiag int
	ls1 := len([]rune(s1))
	ls2 := len([]rune(s2))

	column := make([]int, ls1+1)

	for i := 1; i <= ls1; i++ {
		column[i] = i
	}

	for i := 1; i <= ls2; i++ {
		column[0] = i
		lastdiag = i - 1

		for j := 1; j <= ls1; j++ {
			olddiag = column[j]

			cost = 0
			if s1[j-1] != s2[i-1] {
				cost = 1
			}

			column[j] = distanceMin(column[j]+1, column[j-1]+1, lastdiag+cost)
			lastdiag = olddiag
		}
	}

	return column[ls1]
}

func distanceMin(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}
