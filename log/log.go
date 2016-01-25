package log

import (
	"os"

	"github.com/kdar/factorlog"

	"github.com/xh3b4sd/anna/spec"
)

type Config struct {
	// Color decides to whether color log output or not.
	Color bool

	// Format describes how to structure log output.
	Format string

	// SeverityRange defines the log severity type to output by a min and a max
	// value. Note that the list MUST provide 2 options. The first is the min,
	// the last is the max value. If you want to dedicate to one specific
	// severity type, just provide the same severity type twice. See also
	// OnlySeverity. By convention this can be a range between the following
	// options.
	//
	//   factorlog.DEBUG
	//   factorlog.INFO
	//   factorlog.WARN
	//   factorlog.ERROR
	//
	SeverityRange []factorlog.Severity

	// Verbosity describes the verbosity level. By convention this can be between
	// 0 and 12. Reason for that are the 4 conventional severity types. This
	// should help identifying and using the proper verbosity level for each
	// severity type. So you can use 3 verbosity levels for each severity type as
	// follows.
	//
	//         0  disable logging
	//    1 -  3  verbosity for factorlog.ERROR
	//    4 -  6  verbosity for factorlog.WARN
	//    6 -  9  verbosity for factorlog.INFO
	//   10 - 12  verbosity for factorlog.DEBUG
	//
	Verbosity int
}

func DefaultConfig() Config {
	newDefaultConfig := Config{
		Color:         true,
		Format:        "[%{S}] [%{Date} %{Time}] %{Message}",
		SeverityRange: []factorlog.Severity{factorlog.DEBUG, factorlog.ERROR},
		Verbosity:     12,
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

	newLog.Log.SetMinMaxSeverity(config.SeverityRange[0], config.SeverityRange[1])
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
