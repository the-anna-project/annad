package endpoint

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/annactl/command/endpoint/text"
)

// New creates a new endpoint command.
func New() *Command {
	command := &Command{}

	command.SetTextCommand(text.New())

	return command
}

// Command represents the endpoint command.
type Command struct {
	// Settings.

	textCommand *text.Command
}

// Execute represents the cobra run method.
func (c *Command) Execute(cmd *cobra.Command, args []string) {
	cmd.HelpFunc()(cmd, nil)
}

// New creates a new cobra command for the endpoint command.
func (c *Command) New() *cobra.Command {
	newCommand := &cobra.Command{
		Use:   "endpoint",
		Short: "Manage the network endpoints of the API of the anna project.",
		Long:  "Manage the network endpoints of the API of the anna project.",
		Run:   c.Execute,
	}

	newCommand.AddCommand(c.textCommand.New())

	return newCommand
}

// TextCommand returns the text subcommand of the endpoint command.
func (c *Command) TextCommand() *text.Command {
	return c.textCommand
}

// SetTextCommand sets the text subcommand for the endpoint command.
func (c *Command) SetTextCommand(textCommand *text.Command) {
	c.textCommand = textCommand
}
