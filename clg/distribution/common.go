package distribution

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
