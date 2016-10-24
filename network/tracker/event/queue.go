package event

import (
	"strings"
	"sync"
	"sync/atomic"
)

// QueueConfig represents the configuration used to create a new queue object.
type QueueConfig struct {
	// Destination represents the destination of the network payload currently
	// being processed.
	Destination string

	// Source represents the sources of the network payload currently being
	// processed.
	Sources []string
}

// DefaultConfig provides a default configuration to create a new queue object
// by best effort.
func DefaultConfig() QueueConfig {
	newConfig := QueueConfig{
		Destination: "",
		Sources:     nil,
	}

	return newConfig
}

// TODO comment
type Queue interface {
	// TODO comment
	// Boot initializes and starts the whole network like booting a machine. The
	// call to Boot blocks until the network is completely initialized, so you
	// might want to call it in a separate goroutine.
	//
	// Boot makes the network listen on requests from the outside. Here each
	// CLG input channel is managed. This way Listen acts as kind of cortex in
	// which signals are dispatched into all possible direction and finally flow
	// back again. Errors during processing of the neural network will be logged
	// to the provided logger.
	Boot()

	GetComplete() chan struct{}

	GetError() chan error

	GetInput() chan string

	GetOutput() chan Event

	// Shutdown ends all processes of the network like shutting down a machine.
	// The call to Shutdown blocks until the network is completely shut down, so
	// you might want to call it in a separate goroutine.
	Shutdown()
}

// NewQueue creates a new configured queue object.
func NewQueue(config QueueConfig) (Queue, error) {
	newQueue := &queue{
		QueueConfig: config,

		BootOnce:          sync.Once{},
		Closer:            make(chan struct{}, 1),
		Complete:          make(chan struct{}, 1),
		ConnectionDetails: nil,
		Error:             make(chan error, 1),
		Input:             make(chan string, 1),
		ProcessedEvents:   0,
		Output:            make(chan Event, 1),
		ShutdownOnce:      sync.Once{},
	}

	if newQueue.Destination == "" {
		return nil, maskAnyf(invalidConfigError, "destination must not be empty")
	}
	if newQueue.Sources == nil {
		return nil, maskAnyf(invalidConfigError, "sources must not be empty")
	}

	return newQueue, nil
}

type ConnectionDetail struct {
	Connection  string
	Destination string
	Source      string
}

type queue struct {
	QueueConfig

	BootOnce          sync.Once
	Closer            chan struct{}
	Complete          chan struct{}
	ConnectionDetails []ConnectionDetail
	Error             chan error
	Input             chan string
	Output            chan Event
	ProcessedEvents   int32
	ShutdownOnce      sync.Once
}

// TODO comment
func (q *queue) Boot() {
	q.BootOnce.Do(func() {
		// TODO comment
		for _, s := range q.Sources {
			cd := ConnectionDetail{
				Connection:  s + "," + q.Destination,
				Destination: q.Destination,
				Source:      s,
			}
			q.ConnectionDetails = append(q.ConnectionDetails, cd)
		}

		// TODO comment
		for {
			select {
			case <-q.Closer:
				return
			case <-q.Complete:
				return
			case connectionPath := <-q.Input:
				events, err := q.Process(connectionPath)
				if err != nil {
					q.Error <- maskAny(err)
					close(q.Complete)
					return
				}

				for _, e := range events {
					q.Output <- e
				}
			}
		}
	})
}

func (q *queue) GetComplete() chan struct{} {
	return q.Complete
}

func (q *queue) GetError() chan error {
	return q.Error
}

func (q *queue) GetInput() chan string {
	return q.Input
}

func (q *queue) GetOutput() chan Event {
	return q.Output
}

