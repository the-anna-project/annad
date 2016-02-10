package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	controlLogResetObjectsCmd = &cobra.Command{
		Use:   "objects",
		Short: "make Anna reset log objects",
		Long:  "make Anna reset log objects",
		Run:   controlLogResetObjectsRun,
	}
)

func controlLogResetObjectsRun(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		cmd.Help()
		os.Exit(1)
	}

	ctx := context.Background()

	err := logControl.ResetObjects(ctx)
	if err != nil {
		fmt.Printf("%#v\n", maskAny(err))
		os.Exit(1)
	}
}
