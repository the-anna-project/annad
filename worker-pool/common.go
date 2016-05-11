package workerpool

import (
	"sync"
)

// executeOnce executes the worker pool's workers as described by
// spec.WorkerPool.Execute. Note that executeOnce is only supposed to be
// executed once. This is why it is wrapped and synchronized by
// spec.WorkerPool.Execute.
func (wp *workerPool) executeOnce() chan error {
	var wg sync.WaitGroup
	var errors chan error

	canceler := make(chan struct{}, 1)

	go func() {
		select {
		case <-wp.Canceler:
			// Receiving a signal from the global canceler will forward the
			// cancelation to all workers. Simply closing the workers canceler wil
			// broadcast the signal to each listener.
			close(canceler)
		}
	}()

	for n := 0; n < wp.NumWorkers; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			err := wp.WorkerFunc(canceler)
			if err != nil {
				if wp.CancelOnError {
					// Closing the canceler channel acts as broadcast to all workers that
					// should listen to the canceler.
					close(canceler)
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
	close(wp.Canceler)

	return errors
}
