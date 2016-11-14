// Package main implements a command line client for Anna. Cobra CLI is used as
// framework. The commands are simple wrappers around the client package.
package main

import (
	"sync"

	"github.com/spf13/cobra"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
)

var (
	// version is the project version. It is given via buildflags that inject the
	// commit hash.
	version string
)

// New creates a new annactl service.
func New() systemspec.Annactl {
	return &annactl{}
}

type annactl struct {
	// Dependencies.

	serviceCollection servicespec.Collection
	textInterface     systemspec.TextInterfaceClient

	// Settings.

	bootOnce     sync.Once
	closer       chan struct{}
	flags        Flags
	metadata     map[string]string
	sessionID    string
	shutdownOnce sync.Once
	version      string
}

func (a *annactl) Boot() {
	a.bootOnce.Do(func() {
		id, err := a.Service().ID().New()
		if err != nil {
			panic(err)
		}
		a.metadata = map[string]string{
			"id":   id,
			"name": "annactl",
			"type": "service",
		}

		a.bootOnce = sync.Once{}
		a.closer = make(chan struct{}, 1)
		a.flags = Flags{}

		sessionID, err := a.Service().ID().New()
		if err != nil {
			panic(err)
		}
		a.sessionID = sessionID

		a.shutdownOnce = sync.Once{}
		a.version = version

		go a.listenToSignal()
	})
}

func (a *annactl) ExecAnnactlCmd(cmd *cobra.Command, args []string) {
	a.Service().Log().Line("func", "ExecAnnactlCmd")

	cmd.HelpFunc()(cmd, nil)
}

func (a *annactl) InitAnnactlCmd() *cobra.Command {
	//a.Service().Log().Line("func", "InitAnnactlCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "annactl",
		Short: "Interact with Anna's network API. For more information see https://github.com/xh3b4sd/anna.",
		Long:  "Interact with Anna's network API. For more information see https://github.com/xh3b4sd/anna.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Service collection.
			a.serviceCollection = a.newServiceCollection()

			// Text interface.
			a.textInterface = newTextInterface(a.serviceCollection, a.flags.GRPCAddr)
		},
		Run: a.ExecAnnactlCmd,
	}

	// Add sub commands.
	newCmd.AddCommand(a.InitAnnactlInterfaceCmd())
	newCmd.AddCommand(a.InitAnnactlVersionCmd())

	// Define command line flags.
	newCmd.PersistentFlags().StringVar(&a.flags.GRPCAddr, "grpc-addr", "127.0.0.1:9119", "host:port to bind Anna's gRPC server to")
	newCmd.PersistentFlags().StringVar(&a.flags.HTTPAddr, "http-addr", "127.0.0.1:9120", "host:port to bind Anna's HTTP server to")

	return newCmd
}

func (a *annactl) Metadata() map[string]string {
	return a.metadata
}

func (a *annactl) Service() servicespec.Collection {
	return a.serviceCollection
}

func (a *annactl) SetServiceCollection(sc servicespec.Collection) {
	a.serviceCollection = sc
}

func (a *annactl) SetTextInterface(ti systemspec.TextInterfaceClient) {
	a.textInterface = ti
}

func (a *annactl) Shutdown() {
	a.Service().Log().Line("func", "Shutdown")

	a.shutdownOnce.Do(func() {
		close(a.closer)
	})
}

func main() {
	newAnnactl := New()
	newAnnactl.Boot()
}
