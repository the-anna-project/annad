package main

import "github.com/spf13/cobra"

func (a *annactl) InitAnnactlInterfaceCmd() *cobra.Command {
	a.Service().Log().Line("func", "InitAnnactlInterfaceCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "interface",
		Short: "Interface for Anna's behaviours.",
		Long:  "Interface for Anna's behaviours.",
		Run:   a.ExecAnnactlInterfaceCmd,
	}

	// Add sub commands.
	newCmd.AddCommand(a.InitAnnactlInterfaceTextCmd())

	return newCmd
}

func (a *annactl) ExecAnnactlInterfaceCmd(cmd *cobra.Command, args []string) {
	a.Service().Log().Line("func", "ExecAnnactlInterfaceCmd")

	cmd.HelpFunc()(cmd, nil)
}
