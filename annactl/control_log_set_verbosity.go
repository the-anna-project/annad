package main

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/spec"
)

var (
	controlLogSetVerbosityCmd = &cobra.Command{
		Use:   "verbosity [verbosity]",
		Short: "Make Anna set log verbosity.",
		Long:  "Make Anna set log verbosity.",
		Run:   controlLogSetVerbosityRun,
	}
)

func controlLogSetVerbosityRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Help()
		os.Exit(1)
	}

	ctx := context.Background()

	v, err := strconv.Atoi(args[0])
	if err != nil {
		log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}

	err = logControl.SetVerbosity(ctx, v)
	if err != nil {
		log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}
}
