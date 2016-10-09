package key

import (
	"fmt"
)

const (
	// NetworkKeyPrefix represents the prefix used to create storage keys for the
	// network scope. "s" stands for the scope, that is, the network scope.
	NetworkKeyPrefix = "s:network"
)

// NewNetworkKey returns a well defined key used to identify data in some
// underlying storage. Keys generated with NewNetworkKey should only be used by
// objects related to the network scope. This can be e.g. the neural network
// itself, or even each CLG acting inside of the neural network. All of these
// objects generate and structure dynamic information that are used to learn.
// Note that there might be a lot of different keys used. In case it is
// necessary, a list of all used network key structures should be possible to be
// created by grepping all occurrences of the NewNetworkKey function. See the
// documentation of the development cheat sheet for an example.
//
//     https://github.com/xh3b4sd/anna/blob/master/doc/development/cheat_sheet.md#list-storage-keys
//
func NewNetworkKey(f string, v ...interface{}) string {
	return NetworkKeyPrefix + ":" + fmt.Sprintf(f, v...)
}
