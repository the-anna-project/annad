package main

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlControlLogSetObjectsCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call InitAnnactlControlLogSetObjectsCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "objects [object] ...",
		Short: "Make Anna set log objects.",
		Long:  "Make Anna set log objects.",
		Run:   a.ExecAnnactlControlLogSetObjectsCmd,
	}

	return newCmd
}

func (a *annactl) ExecAnnactlControlLogSetObjectsCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call ExecAnnactlControlLogSetObjectsCmd")

	if len(args) == 0 {
		cmd.HelpFunc()(cmd, nil)
		os.Exit(1)
	}

	ctx := context.Background()

	err := a.LogControl.SetObjects(ctx, strings.Join(args, ","))
	if err != nil {
		a.Log.WithTags(spec.Tags{C: nil, L: "F", O: a, V: 1}, "%#v", maskAny(err))
	}
}
