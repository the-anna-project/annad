package key

import (
	"fmt"

	"github.com/xh3b4sd/anna/spec"
)

const (
	// CLGKeyFormat represents the format used to create storage keys for the
	// CLG scope.
	CLGKeyFormat = "s:clg:o:%s:%s"
)

// NewCLGKey returns a well configured key used to store and fetch data. Keys
// generated with NewCLGKey should only be used by objects related to the CLG
// scope. This can be e.g. each CLG. These objects generate and structure
// dynamic information used to learn. The returned key has the following
// scheme. "s" stands for the scope, that is, the CLG scope. "o" stands for the
// object requesting the key. "<key>" stands for the key-value pair identifying
// the most specific part of the key being used.
//
//     s:clg:o:<object>:<key>
//
func NewCLGKey(o, f string, v ...interface{}) string {
	return fmt.Sprintf(CLGKeyFormat, o, fmt.Sprintf(f, v...))
}

const (
	// CoreKeyFormat represents the format used to create storage keys for the
	// core's neural network scope.
	CoreKeyFormat = "s:core:o:%s:%s"
)

// NewCoreKey returns a well configured key used to store and fetch data. Keys
// generated with NewCoreKey should only be used by objects related to the
// core's neural network scope. This can be e.g. the core's neural network.
// These objects generate and structure dynamic information used to learn. The
// returned key has the following scheme. "s" stands for the scope, that is,
// the core scope.  "o" stands for the object requesting the key. "<key>"
// stands for the key-value pair identifying the most specific part of the key
// being used.
//
//     s:core:o:<object>:<key>
//
func NewCoreKey(o spec.Object, f string, v ...interface{}) string {
	return fmt.Sprintf(CoreKeyFormat, o.GetType(), fmt.Sprintf(f, v...))
}

const (
	// SysKeyFormat represents the format used to create storage keys for the
	// system scope.
	SysKeyFormat = "s:sys:o:%s:%s"
)

// NewSysKey returns a well configured key used to store and fetch data. Keys
// generated with NewSysKey should only be used by objects related to the
// system scope. This can be e.g. the Scheduler. These objects generate and
// structure fundamental information used to manage the system. The returned
// key has the following scheme. "s" stands for the scope, that is, the system
// scope. "o" stands for the object requesting the key. "<key>" stands for the
// key-value pair identifying the most specific part of the key being used.
//
//     s:sys:o:<object>:<key>
//
func NewSysKey(o spec.Object, f string, v ...interface{}) string {
	return fmt.Sprintf(SysKeyFormat, o.GetType(), fmt.Sprintf(f, v...))
}
