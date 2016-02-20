package state

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/spec"
)

func (s *state) UnmarshalJSON(bytes []byte) error {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 15}, "call UnmarshalJSON")

	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	var generic map[string]interface{}
	err := json.Unmarshal(bytes, &generic)
	if err != nil {
		return maskAny(err)
	}

	s.Bytes, err = s.unmarshalJSONBytes(generic)
	if err != nil {
		return maskAny(err)
	}

	s.Cores, err = s.unmarshalJSONCores(generic)
	if err != nil {
		return maskAny(err)
	}

	s.CreatedAt, err = s.unmarshalJSONCreatedAt(generic)
	if err != nil {
		return maskAny(err)
	}

	s.Impulses, err = s.unmarshalJSONImpulses(generic)
	if err != nil {
		return maskAny(err)
	}

	s.Networks, err = s.unmarshalJSONNetworks(generic)
	if err != nil {
		return maskAny(err)
	}

	s.Neurons, err = s.unmarshalJSONNeurons(generic)
	if err != nil {
		return maskAny(err)
	}

	s.ObjectID, err = s.unmarshalJSONObjectID(generic)
	if err != nil {
		return maskAny(err)
	}

	s.ObjectType, err = s.unmarshalJSONObjectType(generic)
	if err != nil {
		return maskAny(err)
	}

	s.StateReader, err = s.unmarshalJSONStateReader(generic)
	if err != nil {
		return maskAny(err)
	}

	s.StateWriter, err = s.unmarshalJSONStateWriter(generic)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *state) unmarshalJSONBytes(generic map[string]interface{}) (map[string][]byte, error) {
	newBytes := map[string][]byte{}

	if genericBytes, ok := generic["bytes"].(map[string]interface{}); ok {
		for key, genericByteSlice := range genericBytes {
			if base64Encoded, ok := genericByteSlice.(string); ok {
				base64Decoded, err := base64.StdEncoding.DecodeString(base64Encoded)
				if err != nil {
					return nil, maskAny(err)
				}
				newBytes[key] = []byte(base64Decoded)
			}
		}
	}

	return newBytes, nil
}

func (s *state) unmarshalJSONCores(generic map[string]interface{}) (map[spec.ObjectID]spec.Core, error) {
	newCores := map[spec.ObjectID]spec.Core{}

	if genericCores, ok := generic["cores"].(map[string]interface{}); ok {
		for genericObjectID, genericCore := range genericCores {
			coreBytes, err := json.Marshal(genericCore)
			if err != nil {
				return nil, maskAny(err)
			}
			coreSpec, err := s.FactoryClient.NewCore()
			if err != nil {
				return nil, maskAny(err)
			}
			err = json.Unmarshal(coreBytes, coreSpec)
			if err != nil {
				return nil, maskAny(err)
			}

			newCores[spec.ObjectID(genericObjectID)] = coreSpec
		}
	}

	return newCores, nil
}

func (s *state) unmarshalJSONCreatedAt(generic map[string]interface{}) (time.Time, error) {
	newCreatedAt := time.Time{}

	if genericCreatedAt, ok := generic["created_at"].(string); ok {
		t, err := time.Parse(time.RFC3339Nano, genericCreatedAt)
		if err != nil {
			return time.Time{}, maskAny(err)
		}
		newCreatedAt = t
	}

	return newCreatedAt, nil
}

func (s *state) unmarshalJSONImpulses(generic map[string]interface{}) (map[spec.ObjectID]spec.Impulse, error) {
	newImpulses := map[spec.ObjectID]spec.Impulse{}

	if genericImpulses, ok := generic["impulses"].(map[string]interface{}); ok {
		for genericObjectID, genericImpulse := range genericImpulses {
			impulseBytes, err := json.Marshal(genericImpulse)
			if err != nil {
				return nil, maskAny(err)
			}
			impulseSpec, err := s.FactoryClient.NewImpulse()
			if err != nil {
				return nil, maskAny(err)
			}
			err = json.Unmarshal(impulseBytes, impulseSpec)
			if err != nil {
				return nil, maskAny(err)
			}

			newImpulses[spec.ObjectID(genericObjectID)] = impulseSpec
		}
	}

	return newImpulses, nil
}

