package main

import (
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/spec"
)

var (
	controlLogResetVerbosityCmd = &cobra.Command{
		Use:   "verbosity",
		Short: "Make Anna reset log verbosity.",
		Long:  "Make Anna reset log verbosity.",
		Run:   controlLogResetVerbosityRun,
	}
)

func controlLogResetVerbosityRun(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		cmd.Help()
		os.Exit(1)
	}

	ctx := context.Background()

	err := logControl.ResetVerbosity(ctx)
	if err != nil {
		log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}
}
