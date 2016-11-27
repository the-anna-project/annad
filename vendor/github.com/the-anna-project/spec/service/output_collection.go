package service

// OutputCollection represents a collection of output services. This scopes
// different output service implementations in a simple container, which can
// easily be passed around.
type OutputCollection interface {
	Boot()
	SetTextService(textService OutputService)
	Text() OutputService
}
