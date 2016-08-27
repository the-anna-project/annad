package log

import (
	"fmt"
	"time"

	"github.com/xh3b4sd/anna/spec"
)

const (
	// dateTimeMilli represents the date-time format used in log messages.
	dateTimeMilli = "06/01/02 15:04:05.000"
)

var (
	// LevelColors represent the log levels and their related colors the logger
	// implementations should use.
	LevelColors = map[string]string{
		"D": "cyan",
		"E": "red",
		"F": "magenta",
		"I": "white",
		"W": "yellow",
	}
)

func extendFormatWithTags(f string, tags spec.Tags) string {
	newFormat := ""

	newFormat += fmt.Sprintf("[%s] ", time.Now().Format(dateTimeMilli))
	if tags.L != "" {
		newFormat += fmt.Sprintf("[L: %s] ", tags.L)
	}
	if tags.O != nil {
		newFormat += fmt.Sprintf("[O: %s / %s] ", tags.O.GetType(), tags.O.GetID())
	}
	if tags.T != nil {
		newFormat += fmt.Sprintf("[T: %s] ", tags.T.GetID()) // TODO
	}
	newFormat += fmt.Sprintf("[V: %2d] ", tags.V)
	newFormat += f

	return newFormat
}

func containsString(list []string, item string) bool {
	for _, member := range list {
		if item == member {
			return true
		}
	}

	return false
}

func containsObjectType(list []spec.ObjectType, item spec.ObjectType) bool {
	for _, member := range list {
		if item == member {
			return true
		}
	}

	return false
}

func colorForLevel(level string) (string, error) {
	if c, ok := LevelColors[level]; ok {
		return c, nil
	}

	return "", maskAny(invalidLogLevelError)
}
