package state

import (
	"encoding/json"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/spec"
)

func (s *state) Read() error {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 14}, "call Read")

	version := s.GetVersion()

	switch s.StateReader {
	case common.StateType.FSReader:
		s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "restoring state backup from file")

		err := s.ReadFile(common.DefaultStateFile)
		if err != nil {
			return maskAny(err)
		}
	case common.StateType.NoneReader:
		// Do nothing.
		s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "NOT restoring state backup")
	default:
		return maskAny(invalidStateReaderError)
	}

	s.SetVersion(version)

	return nil
}

func (s *state) ReadFile(filename string) error {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 14}, "call ReadFile")

	bytes, err := s.FileSystem.ReadFile(filename)
	if err != nil {
		return maskAny(err)
	}

	err = json.Unmarshal(bytes, s)
	if err != nil {
		return maskAny(err)
	}

	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "state backup restored")

	return nil
}
