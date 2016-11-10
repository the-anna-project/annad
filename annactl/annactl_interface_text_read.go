package main

import "github.com/spf13/cobra"

func (a *annactl) InitAnnactlInterfaceTextReadCmd() *cobra.Command {
	a.Service().Log().Line("func", "InitAnnactlInterfaceTextReadCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "read",
		Short: "Make Anna read text.",
		Long:  "Make Anna read text.",
		Run:   a.ExecAnnactlInterfaceTextReadCmd,
	}

	// Add sub commands.
	newCmd.AddCommand(a.InitAnnactlInterfaceTextReadPlainCmd())

	return newCmd
}

func (a *annactl) ExecAnnactlInterfaceTextReadCmd(cmd *cobra.Command, args []string) {
	a.Service().Log().Line("func", "ExecAnnactlInterfaceTextReadCmd")

	cmd.HelpFunc()(cmd, nil)
}
