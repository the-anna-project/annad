package command

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/annactl/command/endpoint"
	"github.com/xh3b4sd/anna/annactl/command/version"
)

// New creates a new annactl command.
func New() *Command {
	command := &Command{}

	command.SetEndpointCommand(endpoint.New())
	command.SetVersionCommand(version.New())

	return command
}

// Command represents the annactl command.
type Command struct {
	// Dependencies.

	endpointCommand *endpoint.Command
	versionCommand  *version.Command
}

// Execute represents the cobra run method.
func (c *Command) Execute(cmd *cobra.Command, args []string) {
	cmd.HelpFunc()(cmd, nil)
}

// New creates a new cobra command for the annactl command.
func (c *Command) New() *cobra.Command {
	newCommand := &cobra.Command{
		Use:   "annactl",
		Short: "Manage the API of the anna project. For more information see https://github.com/the-anna-project/annactl.",
		Long:  "Manage the API of the anna project. For more information see https://github.com/the-anna-project/annactl.",
		Run:   c.Execute,
	}

	newCommand.AddCommand(c.endpointCommand.New())
	newCommand.AddCommand(c.versionCommand.New())

	return newCommand
}

// EndpointCommand returns the endpoint subcommand of the annactl command.
func (c *Command) EndpointCommand() *endpoint.Command {
	return c.endpointCommand
}

// SetEndpointCommand sets the endpoint subcommand for the annactl command.
func (c *Command) SetEndpointCommand(endpointCommand *endpoint.Command) {
	c.endpointCommand = endpointCommand
}

// SetVersionCommand sets the version subcommand for the annactl command.
func (c *Command) SetVersionCommand(versionCommand *version.Command) {
	c.versionCommand = versionCommand
}

// VersionCommand returns the version subcommand of the annactl command.
func (c *Command) VersionCommand() *version.Command {
	return c.versionCommand
}