func (s *state) unmarshalJSONNetworks(generic map[string]interface{}) (map[spec.ObjectID]spec.Network, error) {
	newNetworks := map[spec.ObjectID]spec.Network{}

	if genericNetworks, ok := generic["networks"].(map[string]interface{}); ok {
		for genericObjectID, genericNetwork := range genericNetworks {
			networkBytes, err := json.Marshal(genericNetwork)
			if err != nil {
				return nil, maskAny(err)
			}
			networkSpec, err := s.FactoryClient.NewNetwork()
			if err != nil {
				return nil, maskAny(err)
			}
			err = json.Unmarshal(networkBytes, networkSpec)
			if err != nil {
				return nil, maskAny(err)
			}

			newNetworks[spec.ObjectID(genericObjectID)] = networkSpec
		}
	}

	return newNetworks, nil
}

func (s *state) unmarshalJSONNeurons(generic map[string]interface{}) (map[spec.ObjectID]spec.Neuron, error) {
	newNeurons := map[spec.ObjectID]spec.Neuron{}

	if genericNeurons, ok := generic["neurons"].(map[string]interface{}); ok {
		for genericObjectID, genericNeuron := range genericNeurons {
			objectType, err := objectTypeFromGenericObject(genericNeuron)
			if err != nil {
				return nil, maskAny(err)
			}

			neuronBytes, err := json.Marshal(genericNeuron)
			if err != nil {
				return nil, maskAny(err)
			}

			var neuronSpec spec.Neuron
			switch objectType {
			case common.ObjectType.CharacterNeuron:
				neuronSpec, err = s.FactoryClient.NewCharacterNeuron()
				if err != nil {
					return nil, maskAny(err)
				}
			case common.ObjectType.FirstNeuron:
				neuronSpec, err = s.FactoryClient.NewFirstNeuron()
				if err != nil {
					return nil, maskAny(err)
				}
			case common.ObjectType.JobNeuron:
				neuronSpec, err = s.FactoryClient.NewJobNeuron()
				if err != nil {
					return nil, maskAny(err)
				}
			default:
				return nil, maskAny(invalidObjectTypeError)
			}
			err = json.Unmarshal(neuronBytes, neuronSpec)
			if err != nil {
				return nil, maskAny(err)
			}

			newNeurons[spec.ObjectID(genericObjectID)] = neuronSpec
		}
	}

	return newNeurons, nil
}

func (s *state) unmarshalJSONObjectID(generic map[string]interface{}) (spec.ObjectID, error) {
	newObjectID := spec.ObjectID("")

	if genericObjectID, ok := generic["object_id"].(string); ok {
		newObjectID = spec.ObjectID(genericObjectID)
	}

	return newObjectID, nil
}

func (s *state) unmarshalJSONObjectType(generic map[string]interface{}) (spec.ObjectType, error) {
	newObjectType := spec.ObjectType("")

	if genericObjectType, ok := generic["object_type"].(string); ok {
		newObjectType = spec.ObjectType(genericObjectType)
	}

	return newObjectType, nil
}

func (s *state) unmarshalJSONStateReader(generic map[string]interface{}) (spec.StateType, error) {
	newStateReader := spec.StateType("")

	if genericStateReader, ok := generic["state_reader"].(string); ok {
		newStateReader = spec.StateType(genericStateReader)
	}

	return newStateReader, nil
}

func (s *state) unmarshalJSONStateWriter(generic map[string]interface{}) (spec.StateType, error) {
	newStateWriter := spec.StateType("")

	if genericStateWriter, ok := generic["state_writer"].(string); ok {
		newStateWriter = spec.StateType(genericStateWriter)
	}

	return newStateWriter, nil
}
