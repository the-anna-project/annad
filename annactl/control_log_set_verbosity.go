package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	controlLogSetVerbosityCmd = &cobra.Command{
		Use:   "verbosity [verbosity]",
		Short: "make Anna set log verbosity",
		Long:  "make Anna set log verbosity",
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
		fmt.Printf("%#v\n", maskAny(err))
		os.Exit(1)
	}

	err = logControl.SetVerbosity(ctx, v)
	if err != nil {
		fmt.Printf("%#v\n", maskAny(err))
		os.Exit(1)
	}
}
