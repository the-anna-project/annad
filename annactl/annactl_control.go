package main

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlControlCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call InitAnnactlControlCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "control",
		Short: "Control for Anna's behaviours.",
		Long:  "Control for Anna's behaviours.",
		Run:   a.ExecAnnactlControlCmd,
	}

	// Add sub commands.
	newCmd.AddCommand(a.InitAnnactlControlLogCmd())

	return newCmd
}

func (a *annactl) ExecAnnactlControlCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call ExecAnnactlControlCmd")

	cmd.HelpFunc()(cmd, nil)
}
