// Package feature provides feature detection within sequences. A feature is
// considered a recognized property of a sequence. A sequence can be any string.
package feature

import (
	"strings"

	"github.com/xh3b4sd/anna/object/feature"
	objectspec "github.com/xh3b4sd/anna/object/spec"
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// New creates a new feature service. The feature service tries to detect all
// patterns within the configured input sequences.
func New() servicespec.Feature {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.Collection

	// Settings.

	metadata map[string]string
}

func (s *service) Configure() error {
	// Settings.

	id, err := s.Service().ID().New()
	if err != nil {
		return maskAny(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "feature",
		"type": "service",
	}

	return nil
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

// DefaultScanConfig provides a default configuration to scan for new feature
// objects by best effort.
func DefaultScanConfig() servicespec.ScanConfig {
	newConfig := servicespec.ScanConfig{
		MaxLength: -1,
		MinLength: 1,
		MinCount:  1,
		Separator: "",
		Sequences: []string{},
	}

	return newConfig
}

func (s *service) Scan(config servicespec.ScanConfig) ([]objectspec.Feature, error) {
	// Validate.
	err := config.Validate()
	if err != nil {
		return nil, maskAny(err)
	}

	// Prepare sequence combinations.
	var allSeqs []string
	for _, sequence := range config.Sequences {
		for _, seq := range seqCombinations(sequence, config.Separator, config.MinLength, config.MaxLength) {
			if !containsString(allSeqs, seq) {
				allSeqs = append(allSeqs, seq)
			}
		}
	}

	// Find sequence positions.
	positions := map[string][][]float64{}
	for _, sequence := range config.Sequences {
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
	var newFeatures []objectspec.Feature
	for seq, ps := range positions {
		if len(ps) < config.MinCount {
			continue
		}

		newObject := feature.New()
		newObject.SetPositions(ps)
		newObject.SetSequence(seq)
		err := newObject.Validate()
		if err != nil {
			return nil, maskAny(err)
		}
		err = newObject.Configure()
		if err != nil {
			return nil, maskAny(err)
		}

		newFeatures = append(newFeatures, newObject)
	}

	return newFeatures, nil
}

func (s *service) Service() servicespec.Collection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(sc servicespec.Collection) {
	s.serviceCollection = sc
}

func (s *service) Validate() error {
	// Dependencies.

	if s.serviceCollection == nil {
		return maskAnyf(invalidConfigError, "service collection must not be empty")
	}

	return nil
}