func (q *queue) NewExtendHeadEvent(cd ConnectionDetail, cp string) (Event, error) {
	var event Event
	var err error

	// In case the last peer of the current connection matches the first peer of
	// the given connection path, we emit a ExtendHead event.
	if strings.HasPrefix(cp, cd.Destination) {
		config := DefaultExtendHeadConfig()
		config.Connection = cd.Connection
		config.ConnectionPath = cp
		config.Destination = cd.Destination
		config.Source = cd.Source
		event, err = NewExtendHead(config)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	return event, nil
}

func (q *queue) NewExtendTailEvent(cd ConnectionDetail, cp string) (Event, error) {
	var event Event
	var err error

	// In case the first peer of the current connection matches the last peer of
	// the given connection path, we emit a ExtendTail event.
	if strings.HasSuffix(cp, cd.Source) {
		config := DefaultExtendTailConfig()
		config.Connection = cd.Connection
		config.ConnectionPath = cp
		config.Destination = cd.Destination
		config.Source = cd.Source
		event, err = NewExtendTail(config)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	return event, nil
}

func (q *queue) NewMatchBodyEvent(cd ConnectionDetail, cp string) (Event, error) {
	var event Event
	var err error

	// In case the current connection matches the body of the given connection
	// path, we emit a MatchBody event.
	if strings.Contains(cp, cd.Connection) {
		config := DefaultMatchBodyConfig()
		config.Connection = cd.Connection
		config.ConnectionPath = cp
		config.Destination = cd.Destination
		config.Source = cd.Source
		event, err = NewMatchBody(config)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	return event, nil
}

func (q *queue) NewMatchHeadEvent(cd ConnectionDetail, cp string) (Event, error) {
	var event Event
	var err error

	// In case the current connection matches the head of the given connection
	// path, we emit a MatchHead event.
	if strings.HasPrefix(cp, cd.Connection) {
		config := DefaultMatchHeadConfig()
		config.Connection = cd.Connection
		config.ConnectionPath = cp
		config.Destination = cd.Destination
		config.Source = cd.Source
		event, err = NewMatchHead(config)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	return event, nil
}

// TODO comment
func (q *queue) NewMatchPathEvent(cd ConnectionDetail, cp string) (Event, error) {
	var event Event
	var err error

	// In case the current connection equals exactly the given connection path,
	// we emit a MatchPath event.
	if cd.Connection == cp {
		// TODO comment
		atomic.AddInt32(&q.ProcessedEvents, 1)

		config := DefaultMatchPathConfig()
		config.Connection = cd.Connection
		config.ConnectionPath = cp
		config.Destination = cd.Destination
		config.Source = cd.Source
		event, err = NewMatchPath(config)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	return event, nil
}

// TODO
func (q *queue) NewMatchTailEvent(cd ConnectionDetail, cp string) (Event, error) {
	var event Event
	var err error

	// In case the current connection matches the tail of the given connection
	// path, we emit a MatchTail event.
	if strings.HasSuffix(cp, cd.Connection) {
		config := DefaultMatchTailConfig()
		config.Connection = cd.Connection
		config.ConnectionPath = cp
		config.Destination = cd.Destination
		config.Source = cd.Source
		event, err = NewMatchTail(config)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	return event, nil
}

// TODO
func (q *queue) NewNewPathEvent(cd ConnectionDetail, cp string) (Event, error) {
	var event Event
	var err error

	// TODO comment
	if len(events) == 0 {
		// When we get to here no other rule matched. That means
		atomic.AddInt32(&q.ProcessedEvents, 1)
	}

	return event, nil
}

// TODO
func (q *queue) NewSplitPathEvent(cd ConnectionDetail, cp string) (Event, error) {
	var event Event
	var err error

	// TODO detect split

	return event, nil
}

// TODO comment
// cp is connection path
func (q *queue) Process(cp string) ([]Event, error) {
	var events []Event

	lookups := []func(cd ConnectionDetail, cp string) (Event, error){
		q.NewExtendHeadEvent,
		q.NewExtendTailEvent,
		q.NewMatchBodyEvent,
		q.NewMatchHeadEvent,
		q.NewMatchPathEvent,
		q.NewMatchTailEvent,
		q.NewNewPathEvent,
		q.NewSplitPathEvent,
	}

	for _, cd := range q.ConnectionDetails {
		for _, l := range lookups {
			e, err := l(cd, cp)
			if err != nil {
				return nil, maskAny(err)
			}
			if e != nil {
				events = append(events, e)
			}
		}
	}

	// TODO comment
	if int(atomic.LoadInt32(&q.ProcessedEvents)) == len(q.ConnectionDetails) {
		close(q.Complete)
	}

	// TODO filter relevant events
	// e.g. NewPath must only be emited in case there is no other event

	return events, nil
}

func (q *queue) Shutdown() {
	q.ShutdownOnce.Do(func() {
		close(q.Closer)
	})
}
