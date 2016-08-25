package main

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlControlLogCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitAnnactlControlLogCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "log",
		Short: "Log control for Anna.",
		Long:  "Log control for Anna.",
		Run:   a.ExecAnnactlControlLogCmd,
	}

	// Add sub commands.
	newCmd.AddCommand(a.InitAnnactlControlLogResetCmd())
	newCmd.AddCommand(a.InitAnnactlControlLogSetCmd())

	return newCmd
}

func (a *annactl) ExecAnnactlControlLogCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecAnnactlControlLogCmd")

	cmd.HelpFunc()(cmd, nil)
}
