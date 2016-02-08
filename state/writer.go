package state

import (
	"encoding/json"
	"os"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/spec"
)

func (s *state) Write() error {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 14}, "call Write")

	switch s.StateWriter {
	case common.StateType.FSWriter:
		s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "backing up state to file")

		err := s.WriteFile(common.DefaultStateFile)
		if err != nil {
			return maskAny(err)
		}
	case common.StateType.NoneWriter:
		// Do nothing.
		s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "NOT backing up state")
	default:
		return maskAny(invalidStateWriterError)
	}

	return nil
}

func (s *state) WriteFile(filename string) error {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 14}, "call WriteFile")

	bytes, err := json.Marshal(s)
	if err != nil {
		return maskAny(err)
	}

	err = s.FileSystem.WriteFile(filename, bytes, os.FileMode(0660)) // ug+rw (user and group can read and write)
	if err != nil {
		return maskAny(err)
	}
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "state backed up")

	return nil
}
