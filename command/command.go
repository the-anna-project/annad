package command

import "github.com/spf13/cobra"

func New() *Command {
	return &Command{}
}

type Command struct {
	// Dependencies.

	bootCommand    *cobra.Command
	versionCommand *cobra.Command
}

func (c *Command) Execute(cmd *cobra.Command, args []string) {
	cmd.HelpFunc()(cmd, nil)
}

func (c *Command) New() *cobra.Command {
	newCommand := &cobra.Command{
		Use:   "annad",
		Short: "Manage the daemon of the anna project. For more information see https://github.com/xh3b4sd/anna.",
		Long:  "Manage the daemon of the anna project. For more information see https://github.com/xh3b4sd/anna.",
		Run:   c.Execute,
	}

	newCommand.AddCommand(a.bootCommand.New())
	newCommand.AddCommand(a.versionCommand.New())

	return newCommand
}

func (c *Command) SetBootCommand(command *cobra.Command) {
	c.bootCommand = command
}

func (c *Command) SetVersionCommand(command *cobra.Command) {
	c.versionCommand = command
}
