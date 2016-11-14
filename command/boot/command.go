package boot

import (
	"os"
	"os/signal"
	"sync"

	"github.com/spf13/cobra"
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// New creates a new boot command.
func New() *Command {
	return &Command{}
}

type Command struct {
	// Dependencies.

	configCollection  *config.Config
	serviceCollection servicespec.Collection

	// Settings.

	bootOnce       sync.Once
	flags          Flags
	projectVersion string
	shutdownOnce   sync.Once
	version        string
}

func (c *Command) Boot() {
	c.serviceCollection = c.newServiceCollection()
	go c.serviceCollection.Boot()

	go c.ListenToSignal()

	// Block the main goroutine forever. The process is only supposed to be ended
	// by a call to Shutdown or ForceShutdown.
	select {}
}

func (c *Command) Execute(cmd *cobra.Command, args []string) {
	s.Boot()
}

func (c *Command) ForceShutdown() {
	os.Exit(0)
}

func (c *Command) New() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "boot",
		Short: "Boot and run the anna daemon.",
		Long:  "Boot and run the anna daemon.",
		Run:   c.Execute,
	}

	c.configCollection.Endpoint().Text().SetAddress(newCmd.PersistentFlags().String("endpoint.text.address", "127.0.0.1:9119", "host:port to bind Anna's gRPC server to"))
	c.configCollection.Endpoint().Metric().SetAddress(newCmd.PersistentFlags().String("endpoint.metric.address", "127.0.0.1:9120", "host:port to bind Anna's HTTP server to"))

	c.configCollection.Space().Connection().SetWeight(newCmd.PersistentFlags().Int("space.connection.weight", 0, "default weight of new connections within the connection space"))

	c.configCollection.Space().Dimension().SetCount(newCmd.PersistentFlags().Int("space.dimension.count", 3, "default number of directional coordinates within the connection space"))
	c.configCollection.Space().Dimension().SetDepth(newCmd.PersistentFlags().Int("space.dimension.depth", 1000000, "default size of each directional coordinate within the connection space"))

	c.configCollection.Space().Peer().SetPosition(newCmd.PersistentFlags().String("space.peer.position", "0,0,0", "default position of new peers within the connection space"))

	c.configCollection.Storage().Connection().SetAddress(newCmd.PersistentFlags().String("storage.connection.address", "127.0.0.1:6379", "host:port to connect to connection storage"))
	c.configCollection.Storage().Connection().SetKind(newCmd.PersistentFlags().String("storage.connection.kind", "memory", "storage kind to use for persistency (e.g. redis)"))
	c.configCollection.Storage().Connection().SetPrefix(newCmd.PersistentFlags().String("storage.connection.prefix", "anna", "prefix used to prepend to connection storage keys"))

	c.configCollection.Storage().Feature().SetAddress(newCmd.PersistentFlags().String("storage.feature.address", "127.0.0.1:6380", "host:port to connect to feature storage"))
	c.configCollection.Storage().Feature().SetKind(newCmd.PersistentFlags().String("storage.feature.kind", "memory", "storage kind to use for persistency (e.g. redis)"))
	c.configCollection.Storage().Feature().SetPrefix(newCmd.PersistentFlags().String("storage.feature.prefix", "anna", "prefix used to prepend to feature storage keys"))

	c.configCollection.Storage().General().SetAddress(newCmd.PersistentFlags().String("storage.general.address", "127.0.0.1:6381", "host:port to connect to general storage"))
	c.configCollection.Storage().General().SetKind(newCmd.PersistentFlags().String("storage.general.kind", "memory", "storage kind to use for persistency (e.g. redis)"))
	c.configCollection.Storage().General().SetPrefix(newCmd.PersistentFlags().String("storage.general.prefix", "anna", "prefix used to prepend to general storage keys"))

	return newCmd
}

func (c *Command) ListenToSignal() {
	listener := make(chan os.Signal, 2)
	signal.Notify(listener, os.Interrupt, os.Kill)

	<-listener

	go c.Shutdown()

	<-listener

	c.ForceShutdown()
}

func (c *Command) SetConfigCollection(cc *config.Config) {
	c.configCollection = cc
}

func (c *Command) SetProjectVersion(projectVersion string) {
	c.projectVersion = projectVersion
}

func (c *Command) Shutdown() {
	c.shutdownOnce.Do(func() {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			c.serviceCollection.Shutdown()
			wg.Done()
		}()

		wg.Wait()

		os.Exit(0)
	})
}
