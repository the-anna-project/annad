package spec

import (
	"github.com/spf13/cobra"

	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// Annad represents the anna daemon object. It contains all further objects and
// manages the whole control flow around Anna's implementations.
type Annad interface {
	// Boot initializes and starts Anna like booting a machine. The call to Boot
	// blocks until Anna is completely initialized, so you might want to call it
	// in a separate goroutine.
	Boot()

	// ExecAnnadCmd executes the root command.
	ExecAnnadCmd(cmd *cobra.Command, args []string)

	// ExecAnnadVersionCmd executes the version command.
	ExecAnnadVersionCmd(cmd *cobra.Command, args []string)

	// ForceShutdown ends all processes of Anna immediately.
	ForceShutdown()

	// InitAnnadCmd initializes the root command.
	InitAnnadCmd() *cobra.Command

	// InitAnnadVersionCmd initializes the version command.
	InitAnnadVersionCmd() *cobra.Command

	Service() servicespec.Collection

	SetServiceCollection(sc servicespec.Collection)

	// Shutdown ends all processes of Anna like shutting down a machine. The call
	// to Shutdown blocks until Anna is completely shut down, so you might want
	// to call it in a separate goroutine.
	Shutdown()
}
