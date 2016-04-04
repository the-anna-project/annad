package patnet

import (
	"strings"
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeFeatureSet represents the object type of the feature set object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeFeatureSet spec.ObjectType = "feature-set"
)

// FeatureSetConfig represents the configuration used to create a new feature
// set object.
type FeatureSetConfig struct {
	// MinLength represents the length of a sequence detected as feature. E.g.
	// MinLength set to 3 results in sequences having at least the length of 3
	// when detected features.
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

// DefaultFeatureSetConfig provides a default configuration to create a new
// feature set object by best effort.
func DefaultFeatureSetConfig() FeatureSetConfig {
	newConfig := FeatureSetConfig{
		MinLength: 1,
		MinCount:  1,
		Separator: "",
		Sequences: []string{},
	}

	return newConfig
}

// NewFeatureSet creates a new configured feature set object. A feature set
// tries to detect all patterns within the configured input sequences.
func NewFeatureSet(config FeatureSetConfig) (spec.FeatureSet, error) {
	newFeatureSet := &featureSet{
		FeatureSetConfig: config,
		ID:               id.NewObjectID(id.Hex128),
		Mutex:            sync.Mutex{},
		Type:             ObjectTypeFeatureSet,
	}

	if len(newFeatureSet.Sequences) == 0 {
		return nil, maskAnyf(invalidConfigError, "sequences must not be empty")
	}

	return newFeatureSet, nil
}

type featureSet struct {
	FeatureSetConfig

	Features []spec.Feature
	ID       spec.ObjectID
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
		if f.GetSequence() == sequence {
			newFeatures = append(newFeatures, f)
		}
	}

	return newFeatures
}

func (fs *featureSet) GetSequences() []string {
	return fs.Sequences
}

func (fs *featureSet) Scan() error {
	// Prepare sequence combinations.
	var allSeqs []string
	for _, sequence := range fs.Sequences {
		for _, seq := range seqCombinations(sequence, fs.Separator, fs.MinLength) {
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
