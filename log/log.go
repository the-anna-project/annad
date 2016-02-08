// Package log implements spec.Log. This logger interface is to simply log
// output to gather runtime information.
package log

import (
	"fmt"
	builtinLog "log"
	"os"
	"strings"
	"sync"

	"github.com/mgutz/ansi"

	"github.com/xh3b4sd/anna/common"
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

	// ObjectTypes is used to only log messages emitted by objects matching this
	// given object type.
	ObjectTypes []spec.ObjectType

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
		Color:       true,
		Levels:      []string{},
		ObjectTypes: []spec.ObjectType{},
		TraceID:     spec.TraceID(""),
		Verbosity:   10,
	}

	return newDefaultConfig
}

// NewLog creates a new basic logger. Logging is important to comprehensible
// track runtime information.
func NewLog(config Config) spec.Log {
	newLog := log{
		Config: config,
		Logger: builtinLog.New(os.Stdout, "", 0),
		Mutex:  sync.Mutex{},
	}

	return &newLog
}

type log struct {
	Config

	Logger spec.RootLogger
	Mutex  sync.Mutex
}

func (l *log) ResetLevels() {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	l.Levels = DefaultConfig().Levels
}

func (l *log) ResetObjectTypes() {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	l.ObjectTypes = DefaultConfig().ObjectTypes
}

func (l *log) ResetVerbosity() {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	l.Verbosity = DefaultConfig().Verbosity
}

func (l *log) SetLevels(list string) {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	if list == "" {
		return
	}

	newLevels := []string{}
	for _, level := range strings.Split(list, ",") {
		// We only use that here for level validation.
		_, err := colorForLevel(level)
		if err != nil {
			l.WithTags(spec.Tags{L: "E", O: l, T: nil, V: 4}, "%#v", maskAnyf(err, level))
			continue
		}

		newLevels = append(newLevels, level)
	}

	l.Levels = newLevels
}

func (l *log) SetObjectTypes(list string) {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	if list == "" {
		return
	}

	newObjectTypes := []spec.ObjectType{}
	for _, objectType := range strings.Split(list, ",") {
		if !containsObjectType(common.ObjectTypes, spec.ObjectType(objectType)) {
			l.WithTags(spec.Tags{L: "E", O: l, T: nil, V: 4}, "invalid objectType '%s'", objectType)
			continue
		}

		newObjectTypes = append(newObjectTypes, spec.ObjectType(objectType))
	}

	l.ObjectTypes = newObjectTypes
}

func (l *log) SetVerbosity(verbosity int) {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	l.Verbosity = verbosity
}

func (l *log) WithTags(tags spec.Tags, f string, v ...interface{}) {
	if len(l.Levels) != 0 && !containsString(l.Levels, tags.L) {
		return
	}

	if tags.O != nil && len(l.ObjectTypes) != 0 {
		if !containsObjectType(l.ObjectTypes, tags.O.GetObjectType()) {
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

	msg := fmt.Sprintf(extendFormatWithTags(f, tags), v...)
	if l.Color {
		color, err := colorForLevel(tags.L)
		if err != nil {
			l.WithTags(spec.Tags{L: "E", O: l, T: nil, V: 4}, "%#v", maskAnyf(err, tags.L))
			return
		}
		msg = ansi.Color(msg, color)
	}

	l.Logger.Println(msg)
}
