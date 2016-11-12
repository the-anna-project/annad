// Package pairsyntactic implements spec.CLG and provides pairing of information
// sequences.
package pairsyntactic

import (
	"strings"

	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/object/spec"
	"github.com/xh3b4sd/anna/service/storage"
)

// TODO there is nothing that reads pairs
func (s *service) calculate(ctx spec.Context) error {
	// The counter keeps track of the work already being done. We only increment
	// the counter in case we were not able to do our job. As soon as some
	// threshold is reached, we stop trying.
	var counter int

	for {
		// Fetch two random features from the feature storage. This is done by
		// fetching random keys. The keys itself already contain the features. Knowing
		// the key structure makes it possible to simply parse the features from the
		// keys fetched from the feature storage. These features are stored by the
		// split-features CLG. Changes in the key structure there must be aligned with
		// implementation details here. The current key structure looks as follows.
		//
		//     feature:%s:positions
		//
		key1, err := s.Service().Storage().Feature().GetRandom()
		if err != nil {
			return maskAny(err)
		}
		key2, err := s.Service().Storage().Feature().GetRandom()
		if err != nil {
			return maskAny(err)
		}

		// Validate the fetched keys.
		for _, k := range []string{key1, key2} {
			if len(k) != 22 {
				return maskAnyf(invalidFeatureKeyError, "key '%s' must have length '22'", k)
			}
			if !strings.HasPrefix(key1, "feature:") {
				return maskAnyf(invalidFeatureKeyError, "key '%s' prefix must be 'feature:'", k)
			}
			if !strings.HasSuffix(key1, ":positions") {
				return maskAnyf(invalidFeatureKeyError, "key '%s' suffix must be ':positions'", k)
			}
		}

		// Combine the fetched keys to a new pair.
		pair := key1[8:12] + key2[8:12]

		// Write the new pair into the general storage.
		pairIDKey := key.NewNetworkKey("pair:syntactic:feature:%s:pair-id", pair)
		_, err = s.Service().Storage().General().Get(pairIDKey)
		if storage.IsNotFound(err) {
			// The created pair was not found within the feature storage. That means
			// we created a new one which we can store. Once we stored the new pair,
			// we break the outer loop to be done.
			newID, err := s.Service().ID().New()
			if err != nil {
				return maskAny(err)
			}
			pairID := string(newID)

			err = s.Service().Storage().General().Set(pairIDKey, pairID)
			if err != nil {
				return maskAny(err)
			}

			break
		}

		if counter >= 3 {
			// We tried to create a new pair 3 times, but it looks like there is not
			// enough data. That is why the created pair was found all the time. It
			// looks like we cannot create any new pair and stop doing anything here.
			break
		}

		counter++
	}

	return nil
}
