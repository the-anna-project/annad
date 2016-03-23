package spec

// Feature represents a charactistic within a sequence. During pattern
// recognition it is tried to detect features. Their distributions describe
// location patterns within space.
type Feature interface {
	// AddPosition provides a way to add more positions to the initialized
	// feature. Note positions are vectors in distribution terms.
	AddPosition(position []float64) error

	// GetCount returns the number of occurrences within analysed sequences. That
	// is, how often this feature was found.
	GetCount() int

	// GetDistribution returns the distribution representing this feature. See
	// documentation about the Distribution object for more information.
	GetDistribution() Distribution

	// GetPositions returns the feature's configured positions.
	GetPositions() [][]float64

	// GetSequence returns the sequence that represents this feature. This is the
	// sub-sequence, the charactistic detected within analysed sequences.
	GetSequence() string
}
