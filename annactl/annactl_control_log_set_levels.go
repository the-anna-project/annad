package main

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlControlLogSetLevelsCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitAnnactlControlLogSetLevelsCmd")

	newCmd := &cobra.Command{
		Use:   "levels [level] ...",
		Short: "Make Anna set log levels.",
		Long:  "Make Anna set log levels.",
		Run:   a.ExecAnnactlControlLogSetLevelsCmd,
	}

	return newCmd
}

func (a *annactl) ExecAnnactlControlLogSetLevelsCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecAnnactlControlLogSetLevelsCmd")

	if len(args) == 0 {
		cmd.HelpFunc()(cmd, nil)
		os.Exit(1)
	}

	ctx := context.Background()

	err := a.LogControl.SetLevels(ctx, strings.Join(args, ","))
	if err != nil {
		a.Log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}
}
