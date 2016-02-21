package main

import (
	"os"
	"os/signal"

	"github.com/xh3b4sd/anna/spec"
)

func (a *anna) listenToSignal() {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call listenToSignal")

	listener := make(chan os.Signal, 1)
	signal.Notify(listener, os.Interrupt, os.Kill)

	<-listener

	a.Shutdown()
}
