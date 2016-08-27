package redis

import (
	"time"

	"github.com/xh3b4sd/anna/spec"
)

func (s *storage) retryErrorLogger(err error, d time.Duration) {
	s.Log.WithTags(spec.Tags{C: nil, L: "E", O: s, V: 4}, "retry error: %#v", maskAny(err))
}
