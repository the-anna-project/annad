package text

import (
	"os"

	kitlog "github.com/go-kit/kit/log"

	endpointcollection "github.com/the-anna-project/client/collection"
	textendpoint "github.com/the-anna-project/client/service/text"
	"github.com/the-anna-project/collection"
	"github.com/the-anna-project/fs/memory"
	"github.com/the-anna-project/id"
	inputcollection "github.com/the-anna-project/input/collection"
	textinputservice "github.com/the-anna-project/input/service/text"
	"github.com/the-anna-project/log"
	outputcollection "github.com/the-anna-project/output/collection"
	textoutputservice "github.com/the-anna-project/output/service/text"
	"github.com/the-anna-project/permutation/service"
	servicespec "github.com/the-anna-project/spec/service"
)

func (c *Command) newServiceCollection() servicespec.ServiceCollection {
	collection := collection.New()

	collection.SetEndpointCollection(c.newEndpointCollection())
	collection.SetFSService(c.newFSService())
	collection.SetIDService(c.newIDService())
	collection.SetInputCollection(c.newInputCollection())
	collection.SetLogService(c.newLogService())
	collection.SetOutputCollection(c.newOutputCollection())
	collection.SetPermutationService(c.newPermutationService())

	collection.Endpoint().Text().SetServiceCollection(collection)
	collection.FS().SetServiceCollection(collection)
	collection.ID().SetServiceCollection(collection)
	collection.Input().Text().SetServiceCollection(collection)
	collection.Log().SetServiceCollection(collection)
	collection.Output().Text().SetServiceCollection(collection)
	collection.Permutation().SetServiceCollection(collection)

	return collection
}

func (c *Command) newEndpointCollection() servicespec.EndpointCollection {
	newCollection := endpointcollection.New()

	textService := textendpoint.New()
	textService.SetAddress(c.configCollection.Endpoint().Text().Address())

	newCollection.SetTextService(textService)

	return newCollection
}

// TODO make mem/os configurable
func (c *Command) newFSService() servicespec.FSService {
	return memory.New()
}

func (c *Command) newIDService() servicespec.IDService {
	return id.New()
}

func (c *Command) newInputCollection() servicespec.InputCollection {
	newCollection := inputcollection.New()

	newCollection.SetTextService(textinputservice.New())

	return newCollection
}

func (c *Command) newLogService() servicespec.LogService {
	newService := log.New()

	newService.SetRootLogger(kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr)))

	return newService
}

func (c *Command) newOutputCollection() servicespec.OutputCollection {
	newCollection := outputcollection.New()

	newCollection.SetTextService(textoutputservice.New())

	return newCollection
}

func (c *Command) newPermutationService() servicespec.PermutationService {
	return permutation.New()
}
