package object

type WorkerExecuteConfig interface {
	Action() func(canceler <-chan struct{}) error
	Canceler() chan struct{}
	CancelOnError() bool
	NumWorkers() int
	SetAction(action func(canceler <-chan struct{}) error)
	SetCanceler(canceler chan struct{})
	SetCancelOnError(cancelOnError bool)
	SetNumWorkers(numWorkers int)
}
