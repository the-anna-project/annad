package key

import (
	"strings"
)

// NewPromKey returns a new key, used for prometheus instrumentation, having
// all given parts properly joined using underscores. The schema of the
// returned key looks as follows.
//
//     s_s_s_s
//
func NewPromKey(s ...string) string {
	return strings.Join(s, "_")
}
