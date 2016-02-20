package state

import (
	"encoding/json"

	"github.com/xh3b4sd/anna/spec"
)

func (s *state) SetStateFromObjectBytes(bytes []byte) error {
	var generic map[string]interface{}
	err := json.Unmarshal(bytes, &generic)
	if err != nil {
		return maskAny(err)
	}
	stateBytes, err := json.Marshal(generic["state"])
	if err != nil {
		return maskAny(err)
	}
	err = json.Unmarshal(stateBytes, s)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func objectTypeFromGenericObject(v interface{}) (spec.ObjectType, error) {
	if genericObject, ok := v.(map[string]interface{}); ok {
		if genericState, ok := genericObject["state"].(map[string]interface{}); ok {
			if genericObjectType, ok := genericState["object_type"].(string); ok {
				return spec.ObjectType(genericObjectType), nil
			}
		}
	}

	return "", maskAny(invalidObjectTypeError)
}
