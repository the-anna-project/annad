package log

import (
	"github.com/xh3b4sd/anna/spec"
)

func (l *log) GetID() string {
	return l.ID
}

func (l *log) GetType() spec.ObjectType {
	return l.Type
}
