package main

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitControlLogSetCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitControlLogSetCmd")

	newCmd := &cobra.Command{
		Use:   "set",
		Short: "Make Anna set log configuration.",
		Long:  "Make Anna set log configuration.",
		Run:   a.ExecControlLogSetCmd,
	}

	newCmd.AddCommand(a.InitControlLogSetLevelsCmd())
	newCmd.AddCommand(a.InitControlLogSetObjectsCmd())
	newCmd.AddCommand(a.InitControlLogSetVerbosityCmd())

	return newCmd
}

func (a *annactl) ExecControlLogSetCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecControlLogSetCmd")

	cmd.HelpFunc()(cmd, nil)
}
