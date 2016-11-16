package spec

import (
	"github.com/spf13/cobra"

	servicespec "github.com/the-anna-project/spec/service"
)

// Annactl represents the business logic implementation usable in command line
// tools. There it is responsible for bridging between the command line
// framework (cobra) and the actual client library.
type Annactl interface {
	// Boot initializes and executes the command line tool.
	Boot()

	// ExecAnnactlCmd executes the root command.
	ExecAnnactlCmd(cmd *cobra.Command, args []string)

	// ExecAnnactlInterfaceCmd executes the interface command.
	ExecAnnactlInterfaceCmd(cmd *cobra.Command, args []string)

	// ExecAnnactlInterfaceTextCmd executes the interface text command.
	ExecAnnactlInterfaceTextCmd(cmd *cobra.Command, args []string)

	// ExecAnnactlInterfaceTextReadCmd executes the interface text read command.
	ExecAnnactlInterfaceTextReadCmd(cmd *cobra.Command, args []string)

	// ExecAnnactlInterfaceTextReadFileCmd executes the interface text read file
	// command.
	ExecAnnactlInterfaceTextReadFileCmd(cmd *cobra.Command, args []string)

	// ExecAnnactlInterfaceTextReadPlainCmd executes the interface text read
	// plain command.
	ExecAnnactlInterfaceTextReadPlainCmd(cmd *cobra.Command, args []string)

	// ExecAnnactlVersionCmd executes the version command.
	ExecAnnactlVersionCmd(cmd *cobra.Command, args []string)

	// GetSessionID returns the current session ID. It looks up the configured
	// session file path to read the session from there. If there is none, a new
	// session is be created and the session file is written using it.
	GetSessionID() (string, error)

	// InitAnnactlCmd initializes the root command.
	InitAnnactlCmd() *cobra.Command

	// InitAnnactlInterfaceCmd initializes the interface command.
	InitAnnactlInterfaceCmd() *cobra.Command

	// InitAnnactlInterfaceTextCmd initializes the interface text command.
	InitAnnactlInterfaceTextCmd() *cobra.Command

	// InitAnnactlInterfaceTextReadCmd initializes the interface text read
	// command.
	InitAnnactlInterfaceTextReadCmd() *cobra.Command

	// InitAnnactlInterfaceTextReadFileCmd initializes the interface text read
	// file command.
	InitAnnactlInterfaceTextReadFileCmd() *cobra.Command

	// InitAnnactlInterfaceTextReadPlainCmd initializes the interface text read
	// plain command.
	InitAnnactlInterfaceTextReadPlainCmd() *cobra.Command

	// InitAnnactlVersionCmd initializes the version command.
	InitAnnactlVersionCmd() *cobra.Command

	Service() servicespec.ServiceCollection

	SetServiceCollection(serviceCollection servicespec.ServiceCollection)

	// Shutdown ends all processes of the command line tool like shutting down a
	// machine. The call to Shutdown blocks until the command line tool is
	// completely shut down, so you might want to call it in a separate
	// goroutine.
	Shutdown()
}
