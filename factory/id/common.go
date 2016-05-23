package id

import (
	"time"

	"github.com/xh3b4sd/anna/spec"
)

func (f *factory) retryErrorLogger(err error, d time.Duration) {
	if f.Log != nil {
		f.Log.WithTags(spec.Tags{L: "E", O: f, T: nil, V: 4}, "retry error: %#v", maskAny(err))
	}
}
