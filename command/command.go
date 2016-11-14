package command

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/command/boot"
	"github.com/xh3b4sd/anna/command/version"
)

// New creates a new annad command.
func New() *Command {
	return &Command{}
}

// Command represents the annad command.
type Command struct {
	// Dependencies.

	bootCommand    *boot.Command
	versionCommand *version.Command
}

// Execute represents the cobra run method.
func (c *Command) Execute(cmd *cobra.Command, args []string) {
	cmd.HelpFunc()(cmd, nil)
}

// New creates a new cobra command for the annad command.
func (c *Command) New() *cobra.Command {
	newCommand := &cobra.Command{
		Use:   "annad",
		Short: "Manage the daemon of the anna project. For more information see https://github.com/xh3b4sd/anna.",
		Long:  "Manage the daemon of the anna project. For more information see https://github.com/xh3b4sd/anna.",
		Run:   c.Execute,
	}

	newCommand.AddCommand(c.bootCommand.New())
	newCommand.AddCommand(c.versionCommand.New())

	return newCommand
}

// SetBootCommand sets the boot subcommand for the annad command.
func (c *Command) SetBootCommand(command *boot.Command) {
	c.bootCommand = command
}

// SetVersionCommand sets the version subcommand for the annad command.
func (c *Command) SetVersionCommand(command *version.Command) {
	c.versionCommand = command
}
