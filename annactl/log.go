package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	logCmd = &cobra.Command{
		Use:   "log",
		Short: "control Anna's log behavior",
		Long:  "control Anna's log behavior",
		Run:   logRun,
	}

	logTags struct {
		L string
		O string
		V int
	}
)

func init() {
	logCmd.PersistentFlags().StringVar(&logTags.L, "log-tag-l", "", "level tags of the logger: comma separated")
	logCmd.PersistentFlags().StringVar(&logTags.O, "log-tag-o", "", "object types tag of the logger: comma separated")
	logCmd.PersistentFlags().IntVar(&logTags.V, "log-tag-v", 10, "verbosity tag of the logger: 0 - 15")
}

func logRun(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		cmd.Help()
		os.Exit(1)
	}

	ctx := context.Background()

	if logTags.L != "" {
		err := log.SetLevels(ctx, logTag.L)
		if err != nil {
			fmt.Printf("%#v\n", maskAny(err))
			os.Exit(1)
		}
	}

	if logTags.O != "" {
		err := log.SetObjectTypes(ctx, logTag.O)
		if err != nil {
			fmt.Printf("%#v\n", maskAny(err))
			os.Exit(1)
		}
	}

	err := log.SetVerbosity(ctx, logTag.V)
	if err != nil {
		fmt.Printf("%#v\n", maskAny(err))
		os.Exit(1)
	}
}
