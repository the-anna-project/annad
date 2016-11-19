package spec

// InputCollection represents a collection of endpoint instances. This scopes
// different endpoint implementations in a simple container, which can easily be
// passed around.
type InputCollection interface {
	Boot()
	SetTextService(textService InputService)
	Text() InputService
}
