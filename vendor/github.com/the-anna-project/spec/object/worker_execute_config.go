package object

type WorkerExecuteConfig interface {
	Actions() []func(canceler <-chan struct{}) error
	Canceler() chan struct{}
	CancelOnError() bool
	NumWorkers() int
	SetActions(actions []func(canceler <-chan struct{}) error)
	SetCanceler(canceler chan struct{})
	SetCancelOnError(cancelOnError bool)
	SetNumWorkers(numWorkers int)
}
