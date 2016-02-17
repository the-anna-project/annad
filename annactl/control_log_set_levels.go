package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	controlLogSetLevelsCmd = &cobra.Command{
		Use:   "levels [level] ...",
		Short: "Make Anna set log levels.",
		Long:  "Make Anna set log levels.",
		Run:   controlLogSetLevelsRun,
	}
)

func controlLogSetLevelsRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(1)
	}

	ctx := context.Background()

	err := logControl.SetLevels(ctx, strings.Join(args, ","))
	if err != nil {
		fmt.Printf("%#v\n", maskAny(err))
		os.Exit(1)
	}
}
