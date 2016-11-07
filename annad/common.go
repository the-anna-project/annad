package main

import (
	"os"
	"os/signal"

	"github.com/xh3b4sd/anna/spec"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func (a *annad) listenToSignal() {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call listenToSignal")

	listener := make(chan os.Signal, 2)
	signal.Notify(listener, os.Interrupt, os.Kill)

	<-listener

	go a.Shutdown()

	<-listener

	a.ForceShutdown()
}
