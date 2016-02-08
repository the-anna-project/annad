package log

import (
	"fmt"
	"time"

	"github.com/xh3b4sd/anna/spec"
)

const (
	DateTimeMilli = "06/01/02 15:04:05.000"
)

func extendFormatWithTags(f string, tags spec.Tags) string {
	newFormat := ""

	newFormat += fmt.Sprintf("[%s] ", time.Now().Format(DateTimeMilli))
	if tags.L != "" {
		newFormat += fmt.Sprintf("[L: %s] ", tags.L)
	}
	if tags.O != nil {
		newFormat += fmt.Sprintf("[O: %-16s / %s] ", tags.O.GetObjectType(), tags.O.GetObjectID())
	}
	if tags.T != nil {
		newFormat += fmt.Sprintf("[T: %s] ", tags.T.GetTraceID())
	}
	newFormat += fmt.Sprintf("[V: %2d] ", tags.V)
	newFormat += f

	return newFormat
}

func contains(list []string, item string) bool {
	for _, member := range list {
		if item == member {
			return true
		}
	}

	return false
}
