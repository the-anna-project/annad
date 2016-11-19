package spec

// InputCollection represents a collection of input services. This scopes
// different input service implementations in a simple container, which can
// easily be passed around.
type InputCollection interface {
	Boot()
	SetTextService(textService InputService)
	Text() InputService
}
