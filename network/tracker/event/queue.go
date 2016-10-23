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

		BootOnce:         sync.Once{},
		Closer:           make(chan struct{}, 1),
		Complete:     make(chan struct{}, 1),
		ConnectionDetails: nil,
		Error:        make(chan error, 1),
		Input:           make(chan string, 1),
		ProcessedEvents:  0,
		Output:          make(chan Event, 1),
		ShutdownOnce:     sync.Once{},
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
	Complete      chan struct{}
	ConnectionDetails []ConnectionDetail
	Error         chan error
	Input            chan string
	Output           chan Event
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

func (q *queue) Process(connectionPath string) ([]Event, error) {
	var events []Event

	for _, connectionDetail := range q.ConnectionDetails {
    // In case the current connection equals exactly the given connection path,
    // we emit a MatchPath event.
		if connectionDetail.Connection == connectionPath {
      // TODO comment
			atomic.AddInt32(&q.ProcessedEvents, 1)

      config := DefaultMatchPathConfig()
      config.Connection = connectionDetail.Connection
      config.ConnectionPath = connectionPath
      config.Destination = connectionDetail.Destination
      config.Source = connectionDetail.Source
      e, err := NewMatchPath(config)
      if err != nil {
        return nil, maskAny(err)
      }
			events = append(events, e)
		}

    // In case the current connection matches the head of the given connection
    // path, we emit a MatchHead event.
		if strings.HasPrefix(connectionPath, connectionDetail.Connection) {
      config := DefaultMatchHeadConfig()
      config.Connection = connectionDetail.Connection
      config.ConnectionHead = connectionHead
      config.Destination = connectionDetail.Destination
      config.Source = connectionDetail.Source
      e, err := NewMatchHead(config)
      if err != nil {
        return nil, maskAny(err)
      }
			events = append(events, e)
		}

    // In case the current connection matches the tail of the given connection
    // path, we emit a MatchTail event.
		if strings.HasSuffix(connectionPath, connectionDetail.Connection) {
      config := DefaultMatchTailConfig()
      config.Connection = connectionDetail.Connection
      config.ConnectionTail = connectionTail
      config.Destination = connectionDetail.Destination
      config.Source = connectionDetail.Source
      e, err := NewMatchTail(config)
      if err != nil {
        return nil, maskAny(err)
      }
			events = append(events, e)
		}

    // In case the current connection matches the body of the given connection
    // path, we emit a MatchBody event.
		if strings.Contains(connectionPath, connectionDetail.Connection) {
      config := DefaultMatchBodyConfig()
      config.Connection = connectionDetail.Connection
      config.ConnectionBody = connectionBody
      config.Destination = connectionDetail.Destination
      config.Source = connectionDetail.Source
      e, err := NewMatchBody(config)
      if err != nil {
        return nil, maskAny(err)
      }
			events = append(events, e)
		}

    // In case the last peer of the current connection matches the first peer of
    // the given connection path, we emit a ExtendHead event.
		if strings.HasPrefix(connectionPath, connectionDetail.Source) {
      config := DefaultExtendHeadConfig()
      config.Connection = connectionDetail.Connection
      config.ConnectionBody = connectionBody
      config.Destination = connectionDetail.Destination
      config.Source = connectionDetail.Source
      e, err := NewExtendHead(config)
      if err != nil {
        return nil, maskAny(err)
      }
			events = append(events, e)
		}

    // In case the first peer of the current connection matches the last peer of
    // the given connection path, we emit a ExtendTail event.
		if strings.HasSuffix(connectionPath, connectionDetail.Destination) {
      config := DefaultExtendTailConfig()
      config.Connection = connectionDetail.Connection
      config.ConnectionBody = connectionBody
      config.Destination = connectionDetail.Destination
      config.Source = connectionDetail.Source
      e, err := NewExtendTail(config)
      if err != nil {
        return nil, maskAny(err)
      }
			events = append(events, e)
		}

    // TODO detect split
		if strings.HasSuffix(connectionPath, connectionDetail.Destination) {
			t = EventSplitPath
		}

// TODO comment
		if len(events) = 0 {
			// When we get to here no other rule matched. That means
			atomic.AddInt32(&q.ProcessedEvents, 1)
			t = EventNewPath
		}

		if int(atomic.LoadInt32(&q.ProcessedEvents)) == len(q.ConnectionDetails) {
			close(q.Complete)
		}
	}

  return events, nil
}

func (q *queue) Shutdown() {
	q.ShutdownOnce.Do(func() {
		close(q.Closer)
	})
}
