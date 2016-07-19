package connectionpath

func equalDimensionLength(vectors [][]float64) bool {
	if len(vectors) == 0 {
		return false
	}

	l := len(vectors[0])

	for _, v := range vectors {
		if len(v) != l {
			return false
		}
	}

	return true
}
