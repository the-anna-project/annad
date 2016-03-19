package main

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitControlLogCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitControlLogCmd")

	newCmd := &cobra.Command{
		Use:   "log",
		Short: "Log control for Anna.",
		Long:  "Log control for Anna.",
		Run:   a.ExecControlLogCmd,
	}

	newCmd.AddCommand(a.InitControlLogResetCmd())
	newCmd.AddCommand(a.InitControlLogSetCmd())

	return newCmd
}

func (a *annactl) ExecControlLogCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecControlLogCmd")

	cmd.Help()
}
