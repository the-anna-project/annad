package log

import (
	"fmt"
	"time"

	"github.com/kdar/factorlog"

	"github.com/xh3b4sd/anna/spec"
)

const (
	RFC3339NanoAligned = "2006-01-02T15:04:05.000000000Z07:00"
)

func OnlyLevel(level string) []factorlog.Severity {
	return []factorlog.Severity{levelToSeverity(level), levelToSeverity(level)}
}

func extendFormatWithTags(f string, tags spec.Tags) string {
	newFormat := ""

	newFormat += fmt.Sprintf("[D: %s] ", time.Now().Format(RFC3339NanoAligned))
	if tags.L != "" {
		newFormat += fmt.Sprintf("[L: %s] ", tags.L)
	}
	if tags.O != nil {
		newFormat += fmt.Sprintf("[O: %s / %s] ", tags.O.GetObjectType(), tags.O.GetObjectID())
	}
	if tags.T != nil {
		newFormat += fmt.Sprintf("[T: %s] ", tags.T.GetTraceID())
	}
	newFormat += fmt.Sprintf("[V: %2d] ", tags.V)
	newFormat += f

	return newFormat
}

func wrapColorFormat(format string) string {
	colorFormat := `%{Color "cyan" "DEBUG"}%{Color "red" "ERROR"}%{Color "magenta" "FATAL"}%{Color "white" "INFO"}%{Color "yellow" "WARN"}`
	colorFormat += format
	colorFormat += `%{Color "reset"}`
	return colorFormat
}

func levelToSeverity(level string) factorlog.Severity {
	switch level {
	case "D":
		return factorlog.DEBUG
	case "E":
		return factorlog.ERROR
	case "F":
		return factorlog.FATAL
	case "I":
		return factorlog.INFO
	case "W":
		return factorlog.WARN
	default:
		return factorlog.DEBUG
	}
}
