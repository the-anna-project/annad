// Package readseparator implements spec.CLG and provides functionality to
// read a separator stored in association to a specific behaviour ID.
package readseparator

import (
	"fmt"
	"strings"

	objectspec "github.com/the-anna-project/spec/object"
	storagecollection "github.com/the-anna-project/storage/collection"
)

func (s *service) calculate(ctx objectspec.Context) (string, error) {
	behaviourID, ok := ctx.GetBehaviourID()
	if !ok {
		return "", maskAnyf(invalidBehaviourIDError, "must not be empty")
	}

	behaviourIDKey := fmt.Sprintf("behaviour-id:%s:separator", behaviourID)
	separator, err := s.Service().Storage().General().Get(behaviourIDKey)
	if storagecollection.IsNotFound(err) {
		randomKey, err := s.Service().Storage().Feature().GetRandom()
		if err != nil {
			return "", maskAny(err)
		}

		// Make sure the fetched feature is valid based on its key structure. The
		// key itself already contains the feature. Knowing the key structure makes
		// it possible to simply parse the feature from the key fetched from the
		// feature storage. These features are stored by the split-features CLG.
		// Changes in the key structure there must be aligned with implementation
		// details here. The current key structure looks as follows.
		//
		//     feature:%s:positions
		//
		if len(randomKey) != 22 {
			return "", maskAnyf(invalidFeatureKeyError, "key '%s' must have length '22'", randomKey)
		}
		if !strings.HasPrefix(randomKey, "feature:") {
			return "", maskAnyf(invalidFeatureKeyError, "key '%s' prefix must be 'feature:'", randomKey)
		}
		if !strings.HasSuffix(randomKey, ":positions") {
			return "", maskAnyf(invalidFeatureKeyError, "key '%s' suffix must be ':positions'", randomKey)
		}

		// Create a new separator from the fetched random feature. Note that a
		// feature is considered 4 characters long and the random factory takes a
		// max parameter, which is exlusive.
		feature := randomKey[8:12]
		featureIndex, err := s.Service().Random().CreateMax(5)
		if err != nil {
			return "", maskAny(err)
		}
		separator = string(feature[featureIndex])

		// Store the newly created separator using the CLGs own behaviour ID. In case
		// this CLG is asked again to return its separator, it will lookup its
		// separator in the general storage.
		err = s.Service().Storage().General().Set(behaviourIDKey, separator)
		if err != nil {
			return "", maskAny(err)
		}
	} else if err != nil {
		return "", maskAny(err)
	}

	return separator, nil
}
