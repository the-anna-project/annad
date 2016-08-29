package main

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlInterfaceTextReadCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call InitAnnactlInterfaceTextReadCmd")

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
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call ExecAnnactlInterfaceTextReadCmd")

	cmd.HelpFunc()(cmd, nil)
}
