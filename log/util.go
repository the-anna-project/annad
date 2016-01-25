package log

import (
	"github.com/kdar/factorlog"
)

func OnlySeverity(severity factorlog.Severity) []factorlog.Severity {
	return []factorlog.Severity{severity, severity}
}

func wrapColorFormat(format string) string {
	colorFormat := `%{Color "red" "ERROR"}%{Color "yellow" "WARN"}%{Color "white" "INFO"}%{Color "cyan" "DEBUG"}%{Color "blue" "TRACE"}`
	colorFormat += format
	colorFormat += `%{Color "reset"}`
	return colorFormat
}
