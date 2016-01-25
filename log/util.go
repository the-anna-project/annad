package log

import (
	"github.com/kdar/factorlog"
)

func OnlyLevel(level Level) []factorlog.Severity {
	return []factorlog.Severity{levelToSeverity(level), levelToSeverity(level)}
}

func wrapColorFormat(format string) string {
	colorFormat := `%{Color "red" "ERROR"}%{Color "yellow" "WARN"}%{Color "white" "INFO"}%{Color "cyan" "DEBUG"}%{Color "blue" "TRACE"}`
	colorFormat += format
	colorFormat += `%{Color "reset"}`
	return colorFormat
}

func levelToSeverity(level Level) factorlog.Severity {
	switch level {
	case Debug:
		return factorlog.DEBUG
	case Info:
		return factorlog.INFO
	case Warn:
		return factorlog.WARN
	case Error:
		return factorlog.ERROR
	default:
		return factorlog.DEBUG
	}
}
