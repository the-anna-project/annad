package permutation

func shiftIndizes(indizes []int, base int) ([]int, bool) {
	var msdShifted bool
	var reminder int

	for i := len(indizes) - 1; i >= 0; i-- {
		if reminder > 0 {
			current := indizes[i] + reminder
			reminder = 0
			indizes[i] = current
		}

		if indizes[i] >= base {
			indizes[i] = 0
			reminder = 1

			if i == 0 {
				msdShifted = true
			}
		}
	}

	return indizes, msdShifted
}

func prepend(s []int, i, x int) []int {
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = x

	return s
}
