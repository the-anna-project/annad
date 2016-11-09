package main

import (
	"os"
	"os/signal"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func (a *annad) listenToSignal() {
	a.Service().Log().Line("func", "listenToSignal")

	listener := make(chan os.Signal, 2)
	signal.Notify(listener, os.Interrupt, os.Kill)

	<-listener

	go a.Shutdown()

	<-listener

	a.ForceShutdown()
}
