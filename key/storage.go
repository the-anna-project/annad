package key

import (
	"fmt"
)

// TODO rename to NetworkKeyFormat, NewNetworkKey

const (
	// CLGKeyFormat represents the format used to create storage keys for the CLG
	// scope. "s" stands for the scope, that is, the CLG scope. "o" stands for
	// the object requesting the key. "<key>" stands for the key-value pair
	// identifying the most specific part of the key being used.
	CLGKeyFormat = "s:clg:%s"
)

// NewCLGKey returns a well configured key used to store and fetch data. Keys
// generated with NewCLGKey should only be used by objects related to the CLG
// scope. This can be e.g. each CLG. These objects generate and structure
// dynamic information used to learn. Note that there might be a lot of
// different keys used. In case this is necessary, a list of all used CLG key
// structures should be possible to be created by grepping all occurrences of
// the NewCLGKey function. See the documentation of the development cheat sheet
// for an example.
//
//     https://github.com/xh3b4sd/anna/blob/master/doc/development/cheat_sheet.md
//
func NewCLGKey(f string, v ...interface{}) string {
	return fmt.Sprintf(CLGKeyFormat, fmt.Sprintf(f, v...))
}
