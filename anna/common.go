package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/xh3b4sd/anna/spec"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func (a *anna) listenToSignal() {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call listenToSignal")

	listener := make(chan os.Signal, 2)
	signal.Notify(listener, os.Interrupt, os.Kill)

	<-listener

	go a.Shutdown()

	<-listener

	a.ForceShutdown()
}

// writeStateInfo writes state information to the configured general storage.
// The information look like this.
//
//     "time":       "16/03/27 21:14:35"
//     "version":    "84ehdv0"
//
func (a *anna) writeStateInfo() {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call writeStateInfo")

	err := a.GeneralStorage.Set("version", version)
	panicOnError(err)

	for {
		dateTime := time.Now().Format("06/01/02 15:04:05")
		err := a.GeneralStorage.Set("time", dateTime)
		if err != nil {
			a.Log.WithTags(spec.Tags{C: nil, L: "E", O: a, V: 4}, "%#v", maskAny(err))
		}

		time.Sleep(5 * time.Second)
	}
}
