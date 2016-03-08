package main

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/spec"
)

var (
	controlLogSetObjectsCmd = &cobra.Command{
		Use:   "objects [object] ...",
		Short: "Make Anna set log objects.",
		Long:  "Make Anna set log objects.",
		Run:   controlLogSetObjectsRun,
	}
)

func controlLogSetObjectsRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(1)
	}

	ctx := context.Background()

	err := logControl.SetObjects(ctx, strings.Join(args, ","))
	if err != nil {
		log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}
}
