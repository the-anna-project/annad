package spec

// FeatureSet simply represents a feature container. This is basically the
// object created when analysing sequences. Based on the created feature
// distributions information will be stored.
type FeatureSet interface {
	// GetFeatures returns all features detected during scanning of sequences.
	GetFeatures() []Feature

	// GetFeaturesByCount returns all features that were found at least count
	// times.
	GetFeaturesByCount(count int) []Feature

	// GetFeaturesByLength returns all features which underlying sequences are at
	// least length characters long.
	GetFeaturesByLength(length int) []Feature

	// GetFeaturesBySequence returns all features that are represented by the
	// given sequence. Note this is usually only one feature.
	GetFeaturesBySequence(sequence string) []Feature

	// GetSequences simply returns the configured sequences of the current
	// feature set.
	GetSequences() []string

	// Scan analyses the given sequences to detect patterns. Found patterns are
	// used to create and attach new features. These features are represented by
	// their corresponding distribution which is basically used to store, compare
	// and learn.
	Scan() error
}
