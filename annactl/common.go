package main

import (
	"os"
	"os/signal"

	"github.com/juju/errgo"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/spec"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func (a *annactl) listenToSignal() {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call listenToSignal")

	listener := make(chan os.Signal, 1)
	signal.Notify(listener, os.Interrupt, os.Kill)

	<-listener

	a.Shutdown()
}

const (
	// SessionFilePath represents the file path used to read and write session IDs.
	SessionFilePath = ".annasession"
)

func (a *annactl) GetSessionID() (string, error) {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call GetSession")

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
	newSessionID, err := a.IDFactory.WithType(id.Hex128)
	if err != nil {
		return "", maskAny(err)
	}
	err = a.FileSystem.WriteFile(SessionFilePath, []byte(newSessionID), os.FileMode(0644))
	if err != nil {
		return "", maskAny(err)
	}

	return string(newSessionID), nil
}
