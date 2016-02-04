package state

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/xh3b4sd/anna/common"
)

func (s *state) Write() error {
	switch s.StateWriter {
	case common.StateType.FSWriter:
		err := s.WriteFile(common.DefaultStateFile)
		if err != nil {
			return maskAny(err)
		}
	case common.StateType.NoneWriter:
		// Do nothing.
	default:
		return maskAny(invalidStateWriterError)
	}

	return nil
}

func (s *state) WriteFile(filename string) error {
	bytes, err := json.Marshal(s)
	if err != nil {
		return maskAny(err)
	}

	err = ioutil.WriteFile(filename, bytes, os.FileMode(0660)) // ug+rw (user and group can read and write)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
