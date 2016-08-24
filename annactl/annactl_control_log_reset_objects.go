package main

import (
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlControlLogResetObjectsCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitAnnactlControlLogResetObjectsCmd")

	newCmd := &cobra.Command{
		Use:   "objects",
		Short: "Make Anna reset log objects.",
		Long:  "Make Anna reset log objects.",
		Run:   a.ExecAnnactlControlLogResetObjectsCmd,
	}

	return newCmd
}

func (a *annactl) ExecAnnactlControlLogResetObjectsCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecAnnactlControlLogResetObjectsCmd")

	if len(args) > 0 {
		cmd.HelpFunc()(cmd, nil)
		os.Exit(1)
	}

	ctx := context.Background()

	err := a.LogControl.ResetObjects(ctx)
	if err != nil {
		a.Log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}
}
