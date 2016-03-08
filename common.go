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
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call listenToSignal")

	listener := make(chan os.Signal, 1)
	signal.Notify(listener, os.Interrupt, os.Kill)

	<-listener

	a.Shutdown()
}

// writeStateInfo writes state information to the configured storage. The
// information look like this.
//
//     "time":    "16/03/27 21:14:35"
//     "version": "84ehdv0"
//
func (a *anna) writeStateInfo() {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call writeStateInfo")

	err := a.Storage.Set("version", version)
	panicOnError(err)

	for {
		dateTime := time.Now().Format("06/01/02 15:04:05")
		err := a.Storage.Set("time", dateTime)
		if err != nil {
			a.Log.WithTags(spec.Tags{L: "E", O: a, T: nil, V: 4}, "%#v", maskAny(err))
		}

		time.Sleep(5 * time.Second)
	}
}
