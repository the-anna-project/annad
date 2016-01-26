package log

import (
	"os"

	"github.com/kdar/factorlog"

	"github.com/xh3b4sd/anna/spec"
)

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
	Warn  Level = "warn"
	Error Level = "error"
)

type Config struct {
	// Color decides to whether color log output or not.
	Color bool

	// Format describes how to structure log output.
	Format string

	// LevelRange defines the log level to output by a min and a max value. Note
	// that the list MUST provide 2 options. The first is the min, the last is
	// the max value. If you want to dedicate to one specific log level, just
	// provide the same log level type twice. See also OnlyLevel. By convention
	// this can be a range between the following options.
	//
	//   Debug
	//   Info
	//   Warn
	//   Error
	//
	LevelRange []Level

	// Verbosity describes the log verbosity. By convention this can be between 0
	// and 12. Reason for that are the 4 conventional severity types. This should
	// help identifying and using the proper log verbosity for each severity
	// type. So you can use 3 log verbosities for each log level as follows.
	//
	//         0  disable logging
	//    1 -  3  log level Error
	//    4 -  6  log level Warn
	//    7 -  9  log level Info
	//   10 - 12  log level Debug
	//
	Verbosity int
}

func DefaultConfig() Config {
	newDefaultConfig := Config{
		Color:      true,
		Format:     "[%{S}] [%{Date} %{Time}] %{Message}",
		LevelRange: []Level{Debug, Error},
		Verbosity:  12,
	}

	return newDefaultConfig
}

func NewLog(config Config) spec.Log {
	newFormat := config.Format
	if config.Color {
		newFormat = wrapColorFormat(newFormat)
	}

	newLog := log{
		Config: config,
		Log:    factorlog.New(os.Stdout, factorlog.NewStdFormatter(newFormat)),
	}

	newLog.Log.SetMinMaxSeverity(levelToSeverity(config.LevelRange[0]), levelToSeverity(config.LevelRange[1]))
	newLog.Log.SetVerbosity(factorlog.Level(config.Verbosity))

	return newLog
}

type log struct {
	Config

	Log *factorlog.FactorLog
}

func (l log) V(verbosity int) spec.Severity {
	return l.Log.V(factorlog.Level(verbosity))
}

func (l log) Debug(v ...interface{}) {
	l.Log.Debug(v)
}

func (l log) Debugf(format string, v ...interface{}) {
	l.Log.Debugf(format, v)
}

func (l log) Error(v ...interface{}) {
	l.Log.Error(v)
}

func (l log) Errorf(format string, v ...interface{}) {
	l.Log.Errorf(format, v)
}

func (l log) Info(v ...interface{}) {
	l.Log.Info(v)
}

func (l log) Infof(format string, v ...interface{}) {
	l.Log.Infof(format, v)
}

func (l log) Warn(v ...interface{}) {
	l.Log.Warn(v)
}

func (l log) Warnf(format string, v ...interface{}) {
	l.Log.Warnf(format, v)
}
