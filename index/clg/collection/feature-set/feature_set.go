// Package featureset provides feature detection within sequences. A feature is
// considered a recognized property of a sequence. A sequence can be any
// string.
package featureset

import (
	"strings"
	"sync"

	"github.com/xh3b4sd/anna/service/id"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeFeatureSet represents the object type of the feature set object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeFeatureSet spec.ObjectType = "feature-set"
)

// Config represents the configuration used to create a new feature set object.
type Config struct {
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

// DefaultConfig provides a default configuration to create a new feature set
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		MaxLength: -1,
		MinLength: 1,
		MinCount:  1,
		Separator: "",
		Sequences: []string{},
	}

	return newConfig
}

// New creates a new configured feature set object. A feature set tries to
// detect all patterns within the configured input sequences.
func New(config Config) (spec.FeatureSet, error) {
	newIDService, err := id.NewService(id.DefaultServiceConfig())
	if err != nil {
		panic(err)
	}
	newID, err := newIDService.WithType(id.Hex128)
	if err != nil {
		panic(err)
	}

	newFeatureSet := &featureSet{
		Config: config,
		ID:     newID,
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeFeatureSet,
	}

	if newFeatureSet.MaxLength < -1 {
		return nil, maskAnyf(invalidConfigError, "MaxLength must be greater than -2")
	}
	if newFeatureSet.MinLength < 1 {
		return nil, maskAnyf(invalidConfigError, "MaxLength must be greater than 0")
	}
	if newFeatureSet.MaxLength != -1 && newFeatureSet.MaxLength < newFeatureSet.MinLength {
		return nil, maskAnyf(invalidConfigError, "MaxLength must be equal to or greater thanMinLength")
	}
	if newFeatureSet.MinCount < 0 {
		return nil, maskAnyf(invalidConfigError, "MinCount must be greater than -1")
	}
	if len(newFeatureSet.Sequences) == 0 {
		return nil, maskAnyf(invalidConfigError, "sequences must not be empty")
	}

	return newFeatureSet, nil
}

type featureSet struct {
	Config

	Features []spec.Feature
	ID       string
	Mutex    sync.Mutex
	Type     spec.ObjectType
}

func (fs *featureSet) GetFeatures() []spec.Feature {
	return fs.Features
}

func (fs *featureSet) GetFeaturesByCount(count int) []spec.Feature {
	var newFeatures []spec.Feature

	for _, f := range fs.Features {
		if f.GetCount() >= count {
			newFeatures = append(newFeatures, f)
		}
	}

	return newFeatures
}

func (fs *featureSet) GetFeaturesByLength(min, max int) []spec.Feature {
	var newFeatures []spec.Feature

	for _, f := range fs.Features {
		l := len(f.GetSequence())
		if l >= min && (l <= max || max == -1) {
			newFeatures = append(newFeatures, f)
		}
	}

	return newFeatures
}

func (fs *featureSet) GetFeaturesBySequence(sequence string) []spec.Feature {
	var newFeatures []spec.Feature

	for _, f := range fs.Features {
		if strings.Contains(f.GetSequence(), sequence) {
			newFeatures = append(newFeatures, f)
		}
	}

	return newFeatures
}

func (fs *featureSet) GetMaxLength() int {
	return fs.MaxLength
}

func (fs *featureSet) GetMinLength() int {
	return fs.MinLength
}

func (fs *featureSet) GetMinCount() int {
	return fs.MinCount
}

func (fs *featureSet) GetSeparator() string {
	return fs.Separator
}

func (fs *featureSet) GetSequences() []string {
	return fs.Sequences
}

func (fs *featureSet) Scan() error {
	// Prepare sequence combinations.
	var allSeqs []string
	for _, sequence := range fs.Sequences {
		for _, seq := range seqCombinations(sequence, fs.Separator, fs.MinLength, fs.MaxLength) {
			if !containsString(allSeqs, seq) {
				allSeqs = append(allSeqs, seq)
			}
		}
	}

	// Find sequence positions.
	positions := map[string][][]float64{}
	for _, sequence := range fs.Sequences {
		for _, seq := range allSeqs {
			if strings.Contains(sequence, seq) {
				if _, ok := positions[seq]; !ok {
					positions[seq] = [][]float64{}
				}
				positions[seq] = append(positions[seq], seqPositions(sequence, seq)...)
			}
		}
	}

	// Create features for each found sequence.
	var newFeatures []spec.Feature
	for seq, ps := range positions {
		if len(ps) < fs.MinCount {
			continue
		}

		newConfig := DefaultFeatureConfig()
		newConfig.Positions = ps
		newConfig.Sequence = seq
		newFeature, err := NewFeature(newConfig)
		if err != nil {
			return maskAny(err)
		}
		newFeatures = append(newFeatures, newFeature)
	}

	fs.Features = newFeatures

	return nil
}
