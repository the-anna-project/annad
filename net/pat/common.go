package patnet

import (
	"strconv"
	"strings"
)

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

func mapVectorsToChannels(vectors [][]float64, channels []float64) []float64 {
	// Collect all channel range indizes for each vector.
	cri := make([][]int, len(vectors))
	for index, vector := range vectors {
		for i, channel := range channels {
			var prevChan float64
			if i > 0 {
				prevChan = channels[i-1]
			}

			if betweenFloat64(prevChan, minFloat64(vector), maxFloat64(vector)) || betweenFloat64(channel, minFloat64(vector), maxFloat64(vector)) {
				cri[index] = append(cri[index], i)
				continue
			}

			for _, d := range vector {
				if betweenFloat64(d, prevChan, channel) {
					cri[index] = append(cri[index], i)
					break
				}
			}
		}
	}

	// Calculate the distribution relative for each channel. Each vector has a
	// weight of 1. Is a vector located within multiple channels, its weight is
	// divided across them.
	mapping := make([]float64, len(channels))
	for _, vs := range cri {
		t := 1 / float64(len(vs))

		for _, vi := range vs {
			mapping[vi] += t
		}
	}

	return mapping
}

func channelDistance(perc1, perc2 []float64) []float64 {
	var distance []float64

	for index := range perc1 {
		distance = append(distance, perc2[index]-perc1[index])
	}

	return distance
}

func uniqueFloat64(list []float64) bool {
	for _, i1 := range list {
		var c int
		for _, i2 := range list {
			if i1 == i2 {
				c++
			}
		}
		if c > 1 {
			return false
		}
	}

	return true
}

func maxFloat64(list []float64) float64 {
	if len(list) == 0 {
		return 0
	}

	max := list[0]

	for _, i := range list {
		if i > max {
			max = i
		}
	}

	return max
}

func minFloat64(list []float64) float64 {
	if len(list) == 0 {
		return 0
	}

	min := list[0]

	for _, i := range list {
		if i < min {
			min = i
		}
	}

	return min
}

func betweenFloat64(i, min, max float64) bool {
	if i < min {
		return false
	}
	if i > max {
		return false
	}
	return true
}

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

func staticChannelsFromString(value string) ([]float64, error) {
	var newStaticChannels []float64

	for _, c := range strings.Split(value, ",") {
		f, err := strconv.ParseFloat(c, 64)
		if err != nil {
			return nil, maskAny(err)
		}
		newStaticChannels = append(newStaticChannels, f)
	}

	return newStaticChannels, nil
}

func vectorsFromString(value string) ([][]float64, error) {
	var newVectors [][]float64

	for _, v := range strings.Split(value, "|") {
		var newVector []float64
		for _, d := range strings.Split(v, ",") {
			f, err := strconv.ParseFloat(d, 64)
			if err != nil {
				return nil, maskAny(err)
			}
			newVector = append(newVector, f)
		}
		newVectors = append(newVectors, newVector)
	}

	return newVectors, nil
}
