// Package worker implements a service to process work concurrently.
package service

import (
	"sync"

	objectspec "github.com/the-anna-project/spec/object"
	servicespec "github.com/the-anna-project/spec/service"
)

// New creates a new worker service.
func New() servicespec.WorkerService {
	return &service{
		// Dependencies.
		serviceCollection: nil,

		// Settings.
		metadata: map[string]string{},
	}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	metadata map[string]string
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "worker",
		"type": "service",
	}
}

// ExecuteConfig provides a default configuration to execute a new worker pool
// by best effort.
func (s *service) ExecuteConfig() objectspec.WorkerExecuteConfig {
	return newExecuteConfig()
}

func (s *service) Execute(config objectspec.WorkerExecuteConfig) chan error {
	var wg sync.WaitGroup
	var once sync.Once

	canceler := make(chan struct{}, 1)
	errors := make(chan error, config.NumWorkers())

	if config.Canceler() != nil {
		go func() {
			select {
			case <-config.Canceler():
				// Receiving a signal from the global canceler will forward the
				// cancelation to all workers. Simply closing the workers canceler wil
				// broadcast the signal to each listener. Here we also make sure we do
				// not close on a closed channel by only closing once.
				once.Do(func() {
					close(canceler)
				})
			}
		}()
	}

	for n := 0; n < config.NumWorkers(); n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			err := config.Action()(canceler)
			if err != nil {
				if config.CancelOnError() && config.Canceler() != nil {
					// Closing the canceler channel acts as broadcast to all workers that
					// should listen to the canceler. Here we also make sure we do not
					// close on a closed channel by only closing once.
					once.Do(func() {
						close(config.Canceler())
					})
				}
				errors <- err
			}
		}()
	}

	wg.Wait()

	// We can savely close the error and canceler channels here because nobody
	// can write into it anymore. Thus we can clean the environment to not leave
	// uncollectable garbage. It is still save to read from the closed error
	// channel.
	close(errors)

	return errors
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(serviceCollection servicespec.ServiceCollection) {
	s.serviceCollection = serviceCollection
}
