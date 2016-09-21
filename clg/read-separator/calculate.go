// Package readseparator implements spec.CLG and provides functionality to
// read a separator stored in association to a specific behavior ID.
package readseparator

import (
	"strings"

	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
)

func (c *clg) calculate(ctx spec.Context) (string, error) {
	behaviorID := ctx.GetBehaviorID()
	if behaviorID == "" {
		return "", maskAnyf(invalidBehaviorIDError, "must not be empty")
	}

	behaviorIDKey := key.NewCLGKey("behavior-id:%s:separator", behaviorID)
	separator, err := c.Storage().General().Get(behaviorIDKey)
	if storage.IsNotFound(err) {
		randomKey, err := c.Storage().Feature().GetRandomKey()
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
		// max parameter as second argument, which is exlusive.
		feature := randomKey[8:12]
		numbers, err := c.Factory().Random().CreateNMax(1, 5)
		if err != nil {
			return "", maskAny(err)
		}
		separator = string(feature[numbers[0]])

		// Store the newly created separator using the CLGs own behavior ID. In case
		// this CLG is asked again to return its separator, it will lookup its
		// separator in the general storage.
		err = c.Storage().General().Set(behaviorIDKey, separator)
		if err != nil {
			return "", maskAny(err)
		}
	} else if err != nil {
		return "", maskAny(err)
	}

	return separator, nil
}
