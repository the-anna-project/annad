package spec

import (
	"github.com/spf13/cobra"
)

// Anna represents the main object, the programm root. It contains all further
// objects and manages the whole control flow around Anna's implementations.
type Anna interface {
	// Boot initializes and starts Anna like booting a machine. The call to Boot
	// blocks until Anna is completely initialized, so you might want to call it
	// in a separate goroutine.
	Boot()

	// ExecAnnaCmd executes the root command.
	ExecAnnaCmd(cmd *cobra.Command, args []string)

	// ExecAnnaVersionCmd executes the version command.
	ExecAnnaVersionCmd(cmd *cobra.Command, args []string)

	// ForceShutdown ends all processes of Anna immediately.
	ForceShutdown()

	// InitAnnaCmd initializes the root command.
	InitAnnaCmd() *cobra.Command

	// InitAnnaVersionCmd initializes the version command.
	InitAnnaVersionCmd() *cobra.Command

	Object

	// Shutdown ends all processes of Anna like shutting down a machine. The call
	// to Shutdown blocks until Anna is completely shut down, so you might want
	// to call it in a separate goroutine.
	Shutdown()

	ServiceProvider

	StorageProvider
}
