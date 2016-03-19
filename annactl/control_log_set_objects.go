package main

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitControlLogSetObjectsCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitControlLogSetObjectsCmd")

	newCmd := &cobra.Command{
		Use:   "objects [object] ...",
		Short: "Make Anna set log objects.",
		Long:  "Make Anna set log objects.",
		Run:   a.ExecControlLogSetObjectsCmd,
	}

	return newCmd
}

func (a *annactl) ExecControlLogSetObjectsCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecControlLogSetObjectsCmd")

	if len(args) == 0 {
		cmd.Help()
		os.Exit(1)
	}

	ctx := context.Background()

	err := a.LogControl.SetObjects(ctx, strings.Join(args, ","))
	if err != nil {
		a.Log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}
}
