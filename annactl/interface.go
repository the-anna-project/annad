package main

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitInterfaceCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitInterfaceCmd")

	newCmd := &cobra.Command{
		Use:   "interface",
		Short: "Interface for Anna's behaviors.",
		Long:  "Interface for Anna's behaviors.",
		Run:   a.ExecInterfaceCmd,
	}

	newCmd.AddCommand(a.InitInterfaceTextCmd())

	return newCmd
}

func (a *annactl) ExecInterfaceCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecInterfaceCmd")

	cmd.HelpFunc()(cmd, nil)
}
