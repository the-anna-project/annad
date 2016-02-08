package state

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/spec"
)

//func (s *state) MarshalJSON() ([]byte, error) {
//	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 15}, "call UnmarshalJSON")
//	return nil, nil
//}

func (s *state) UnmarshalJSON(bytes []byte) error {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 15}, "call UnmarshalJSON")

	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	var generic map[string]interface{}
	err := json.Unmarshal(bytes, &generic)
	if err != nil {
		return maskAny(err)
	}

	// State.Bytes
	if genericBytes, ok := generic["bytes"].(map[string]interface{}); ok {
		for key, genericByteSlice := range genericBytes {
			if base64Encoded, ok := genericByteSlice.(string); ok {
				base64Decoded, err := base64.StdEncoding.DecodeString(base64Encoded)
				if err != nil {
					return maskAny(err)
				}
				s.Bytes[key] = []byte(base64Decoded)
			}
		}
	}

	// State.Cores
	if genericCores, ok := generic["cores"].(map[string]interface{}); ok {
		for genericObjectID, genericCore := range genericCores {
			coreBytes, err := json.Marshal(genericCore)
			if err != nil {
				return maskAny(err)
			}
			coreSpec, err := s.FactoryClient.NewCore()
			if err != nil {
				return maskAny(err)
			}
			err = json.Unmarshal(coreBytes, coreSpec)
			if err != nil {
				return maskAny(err)
			}

			s.Cores[spec.ObjectID(genericObjectID)] = coreSpec
		}
	}

	// State.CreatedAt
	if genericCreatedAt, ok := generic["created_at"].(string); ok {
		t, err := time.Parse(time.RFC3339Nano, genericCreatedAt)
		if err != nil {
			return maskAny(err)
		}
		s.CreatedAt = t
	}

	// State.Impulses
	if genericImpulses, ok := generic["impulses"].(map[string]interface{}); ok {
		for genericObjectID, genericImpulse := range genericImpulses {
			impulseBytes, err := json.Marshal(genericImpulse)
			if err != nil {
				return maskAny(err)
			}
			impulseSpec, err := s.FactoryClient.NewImpulse()
			if err != nil {
				return maskAny(err)
			}
			err = json.Unmarshal(impulseBytes, impulseSpec)
			if err != nil {
				return maskAny(err)
			}

			s.Impulses[spec.ObjectID(genericObjectID)] = impulseSpec
		}
	}

	// State.Networks
	if genericNetworks, ok := generic["networks"].(map[string]interface{}); ok {
		for genericObjectID, genericNetwork := range genericNetworks {
			networkBytes, err := json.Marshal(genericNetwork)
			if err != nil {
				return maskAny(err)
			}
			networkSpec, err := s.FactoryClient.NewNetwork()
			if err != nil {
				return maskAny(err)
			}
			err = json.Unmarshal(networkBytes, networkSpec)
			if err != nil {
				return maskAny(err)
			}

			s.Networks[spec.ObjectID(genericObjectID)] = networkSpec
		}
	}

	// State.Neurons
	if genericNeurons, ok := generic["neurons"].(map[string]interface{}); ok {
		for genericObjectID, genericNeuron := range genericNeurons {
			objectType, err := objectTypeFromGenericObject(genericNeuron)
			if err != nil {
				return maskAny(err)
			}

			neuronBytes, err := json.Marshal(genericNeuron)
			if err != nil {
				return maskAny(err)
			}

			var neuronSpec spec.Neuron
			switch objectType {
			case common.ObjectType.CharacterNeuron:
				neuronSpec, err = s.FactoryClient.NewCharacterNeuron()
				if err != nil {
					return maskAny(err)
				}
			case common.ObjectType.FirstNeuron:
				neuronSpec, err = s.FactoryClient.NewFirstNeuron()
				if err != nil {
					return maskAny(err)
				}
			case common.ObjectType.JobNeuron:
				neuronSpec, err = s.FactoryClient.NewJobNeuron()
				if err != nil {
					return maskAny(err)
				}
			default:
				return maskAny(invalidObjectTypeError)
			}
			err = json.Unmarshal(neuronBytes, neuronSpec)
			if err != nil {
				return maskAny(err)
			}

			s.Neurons[spec.ObjectID(genericObjectID)] = neuronSpec
		}
	}

	// State.ObjectID
	if genericObjectID, ok := generic["object_id"].(string); ok {
		s.ObjectID = spec.ObjectID(genericObjectID)
	}

	// State.ObjectType
	if genericObjectType, ok := generic["object_type"].(string); ok {
		s.ObjectType = spec.ObjectType(genericObjectType)
	}

	// State.StateReader
	if genericStateReader, ok := generic["state_reader"].(string); ok {
		s.StateReader = spec.StateType(genericStateReader)
	}

	// State.StateWriter
	if genericStateWriter, ok := generic["state_writer"].(string); ok {
		s.StateWriter = spec.StateType(genericStateWriter)
	}

	return nil
}
