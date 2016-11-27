package service

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// FeatureService represents a service being able to scan for features within
// information sequences.
type FeatureService interface {
	Boot()
	Metadata() map[string]string
	// Scan analyses the given sequences to detect patterns. Found patterns are
	// returned in form of a list of features.
	Scan(config ScanConfig) ([]objectspec.Feature, error)
	// ScanConfig returns a default scan config configured by best effort.
	ScanConfig() ScanConfig
	Service() ServiceCollection
	SetServiceCollection(serviceCollection ServiceCollection)
}

// TODO provide interface and move implementation to feature service and remove error.go
// ScanConfig represents the configuration used to scan for new feature objects.
type ScanConfig struct {
	// MaxLength represents the length maximum of a sequence detected as feature.
	// E.g. MaxLength set to 3 results in sequences having a length not larger
	// than 3 when detected as features. The value -1 disables any limitiation.
	MaxLength int
	// MinLength represents the minimum length of a sequence detected as feature.
	// E.g. MinLength set to 3 results in sequences having a length not smaller
	// than 3 when detected as features. The value -1 disables any limitiation.
	MinLength int
	// MinCount represents the number of occurrences at least required to be
	// detected as feature. E.g. MinCount set to 3 requires a feature to be
	// present within a given input sequence at least 3 times.
	MinCount int
	// Separator represents the separator used to split sequences into smaller
	// parts. By default this is an empty string resulting in a character split.
	// This can be set to a whitespace to split for words. Note that the concept
	// of words is a feature known to humans based on contextual information
	// humans connected to create reasonable sences. This needs to be achieved by
	// Anna herself. So later this separator needs to be configured by Anna once
	// she is able to recognize improvements in feature detection, resulting in
	// even more awareness of contextual information.
	Separator string
	// Sequences represents the input sequences being analysed. Out of this
	// information features are detected, if any.
	Sequences []string
}

// Validate checks whether ScanConfig is valid for proper execution in
// Feature.Scan.
func (sc *ScanConfig) Validate() error {
	// Settings.

	if sc.MaxLength < -1 {
		return maskAnyf(invalidConfigError, "max length must be greater than -2")
	}
	if sc.MinLength < 1 {
		return maskAnyf(invalidConfigError, "max length must be greater than 0")
	}
	if sc.MaxLength != -1 && sc.MaxLength < sc.MinLength {
		return maskAnyf(invalidConfigError, "max length must be equal to or greater than min length")
	}
	if sc.MinCount < 0 {
		return maskAnyf(invalidConfigError, "min count must be greater than -1")
	}
	if len(sc.Sequences) == 0 {
		return maskAnyf(invalidConfigError, "sequences must not be empty")
	}

	return nil
}
