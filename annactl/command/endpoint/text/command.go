package text

import (
	"bufio"
	"os"
	"os/signal"
	"sync"

	"github.com/spf13/cobra"

	textinputobject "github.com/the-anna-project/input/object/text"
	servicespec "github.com/the-anna-project/spec/service"
	"github.com/xh3b4sd/anna/annactl/config"
)

// New creates a new text command.
func New() *Command {
	command := &Command{}

	command.SetConfigCollection(config.NewCollection())

	return command
}

// Command represents the text command.
type Command struct {
	// Dependencies.

	configCollection  *config.Collection
	serviceCollection servicespec.ServiceCollection

	// Settings.

	bootOnce     sync.Once
	shutdownOnce sync.Once
}

// Boot opens up a streams to the neural network.
func (c *Command) Boot() {
	go c.ListenToSignal()

	c.serviceCollection = c.newServiceCollection()
	go c.serviceCollection.Boot()

	go func() {
		c.serviceCollection.Log().Line("msg", "Waiting for input.\n")

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			textInputObject := textinputobject.New()
			textInputObject.SetEcho(c.configCollection.Endpoint().Text().Echo())
			textInputObject.SetInput(scanner.Text())
			textInputObject.SetSessionID(c.configCollection.Session().ID())

			c.serviceCollection.Input().Text().Channel() <- textInputObject

			err := scanner.Err()
			if err != nil {
				c.serviceCollection.Log().Line("msg", "%#v", maskAny(err))
			}
		}
	}()

	for {
		select {
		case textResponse := <-c.serviceCollection.Output().Text().Channel():
			c.serviceCollection.Log().Line("response", textResponse.Output())
		}
	}
}

// Execute represents the cobra run method.
func (c *Command) Execute(cmd *cobra.Command, args []string) {
	c.Boot()
}

// ForceShutdown forces the process to stop immediately.
func (c *Command) ForceShutdown() {
	os.Exit(0)
}

// New creates a new cobra command for the text command.
func (c *Command) New() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "text",
		Short: "Feed the neural network with text.",
		Long:  "Feed the neural network with text.",
		Run:   c.Execute,
	}

	c.configCollection.Config().SetDir(newCmd.PersistentFlags().String("config.dir", ".", "directory where to find the config file"))
	c.configCollection.Config().SetName(newCmd.PersistentFlags().String("config.name", "config", "name of the config file without extension"))

	c.configCollection.Endpoint().Text().SetAddress(newCmd.PersistentFlags().String("endpoint.text.address", "127.0.0.1:9119", "host:port to bind the text endpoint to"))
	c.configCollection.Endpoint().Text().SetEcho(newCmd.PersistentFlags().Bool("endpoint.text.echo", false, "echo input and bypass the neural network"))

	c.configCollection.Session().SetID(newCmd.PersistentFlags().String("session.id", "", "ID of the session currently being active"))

	return newCmd
}

// ListenToSignal listens to OS signals to be catched and processed if desired.
func (c *Command) ListenToSignal() {
	listener := make(chan os.Signal, 2)
	signal.Notify(listener, os.Interrupt, os.Kill)

	<-listener

	go c.Shutdown()

	<-listener

	c.ForceShutdown()
}

// SetConfigCollection sets the config collection for the text command to
// configure the endpoint client.
func (c *Command) SetConfigCollection(configCollection *config.Collection) {
	c.configCollection = configCollection
}

// Shutdown initializes the shutdown of the neural network.
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
