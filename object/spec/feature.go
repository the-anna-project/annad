package spec

// Feature represents a charactistic within a sequence. During pattern
// recognition it is tried to detect features. Their distributions describe
// location patterns within space.
type Feature interface {
	Configure() error
	// AddPosition provides a way to add more positions to the initialized
	// feature. Note positions are vectors in distribution terms.
	AddPosition(position []float64) error
	// Count returns the number of occurrences within analysed sequences. That is,
	// how often this feature was found. Technically spoken,
	// len(feature.Positions).
	Count() int
	// Positions returns the feature's configured positions.
	Positions() [][]float64
	// Sequence returns the sequence that represents this feature. This is the
	// sub-sequence, the charactistic detected within analysed sequences.
	Sequence() string
	SetPositions(ps [][]float64)
	SetSequence(s string)

}
