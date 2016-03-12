// Package key has the authority about key structures and provides simple
// functions to create keys for different purposes.
package key

import (
	"fmt"

	"github.com/xh3b4sd/anna/spec"
)

const (
	// NetKeyFormat represents the format used to create storage keys for the
	// network scope.
	NetKeyFormat = "s:net:%s:%s"
)

// NewNetKey returns a well configured key used to store and fetch data. Keys
// generated with NewNetKey should only be used by objects related to the
// neural network scope. This can be e.g. the CharNet, or the StratNet. These
// objects generate and structure dynamic information used to learn. The
// returned key has the following scheme. "s" stands for the scope, that is,
// the network scope. "o" stands for the object requesting the key. "<key>"
// stands for the key-value pair identifying the most specific part of the key.
//
//     s:net:o:<object>:<key>
//
func NewNetKey(o spec.Object, f string, v ...interface{}) string {
	return fmt.Sprintf(NetKeyFormat, o.GetType(), fmt.Sprintf(f, v...))
}

const (
	// SysKeyFormat represents the format used to create storage keys for the
	// system scope.
	SysKeyFormat = "s:sys:%s:%s"
)

// SysKey returns a well configured key used to store and fetch data. Keys
// generated with NewSysKey should only be used by objects related to the
// system scope. This can be e.g. the Scheduler. These objects generate and
// structure fundamental information used to manage the system. The returned
// key has the following scheme. "s" stands for the scope, that is, the system
// scope. "o" stands for the object requesting the key. "<key>" stands for the
// key-value pair identifying the most specific part of the key.
//
//     s:sys:o:<object>:<key>
//
func NewSysKey(o spec.Object, f string, v ...interface{}) string {
	return fmt.Sprintf(SysKeyFormat, o.GetType(), fmt.Sprintf(f, v...))
}
