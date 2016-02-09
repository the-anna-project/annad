package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	controlLogResetVerbosityCmd = &cobra.Command{
		Use:   "verbosity",
		Short: "reset log verbosity",
		Long:  "reset log verbosity",
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
		fmt.Printf("%#v\n", maskAny(err))
		os.Exit(1)
	}
}
