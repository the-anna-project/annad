package context

// WithValue enriches partent with the given key value pair. The extended
// context is returned. WithValue panics in case parent's underlying type is not
// the package's context type.
func WithValue(parent Context, k, v interface{}) Context {
	underlying := parent.(*context)
	underlying.storage[k] = v
	return underlying
}
