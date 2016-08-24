package main

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlControlLogSetCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitAnnactlControlLogSetCmd")

	newCmd := &cobra.Command{
		Use:   "set",
		Short: "Make Anna set log configuration.",
		Long:  "Make Anna set log configuration.",
		Run:   a.ExecAnnactlControlLogSetCmd,
	}

	newCmd.AddCommand(a.InitAnnactlControlLogSetLevelsCmd())
	newCmd.AddCommand(a.InitAnnactlControlLogSetObjectsCmd())
	newCmd.AddCommand(a.InitAnnactlControlLogSetVerbosityCmd())

	return newCmd
}

func (a *annactl) ExecAnnactlControlLogSetCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecAnnactlControlLogSetCmd")

	cmd.HelpFunc()(cmd, nil)
}
