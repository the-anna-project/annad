package main

import (
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlControlLogResetLevelsCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call InitAnnactlControlLogResetLevelsCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "levels",
		Short: "Make Anna reset log levels.",
		Long:  "Make Anna reset log levels.",
		Run:   a.ExecAnnactlControlLogResetLevelsCmd,
	}

	return newCmd
}

func (a *annactl) ExecAnnactlControlLogResetLevelsCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call ExecAnnactlControlLogResetLevelsCmd")

	if len(args) > 0 {
		cmd.HelpFunc()(cmd, nil)
		os.Exit(1)
	}

	ctx := context.Background()

	err := a.LogControl.ResetLevels(ctx)
	if err != nil {
		a.Log.WithTags(spec.Tags{C: nil, L: "F", O: a, V: 1}, "%#v", maskAny(err))
	}
}
