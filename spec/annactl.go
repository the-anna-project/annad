package spec

import (
	"github.com/spf13/cobra"

	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// Annactl represents the business logic implementation usable in command line
// tools. There it is responsible for bridging between the command line
// framework (cobra) and the actual client library.
type Annactl interface {
	// Boot initializes and executes the command line tool.
	Boot()

	// ExecAnnactlCmd executes the root command.
	ExecAnnactlCmd(cmd *cobra.Command, args []string)

	// ExecAnnactlControlCmd executes the control command.
	ExecAnnactlControlCmd(cmd *cobra.Command, args []string)

	// ExecAnnactlControlLogCmd executes the control log
	ExecAnnactlControlLogCmd(cmd *cobra.Command, args []string)

	// ExecAnnactlControlLogResetCmd executes the control log
	ExecAnnactlControlLogResetCmd(cmd *cobra.Command, args []string)

	// ExecAnnactlControlLogResetLevelsCmd executes the control log
	ExecAnnactlControlLogResetLevelsCmd(cmd *cobra.Command, args []string)

	// ExecAnnactlControlLogResetObjectsCmd executes the control log
	ExecAnnactlControlLogResetObjectsCmd(cmd *cobra.Command, args []string)

	// ExecAnnactlControlLogResetVerbosityCmd executes the control log
	ExecAnnactlControlLogResetVerbosityCmd(cmd *cobra.Command, args []string)

	// ExecAnnactlControlLogSetCmd executes the control log
	ExecAnnactlControlLogSetCmd(cmd *cobra.Command, args []string)

	// ExecAnnactlControlLogSetLevelsCmd executes the control log
	ExecAnnactlControlLogSetLevelsCmd(cmd *cobra.Command, args []string)

	// ExecAnnactlControlLogSetObjectsCmd executes the control log
	ExecAnnactlControlLogSetObjectsCmd(cmd *cobra.Command, args []string)

	// ExecAnnactlControlLogSetVerbosityCmd executes the control log
	ExecAnnactlControlLogSetVerbosityCmd(cmd *cobra.Command, args []string)

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

	// InitAnnactlControlCmd initializes the control command.
	InitAnnactlControlCmd() *cobra.Command

	// InitAnnactlControlLogCmd initializes the control log command.
	InitAnnactlControlLogCmd() *cobra.Command

	// InitAnnactlControlLogResetCmd initializes the control log reset command.
	InitAnnactlControlLogResetCmd() *cobra.Command

	// InitAnnactlControlLogResetLevelsCmd initializes the control log reset
	// levels command.
	InitAnnactlControlLogResetLevelsCmd() *cobra.Command

	// InitAnnactlControlLogResetObjectsCmd initializes the control log reset
	// objects command.
	InitAnnactlControlLogResetObjectsCmd() *cobra.Command

	// InitAnnactlControlLogResetVerbosityCmd initializes the control log reset
	// verbosity command.
	InitAnnactlControlLogResetVerbosityCmd() *cobra.Command

	// InitAnnactlControlLogSetCmd initializes the control log set command.
	InitAnnactlControlLogSetCmd() *cobra.Command

	// InitAnnactlControlLogSetLevelsCmd initializes the control log set levels
	// command.
	InitAnnactlControlLogSetLevelsCmd() *cobra.Command

	// InitAnnactlControlLogSetObjectsCmd initializes the control log set objects
	// command.
	InitAnnactlControlLogSetObjectsCmd() *cobra.Command

	// InitAnnactlControlLogSetVerbosityCmd initializes the control log set
	// verbosity command.
	InitAnnactlControlLogSetVerbosityCmd() *cobra.Command

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

	servicespec.Provider

	// Shutdown ends all processes of the command line tool like shutting down a
	// machine. The call to Shutdown blocks until the command line tool is
	// completely shut down, so you might want to call it in a separate
	// goroutine.
	Shutdown()
}
