package main

import "github.com/spf13/cobra"

func (a *annactl) InitAnnactlInterfaceTextCmd() *cobra.Command {
	a.Service().Log().Line("func", "InitAnnactlInterfaceTextCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "text",
		Short: "Text interface for Anna.",
		Long:  "Text interface for Anna.",
		Run:   a.ExecAnnactlInterfaceTextCmd,
	}

	// Add sub commands.
	newCmd.AddCommand(a.InitAnnactlInterfaceTextReadCmd())

	return newCmd
}

func (a *annactl) ExecAnnactlInterfaceTextCmd(cmd *cobra.Command, args []string) {
	a.Service().Log().Line("func", "ExecAnnactlInterfaceTextCmd")

	cmd.HelpFunc()(cmd, nil)
}
