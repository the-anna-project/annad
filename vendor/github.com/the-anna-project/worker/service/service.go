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

func (s *service) Execute(config objectspec.WorkerExecuteConfig) error {
	var wg sync.WaitGroup
	var once sync.Once

	canceler := make(chan struct{}, 1)
	errors := make(chan error, 1)

	if config.Canceler() != nil {
		go func() {
			<-config.Canceler()
			// Receiving a signal from the global canceler will forward the
			// cancelation to all workers. Simply closing the workers canceler wil
			// broadcast the signal to each listener. Here we also make sure we do
			// not close on a closed channel by only closing once.
			once.Do(func() {
				close(canceler)
			})
		}()
	}

	for n := 0; n < config.NumWorkers(); n++ {
		go func() {
			for _, action := range config.Actions() {
				wg.Add(1)
				go func() {
					defer wg.Done()

					err := action(canceler)
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
		}()
	}

	wg.Wait()

	select {
	case err := <-errors:
		return err
	default:
		return nil
	}
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
