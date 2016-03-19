package main

import (
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitControlLogResetLevelsCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitControlLogResetLevelsCmd")

	newCmd := &cobra.Command{
		Use:   "levels",
		Short: "Make Anna reset log levels.",
		Long:  "Make Anna reset log levels.",
		Run:   a.ExecControlLogResetLevelsCmd,
	}

	return newCmd
}

func (a *annactl) ExecControlLogResetLevelsCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecControlLogResetLevelsCmd")

	if len(args) > 0 {
		cmd.Help()
		os.Exit(1)
	}

	ctx := context.Background()

	err := a.LogControl.ResetLevels(ctx)
	if err != nil {
		a.Log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}
}
