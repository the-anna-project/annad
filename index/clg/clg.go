// Package clg implementes fundamental actions used to create strategies that
// allow to discover new behavior for problem solving.
package clg

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeCLGIndex represents the object type of the CLG index object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeCLGIndex spec.ObjectType = "clg-index"
)

// Config represents the configuration used to create a new CLG index object.
type Config struct {
	// Dependencies.
	Collection spec.CLGCollection
	Log        spec.Log

	// Settings.
	NumCLGProfileWorkers int
}

// DefaultConfig provides a default configuration to create a new CLG index
// object by best effort.
func DefaultConfig() Config {
	newCLGCollection, err := NewCLGCollection(DefaultCLGCollectionConfig())
	if err != nil {
		panic(err)
	}

	newConfig := Config{
		// Dependencies.
		Collection: newCLGCollection,
		Log:        log.NewLog(log.DefaultConfig()),

		// Settings.
		NumCLGProfileWorkers: 10,
	}

	return newConfig
}

// NewCLGIndex creates a new configured CLG index object.
func NewCLGIndex(config Config) (spec.CLGIndex, error) {
	newCLGIndex := &clgIndex{
		Config:       config,
		BootOnce:     sync.Once{},
		ID:           id.NewObjectID(id.Hex128),
		Mutex:        sync.Mutex{},
		Type:         ObjectTypeCLGIndex,
		ShutdownOnce: sync.Once{},
	}

	if newCLGIndex.NumCLGProfileWorkers < 1 {
		return nil, maskAnyf(invalidConfigError, "number of CLG profile workers must be greater than 0")
	}

	newCLGIndex.Log.Register(newCLGIndex.GetType())

	return newCLGIndex, nil
}

type clgIndex struct {
	Config

	BootOnce     sync.Once
	ID           spec.ObjectID
	Mutex        sync.Mutex
	Type         spec.ObjectType
	ShutdownOnce sync.Once
}

func (i *clgIndex) Boot() {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call Boot")

	i.BootOnce.Do(func() {
		go func() {
			err := i.CreateCLGProfiles(i.GetCLGCollection())
			if err != nil {
				i.Log.WithTags(spec.Tags{L: "E", O: i, T: nil, V: 4}, "%#v", maskAny(err))
			}
		}()
	})
}

// TODO
func (i *clgIndex) createCLGProfile(clgCollection spec.CLGCollection, clgName string, closer chan struct{}) (spec.CLGProfile, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call createCLGProfile")

	// Fill a queue.
	clgNames, err := clgCollection.GetNamesMethod()
	if err != nil {
		return maskAny(err)
	}
	queue := make(chan string, len(clgNames))
	for _, clgName := range clgNames {
		queue <- clgName
	}

	// range over channel
	//     find argument types for given clg name
	//     hash method body for given clg name
	//     find right side neighbours for given clg name
	//         if no profile for checked neighbour
	//             push neighbour name back to channel

	return nil, nil
}

// TODO
func (i *clgIndex) CreateCLGProfiles(clgCollection spec.CLGCollection) error {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call CreateCLGProfiles")

	// Fill a queue.
	clgNames, err := clgCollection.GetNamesMethod()
	if err != nil {
		return maskAny(err)
	}
	queue := make(chan string, len(clgNames))
	for _, clgName := range clgNames {
		queue <- clgName
	}
	// We can close the queue channel immediately, because it only provides a one
	// way ticket. As soon as a CLG name was fetched from the queue it is
	// considered WIP. A CLG name must never be requeued.
	close(queue)

	// Prepare synchronization channels.
	done := make(chan struct{}, 1)
	fail := make(chan error, 1)

	// Start N worker goroutines.
	go func() {
		var wg sync.Waitgroup
		workerPoolCloser := make(chan struct{}, 1)

		for n := 0; n < i.NumCLGProfileWorkers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				var err error
				workerCloser := make(chan struct{}, 1)

				for {
					select {
					case <-workerPoolCloser:
						workerCloser <- struct{}{}
						return
					case clgName := <-queue:
						clgProfile, createErr := i.createCLGProfile(clgCollection, clgName, workerCloser)
						if createErr != nil {
							err = createErr
						}

						if clgProfile.HasChanged() {
							storeErr := i.StoreCLGProfile(clgProfile)
							if storeErr != nil {
								err = storeErr
							}
						}

						if err != nil {
							workerCloser <- struct{}{}

							fail <- maskAny(err)
						}
					}
				}
			}()
		}

		wg.Wait()
		done <- struct{}{}
	}()

	var err error
	select {
	case failErr := <-fail:
		err = failErr
	case <-done:
	}

	close(done)
	close(fail)

	return maskAny(err)
}

func (i *clgIndex) GetCLGCollection() spec.CLGCollection {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call GetCLGCollection")

	return i.Collection
}

// TODO
func (i *clgIndex) isRightSideCLGNeighbour(clgCollection spec.CLGCollection, left, right CLGProfile) (bool, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call isRightSideNeighbour")

	// run clg chain
	// if error
	//     return false

	return false, nil
}

func (i *clgIndex) Shutdown() {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call Shutdown")

	i.ShutdownOnce.Do(func() {
	})
}
