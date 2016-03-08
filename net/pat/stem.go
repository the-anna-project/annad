package patnet

// Stem returns the word stem all words provided by the given list have in
// common. Having the following list.
//
//     abc
//     abcd
//     abcde
//     abcdef
//
// Returns the word stem "abc".
//
func Stem(list []string) string {
	if len(list) == 0 {
		return ""
	}

	ri := 0
	li := 0
	ll := len(list)
	ref := list[0]
	rm := ""

	for {
		if ri > len(ref) {
			break
		}
		if ri > len(list[li]) {
			break
		}

		rm = ref[:ri]
		lm := list[li][:ri]

		if rm == lm {
			li++
			if li == ll {
				li = 0
				ri++
			}

			continue
		} else {
			break
		}
	}

	rm = ref[:ri-1]
	return rm
}
