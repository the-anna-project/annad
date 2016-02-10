package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	controlLogSetObjectsCmd = &cobra.Command{
		Use:   "objects [object] ...",
		Short: "make Anna set log objects",
		Long:  "make Anna set log objects",
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
		fmt.Printf("%#v\n", maskAny(err))
		os.Exit(1)
	}
}
