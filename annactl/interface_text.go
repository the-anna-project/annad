package main

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitInterfaceTextCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitInterfaceTextCmd")

	newCmd := &cobra.Command{
		Use:   "text",
		Short: "Text interface for Anna.",
		Long:  "Text interface for Anna.",
		Run:   a.ExecInterfaceTextCmd,
	}

	newCmd.AddCommand(a.InitInterfaceTextReadCmd())

	return newCmd
}

func (a *annactl) ExecInterfaceTextCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecInterfaceTextCmd")

	cmd.Help()
}
