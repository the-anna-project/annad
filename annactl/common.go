package main

import (
	"os"

	"github.com/juju/errgo"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/spec"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	// SessionFilePath represents the file path used to read and write session IDs.
	SessionFilePath = ".annasession"
)

func (a *annactl) GetSessionID() (string, error) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call GetSession")

	// Read session ID.
	raw, err := a.FileSystem.ReadFile(SessionFilePath)
	if _, ok := errgo.Cause(err).(*os.PathError); ok {
		// Session file does not exist. Go ahead to create one.
	} else if err != nil {
		return "", maskAny(err)
	}

	if len(raw) != 0 {
		return string(raw), nil
	}

	// Create session ID.
	newSessionID := id.NewObjectID(id.Hex128)
	err = a.FileSystem.WriteFile(SessionFilePath, []byte(newSessionID), os.FileMode(0644))
	if err != nil {
		return "", maskAny(err)
	}

	return string(newSessionID), nil
}
