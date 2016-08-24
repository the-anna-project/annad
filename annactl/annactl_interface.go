package main

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlInterfaceCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitAnnactlInterfaceCmd")

	newCmd := &cobra.Command{
		Use:   "interface",
		Short: "Interface for Anna's behaviors.",
		Long:  "Interface for Anna's behaviors.",
		Run:   a.ExecAnnactlInterfaceCmd,
	}

	newCmd.AddCommand(a.InitAnnactlInterfaceTextCmd())

	return newCmd
}

func (a *annactl) ExecAnnactlInterfaceCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecAnnactlInterfaceCmd")

	cmd.HelpFunc()(cmd, nil)
}
