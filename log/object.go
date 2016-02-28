package log

import (
	"github.com/xh3b4sd/anna/spec"
)

func (l *log) GetID() spec.ObjectID {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()
	return l.ID
}

func (l *log) GetType() spec.ObjectType {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()
	return l.Type
}
