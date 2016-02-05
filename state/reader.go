package state

import (
	"encoding/json"

	"github.com/xh3b4sd/anna/common"
)

func (s *state) Read() error {
	switch s.StateReader {
	case common.StateType.FSReader:
		err := s.ReadFile(common.DefaultStateFile)
		if err != nil {
			return maskAny(err)
		}
	case common.StateType.NoneReader:
		// Do nothing.
	default:
		return maskAny(invalidStateReaderError)
	}

	return nil
}

func (s *state) ReadFile(filename string) error {
	bytes, err := s.FileSystem.ReadFile(filename)
	if err != nil {
		return maskAny(err)
	}

	err = json.Unmarshal(bytes, s)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
