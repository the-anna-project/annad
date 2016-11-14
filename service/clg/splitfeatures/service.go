// Package splitfeatures implements spec.CLG and provides functionality to
// split information sequences into features.
//
// One possible CLG tree might look like described below. Here the output of the
// two upper CLGs are used to feed the lower CLG.
//
//    |---------------------|     |---------------------|
//    | read-information-id |     |    read-separator   |
//    |---------------------|     |---------------------|
//               |                           |
//               -----------------------------
//                             |
//                             V
//                     |----------------|
//                     | split-features |
//                     |----------------|
//
package splitfeatures

import (
	"encoding/json"
	"fmt"

	objectspec "github.com/xh3b4sd/anna/object/spec"
)

const (
	// FeatureSize represents the number of characters a feature consists of. E.g.
	// a FeatureSize of 4 results in features being registered which are 4
	// characters long. Features are detected by a window sliding over an
	// information sequence. Once a feature is detected the window slides one
	// character farther.
	FeatureSize int = 4
)

func (s *service) calculate(ctx objectspec.Context, informationSequence, separator string) error {
	newConfig := s.Service().Feature().ScanConfig()
	newConfig.MaxLength = FeatureSize
	newConfig.MinLength = FeatureSize
	newConfig.Separator = separator
	newConfig.Sequences = []string{informationSequence}
	newFeatures, err := s.Service().Feature().Scan(newConfig)
	if err != nil {
		return maskAny(err)
	}

	for _, f := range newFeatures {
		// Store the detected feature within the feature storage. It is important to
		// preserve the key structure used here to simply parse the stored features
		// in other places, like in the pair-syntactic and read-separator CLG.
		// Changes in the key structure there must be aligned with implementation
		// details here. The current key structure looks as follows.
		//
		//     feature:%s:positions
		//
		positionKey := fmt.Sprintf("feature:%s:positions", f.Sequence())
		raw, err := json.Marshal(f.Positions())
		if err != nil {
			return maskAny(err)
		}
		err = s.Service().Storage().Feature().Set(positionKey, string(raw))
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}
