package main

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlControlLogSetVerbosityCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitAnnactlControlLogSetVerbosityCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "verbosity [verbosity]",
		Short: "Make Anna set log verbosity.",
		Long:  "Make Anna set log verbosity.",
		Run:   a.ExecAnnactlControlLogSetVerbosityCmd,
	}

	return newCmd
}

func (a *annactl) ExecAnnactlControlLogSetVerbosityCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecAnnactlControlLogSetVerbosityCmd")

	if len(args) != 1 {
		cmd.HelpFunc()(cmd, nil)
		os.Exit(1)
	}

	ctx := context.Background()

	v, err := strconv.Atoi(args[0])
	if err != nil {
		a.Log.WithTags(spec.Tags{L: "F", O: nil, T: nil, V: 1}, "%#v", maskAny(err))
	}

	err = a.LogControl.SetVerbosity(ctx, v)
	if err != nil {
		a.Log.WithTags(spec.Tags{L: "F", O: nil, T: nil, V: 1}, "%#v", maskAny(err))
	}
}
