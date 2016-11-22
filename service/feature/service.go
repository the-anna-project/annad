// Package feature provides feature detection within sequences. A feature is
// considered a recognized property of a sequence. A sequence can be any string.
package feature

import (
	"strings"

	objectspec "github.com/the-anna-project/spec/object"
	servicespec "github.com/the-anna-project/spec/service"
	"github.com/the-anna-project/annad/object/feature"
)

// New creates a new feature service. The feature service tries to detect all
// patterns within the configured input sequences.
func New() servicespec.FeatureService {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	metadata map[string]string
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "feature",
		"type": "service",
	}
}

func (s *service) Metadata() map[string]string {
	return s.metadata
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

		newFeatures = append(newFeatures, newObject)
	}

	return newFeatures, nil
}

func (s *service) ScanConfig() servicespec.ScanConfig {
	newConfig := servicespec.ScanConfig{
		MaxLength: -1,
		MinLength: 1,
		MinCount:  1,
		Separator: "",
		Sequences: []string{},
	}

	return newConfig
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(sc servicespec.ServiceCollection) {
	s.serviceCollection = sc
}
