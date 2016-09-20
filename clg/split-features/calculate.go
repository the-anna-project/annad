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

	"github.com/xh3b4sd/anna/index/clg/collection/feature-set"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// FeatureSize represents the number of characters a feature consists of. E.g.
	// a FeatureSize of 4 results in features being registered which are 4
	// characters long. Features are detected by a window sliding over an
	// information sequence. Once a feature is detected the window slides one
	// character farther.
	FeatureSize int = 4
)

// TODO seperator CLG
func (c *clg) calculate(ctx spec.Context, informationSequence, separator string) error {
	newConfig := featureset.DefaultConfig()
	newConfig.MaxLength = FeatureSize
	newConfig.MinLength = FeatureSize
	newConfig.Separator = separator
	newConfig.Sequences = []string{informationSequence}
	newFeatureSet, err := featureset.New(newConfig)
	if err != nil {
		return maskAny(err)
	}

	err = newFeatureSet.Scan()
	if err != nil {
		return maskAny(err)
	}

	features := newFeatureSet.GetFeatures()
	for _, f := range features {
		// Store the detected feature within the feature storage. It is important to
		// preserve the key structure used here to simply parse the stored features
		// in other places, like in the pair-syntactic CLG. Changes in the key
		// structure there must be aligned with implementation details here. The
		// current key structure looks as follows.
		//
		//     feature:%s:positions
		//
		positionKey := key.NewCLGKey("feature:%s:positions", f.GetSequence())
		raw, err := json.Marshal(f.GetPositions())
		if err != nil {
			return maskAny(err)
		}
		err = c.Storage().Feature().Set(positionKey, string(raw))
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}
