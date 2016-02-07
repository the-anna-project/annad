// Package log implements spec.Log. This logger interface is to simply log
// output to gather runtime information.
package log

import (
	"os"

	"github.com/kdar/factorlog"

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

	// LevelRange defines the log level to output by a min and a max value. Note
	// that the list MUST provide 2 options. The first is the min, the last is
	// the max value. If you want to dedicate to one specific log level, just
	// provide the same log level type twice. See also OnlyLevel. By convention
	// this can be a range between the following options.
	//
	//   D
	//   I
	//   W
	//   E
	//   F
	//
	LevelRange []string

	// Verbosity describes the log verbosity. By convention this can be between 0
	// and 15. Reason for that are the 5 conventional log levels. This should
	// help identifying and using the proper log verbosity for each log level. So
	// you can use 3 log verbosities for each log level as follows.
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
		Format:     "%{Message}",
		LevelRange: []string{"D", "F"},
		Verbosity:  15,
	}

	return newDefaultConfig
}

// NewLog creates a new basic logger. Logging is important to comprehensible
// track runtime information.
func NewLog(config Config) spec.Log {
	newFormat := config.Format
	if config.Color {
		newFormat = wrapColorFormat(newFormat)
	}

	newLog := log{
		Config: config,
		Logger: factorlog.New(os.Stdout, factorlog.NewStdFormatter(newFormat)),
	}

	newLog.Logger.SetMinMaxSeverity(levelToSeverity(config.LevelRange[0]), levelToSeverity(config.LevelRange[1]))
	newLog.Logger.SetVerbosity(factorlog.Level(config.Verbosity))

	return &newLog
}

type log struct {
	Config

	Logger *factorlog.FactorLog
}

func (l *log) WithTags(tags spec.Tags, f string, v ...interface{}) {
	switch tags.L {
	case "D":
		l.Logger.V(factorlog.Level(tags.V)).Debugf(extendFormatWithTags(f, tags), v...)
	case "E":
		l.Logger.V(factorlog.Level(tags.V)).Errorf(extendFormatWithTags(f, tags), v...)
	case "F":
		l.Logger.V(factorlog.Level(tags.V)).Fatalf(extendFormatWithTags(f, tags), v...)
	case "I":
		l.Logger.V(factorlog.Level(tags.V)).Infof(extendFormatWithTags(f, tags), v...)
	case "W":
		l.Logger.V(factorlog.Level(tags.V)).Warnf(extendFormatWithTags(f, tags), v...)
	default:
		l.Logger.V(factorlog.Level(tags.V)).Debugf(extendFormatWithTags(f, tags), v...)
	}
}
