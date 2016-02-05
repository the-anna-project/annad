package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	controlLogResetLevelsCmd = &cobra.Command{
		Use:   "levels",
		Short: "make Anna reset log levels",
		Long:  "make Anna reset log levels",
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
		fmt.Printf("%#v\n", maskAny(err))
		os.Exit(1)
	}
}
