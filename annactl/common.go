package main

import (
	"os"
	"os/signal"

	"github.com/juju/errgo"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func (a *annactl) listenToSignal() {
	a.Service().Log().Line("func", "listenToSignal")

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
	a.Service().Log().Line("func", "GetSession")

	// Read session ID.
	raw, err := a.Service().FS().ReadFile(SessionFilePath)
	if _, ok := errgo.Cause(err).(*os.PathError); ok {
		// Session file does not exist. Go ahead to create one.
	} else if err != nil {
		return "", maskAny(err)
	}

	if len(raw) != 0 {
		return string(raw), nil
	}

	// Create session ID.
	newSessionID, err := a.Service().ID().New()
	if err != nil {
		return "", maskAny(err)
	}
	err = a.Service().FS().WriteFile(SessionFilePath, []byte(newSessionID), os.FileMode(0644))
	if err != nil {
		return "", maskAny(err)
	}

	return string(newSessionID), nil
}
