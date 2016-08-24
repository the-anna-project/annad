package main

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlInterfaceTextReadCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitAnnactlInterfaceTextReadCmd")

	newCmd := &cobra.Command{
		Use:   "read",
		Short: "Make Anna read text.",
		Long:  "Make Anna read text.",
		Run:   a.ExecAnnactlInterfaceTextReadCmd,
	}

	newCmd.AddCommand(a.InitAnnactlInterfaceTextReadPlainCmd())

	return newCmd
}

func (a *annactl) ExecAnnactlInterfaceTextReadCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecAnnactlInterfaceTextReadCmd")

	cmd.HelpFunc()(cmd, nil)
}
