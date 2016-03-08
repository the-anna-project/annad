package main

import (
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/spec"
)

var (
	controlLogResetLevelsCmd = &cobra.Command{
		Use:   "levels",
		Short: "Make Anna reset log levels.",
		Long:  "Make Anna reset log levels.",
		Run:   controlLogResetLevelsRun,
	}
)

func controlLogResetLevelsRun(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		cmd.Help()
		os.Exit(1)
	}

	ctx := context.Background()

	err := logControl.ResetLevels(ctx)
	if err != nil {
		log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}
}
