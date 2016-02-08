// Package log implements spec.Log. This logger interface is to simply log
// output to gather runtime information.
package log

import (
	"fmt"
	builtinLog "log"
	"os"

	"github.com/mgutz/ansi"

	"github.com/xh3b4sd/anna/spec"
)

type Config struct {
	// Color decides to whether color log output or not.
	Color bool

	// Format describes how to structure log output. The log output format should
	// be simple and clean. In the first iteration the log format looks like
	// this.
	//
	//   [short severity] [yyyy-mm-dd hh:mm:ss] message
	//
	// For example a log line then looks like this.
	//
	//   [I] [2016-01-26 23:37:03] hello, I am Anna
	//
	Format string

	// Levels is used to only log messages emitted with a level matching one of
	// the levels given here.
	//
	//   D
	//   I
	//   W
	//   E
	//   F
	//
	Levels []string

	// ObjectType is used to only log messages emitted by objects matching this
	// given object type.
	ObjectType spec.ObjectType

	// TraceID is used to only log messages emitted by requests matching this
	// given trace ID.
	TraceID spec.TraceID

	// Verbosity is used to only log messages emitted with this given verbosity.
	// By convention this can be between 0 and 15. Reason for that are the 5
	// conventional log levels. This should help identifying and using the proper
	// log verbosity for each log level. So you can use 3 log verbosities for
	// each log level as follows.
	//
	//         0  disable logging
	//    1 -  3  log level F
	//    4 -  6  log level E
	//    7 -  9  log level W
	//   10 - 12  log level I
	//   13 - 15  log level D
	//
	Verbosity int
}

func DefaultConfig() Config {
	newDefaultConfig := Config{
		Color:      true,
		Levels:     []string{"D", "E", "F", "I", "W"},
		ObjectType: spec.ObjectType(""),
		TraceID:    spec.TraceID(""),
		Verbosity:  15,
	}

	return newDefaultConfig
}

// NewLog creates a new basic logger. Logging is important to comprehensible
// track runtime information.
func NewLog(config Config) spec.Log {
	newLog := log{
		Config: config,
		Logger: builtinLog.New(os.Stdout, "", 0),
	}

	return &newLog
}

type log struct {
	Config

	Logger spec.RootLogger
}

func (l *log) WithTags(tags spec.Tags, f string, v ...interface{}) {
	if !contains(l.Levels, tags.L) {
		return
	}

	if tags.O != nil && l.ObjectType != spec.ObjectType("") {
		if tags.O.GetObjectType() != l.ObjectType {
			return
		}
	}

	if tags.T != nil && l.TraceID != spec.TraceID("") {
		if tags.T.GetTraceID() != l.TraceID {
			return
		}
	}

	if tags.V == 0 || tags.V > l.Verbosity {
		return
	}

	color := "cyan"
	switch tags.L {
	case "D":
		color = "cyan"
	case "E":
		color = "red"
	case "F":
		color = "magenta"
	case "I":
		color = "white"
	case "W":
		color = "yellow"
	}

	msg := fmt.Sprintf(extendFormatWithTags(f, tags), v...)
	if l.Color {
		msg = ansi.Color(msg, color)
	}

	l.Logger.Println(msg)
}
