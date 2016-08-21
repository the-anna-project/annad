package spec

import (
	"github.com/spf13/cobra"
)

// Annactl represents the business logic implementation usable in command line
// tools. There it is responsible for bridging between the command line
// framework (cobra) and the actual client library.
type Annactl interface {
	// Boot initializes and executes the command line tool.
	Boot()

	// ExecControlCmd executes the control command.
	ExecControlCmd(cmd *cobra.Command, args []string)

	// ExecControlLogCmd executes the control log
	ExecControlLogCmd(cmd *cobra.Command, args []string)

	// ExecControlLogResetCmd executes the control log
	ExecControlLogResetCmd(cmd *cobra.Command, args []string)

	// ExecControlLogResetLevelsCmd executes the control log
	ExecControlLogResetLevelsCmd(cmd *cobra.Command, args []string)

	// ExecControlLogResetObjectsCmd executes the control log
	ExecControlLogResetObjectsCmd(cmd *cobra.Command, args []string)

	// ExecControlLogResetVerbosityCmd executes the control log
	ExecControlLogResetVerbosityCmd(cmd *cobra.Command, args []string)

	// ExecControlLogSetCmd executes the control log
	ExecControlLogSetCmd(cmd *cobra.Command, args []string)

	// ExecControlLogSetLevelsCmd executes the control log
	ExecControlLogSetLevelsCmd(cmd *cobra.Command, args []string)

	// ExecControlLogSetObjectsCmd executes the control log
	ExecControlLogSetObjectsCmd(cmd *cobra.Command, args []string)

	// ExecControlLogSetVerbosityCmd executes the control log
	ExecControlLogSetVerbosityCmd(cmd *cobra.Command, args []string)

	// ExecInterfaceCmd executes the interface command.
	ExecInterfaceCmd(cmd *cobra.Command, args []string)

	// ExecInterfaceTextCmd executes the interface text command.
	ExecInterfaceTextCmd(cmd *cobra.Command, args []string)

	// ExecInterfaceTextReadCmd executes the interface text read command.
	ExecInterfaceTextReadCmd(cmd *cobra.Command, args []string)

	// ExecInterfaceTextReadFileCmd executes the interface text read file
	// command.
	ExecInterfaceTextReadFileCmd(cmd *cobra.Command, args []string)

	// ExecInterfaceTextReadPlainCmd executes the interface text read plain
	// command.
	ExecInterfaceTextReadPlainCmd(cmd *cobra.Command, args []string)

	// ExecVersionCmd executes the version command.
	ExecVersionCmd(cmd *cobra.Command, args []string)

	// GetSessionID returns the current session ID. It looks up the configured
	// session file path to read the session from there. If there is none, a new
	// session is be created and the session file is written using it.
	GetSessionID() (string, error)

	// InitControlCmd initializes the control command.
	InitControlCmd() *cobra.Command

	// InitControlLogCmd initializes the control log command.
	InitControlLogCmd() *cobra.Command

	// InitControlLogResetCmd initializes the control log reset command.
	InitControlLogResetCmd() *cobra.Command

	// InitControlLogResetLevelsCmd initializes the control log reset levels
	// command.
	InitControlLogResetLevelsCmd() *cobra.Command

	// InitControlLogResetObjectsCmd initializes the control log reset objects
	// command.
	InitControlLogResetObjectsCmd() *cobra.Command

	// InitControlLogResetVerbosityCmd initializes the control log reset
	// verbosity command.
	InitControlLogResetVerbosityCmd() *cobra.Command

	// InitControlLogSetCmd initializes the control log set command.
	InitControlLogSetCmd() *cobra.Command

	// InitControlLogSetLevelsCmd initializes the control log set levels command.
	InitControlLogSetLevelsCmd() *cobra.Command

	// InitControlLogSetObjectsCmd initializes the control log set objects
	// command.
	InitControlLogSetObjectsCmd() *cobra.Command

	// InitControlLogSetVerbosityCmd initializes the control log set verbosity
	// command.
	InitControlLogSetVerbosityCmd() *cobra.Command

	// InitInterfaceCmd initializes the interface command.
	InitInterfaceCmd() *cobra.Command

	// InitInterfaceTextCmd initializes the interface text command.
	InitInterfaceTextCmd() *cobra.Command

	// InitInterfaceTextReadCmd initializes the interface text read command.
	InitInterfaceTextReadCmd() *cobra.Command

	// InitInterfaceTextReadFileCmd initializes the interface text read file
	// command.
	InitInterfaceTextReadFileCmd() *cobra.Command

	// InitInterfaceTextReadPlainCmd initializes the interface text read plain
	// command.
	InitInterfaceTextReadPlainCmd() *cobra.Command

	// InitVersionCmd initializes the version command.
	InitVersionCmd() *cobra.Command

	// Shutdown ends all processes of the command line tool like shutting down a
	// machine. The call to Shutdown blocks until the command line tool is
	// completely shut down, so you might want to call it in a separate
	// goroutine.
	Shutdown()
}
