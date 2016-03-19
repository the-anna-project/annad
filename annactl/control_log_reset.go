package main

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitControlLogResetCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitControlLogResetCmd")

	newCmd := &cobra.Command{
		Use:   "reset",
		Short: "Make Anna reset log configuration.",
		Long:  "Make Anna reset log configuration.",
		Run:   a.ExecControlLogResetCmd,
	}

	newCmd.AddCommand(a.InitControlLogResetLevelsCmd())
	newCmd.AddCommand(a.InitControlLogResetObjectsCmd())
	newCmd.AddCommand(a.InitControlLogResetVerbosityCmd())

	return newCmd
}

func (a *annactl) ExecControlLogResetCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecControlLogResetCmd")

	cmd.Help()
}
