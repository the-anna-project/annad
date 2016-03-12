package main

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitInterfaceTextReadCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitInterfaceTextReadCmd")

	newCmd := &cobra.Command{
		Use:   "read",
		Short: "Make Anna read text.",
		Long:  "Make Anna read text.",
		Run:   a.ExecInterfaceTextReadCmd,
	}

	newCmd.AddCommand(a.InitInterfaceTextReadPlainCmd())

	return newCmd
}

func (a *annactl) ExecInterfaceTextReadCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecInterfaceTextReadCmd")

	cmd.Help()
}
