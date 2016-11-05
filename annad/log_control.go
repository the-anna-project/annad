package main

import (
	"github.com/xh3b4sd/anna/server/control/log"
	"github.com/xh3b4sd/anna/spec"
)

func newLogControl(newLog spec.Log) (spec.LogControl, error) {
	newLogControlConfig := log.DefaultControlConfig()
	newLogControlConfig.Log = newLog
	newLogControl, err := log.NewControl(newLogControlConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newLogControl, nil
}
