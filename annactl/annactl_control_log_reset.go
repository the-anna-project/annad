package main

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlControlLogResetCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitAnnactlControlLogResetCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "reset",
		Short: "Make Anna reset log configuration.",
		Long:  "Make Anna reset log configuration.",
		Run:   a.ExecAnnactlControlLogResetCmd,
	}

	// Add sub commands.
	newCmd.AddCommand(a.InitAnnactlControlLogResetLevelsCmd())
	newCmd.AddCommand(a.InitAnnactlControlLogResetObjectsCmd())
	newCmd.AddCommand(a.InitAnnactlControlLogResetVerbosityCmd())

	return newCmd
}

func (a *annactl) ExecAnnactlControlLogResetCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecAnnactlControlLogResetCmd")

	cmd.HelpFunc()(cmd, nil)
}
