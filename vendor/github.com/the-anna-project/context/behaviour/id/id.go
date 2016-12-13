// Package id stores and accesses a behaviour ID from and in a
// github.com/the-anna-project/context.Context.
package id

import (
	"github.com/the-anna-project/context"
)

// key is an unexported type for keys defined in this package. This prevents
// collisions with keys defined in other packages.
type key string

// idKey is the key for behaviour ID values in
// github.com/the-anna-project/context.Context. Clients use id.NewContext and
// id.FromContext instead of using this key directly.
var idKey key = "id"

// NewContext returns a new github.com/the-anna-project/context.Context that
// carries value v.
func NewContext(ctx context.Context, v string) context.Context {
	if v == "" {
		// In case the given value is empty we do not add it, but only return the
		// given context as it is. That way the existence check when reading the
		// context works as expected when no value or an empty value was tried to be
		// added.
		return ctx
	}

	return context.WithValue(ctx, idKey, v)
}

// FromContext returns the behaviour ID value stored in ctx, if any.
func FromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(idKey).(string)
	return v, ok
}
