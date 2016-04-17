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

	// GetFeaturesByLength returns all features which underlying sequences have a
	// character length between min and max.
	GetFeaturesByLength(min, max int) []Feature

	// GetFeaturesBySequence returns all features that are represented by the
	// given sequence. Note this is usually only one feature.
	GetFeaturesBySequence(sequence string) []Feature

	// GetMaxLength returns the feature set's MaxLength configuration.
	GetMaxLength() int

	// GetMinLength returns the feature set's MinLength configuration.
	GetMinLength() int

	// GetMinCount returns the feature set's MinCount configuration.
	GetMinCount() int

	// GetSeparator returns the feature set's Separator configuration.
	GetSeparator() string

	// GetSequences simply returns the configured sequences of the current
	// feature set.
	GetSequences() []string

	// Scan analyses the given sequences to detect patterns. Found patterns are
	// used to create and attach new features. These features are represented by
	// their corresponding distribution which is basically used to store, compare
	// and learn.
	Scan() error
}
