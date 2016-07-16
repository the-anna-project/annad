package main

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitControlCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitControlCmd")

	newCmd := &cobra.Command{
		Use:   "control",
		Short: "Control for Anna's behaviors.",
		Long:  "Control for Anna's behaviors.",
		Run:   a.ExecControlCmd,
	}

	newCmd.AddCommand(a.InitControlLogCmd())

	return newCmd
}

func (a *annactl) ExecControlCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecControlCmd")

	cmd.HelpFunc()(cmd, nil)
}
