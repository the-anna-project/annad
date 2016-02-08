package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	logResetCmd = &cobra.Command{
		Use:   "reset",
		Short: "reset configuration for Anna's log behavior",
		Long:  "reset configuration for Anna's log behavior",
		Run:   logResetRun,
	}

	resetLogTags struct {
		L bool
		O bool
		V bool
	}
)

func init() {
	logResetCmd.PersistentFlags().BoolVar(&resetLogTags.L, "log-tag-l", false, "reset level tags of the logger: all")
	logResetCmd.PersistentFlags().BoolVar(&resetLogTags.O, "log-tag-o", false, "reset object types tag of the logger: all")
	logResetCmd.PersistentFlags().BoolVar(&resetLogTags.V, "log-tag-v", false, "reset verbosity tag of the logger: 10")
}

func logResetRun(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		cmd.Help()
		os.Exit(1)
	}

	if !resetLogTags.L && !resetLogTags.O && !resetLogTags.V {
		fmt.Printf("no log tag provided to be reset\n")
		os.Exit(1)
	}

	ctx := context.Background()

	if resetLogTags.L {
		err := log.ResetLevels(ctx)
		if err != nil {
			fmt.Printf("%#v\n", maskAny(err))
			os.Exit(1)
		}
	}

	if resetLogTags.O {
		err := log.ResetObjectTypes(ctx)
		if err != nil {
			fmt.Printf("%#v\n", maskAny(err))
			os.Exit(1)
		}
	}

	if resetLogTags.V {
		err := log.ResetVerbosity(ctx)
		if err != nil {
			fmt.Printf("%#v\n", maskAny(err))
			os.Exit(1)
		}
	}
}
