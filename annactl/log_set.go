package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	logSetCmd = &cobra.Command{
		Use:   "set",
		Short: "set configuration for Anna's log behavior",
		Long:  "set configuration for Anna's log behavior",
		Run:   logSetRun,
	}

	setLogTags struct {
		L string
		O string
		V int
	}
)

func init() {
	logSetCmd.PersistentFlags().StringVar(&setLogTags.L, "log-tag-l", "", "set level tags of the logger: comma separated")
	logSetCmd.PersistentFlags().StringVar(&setLogTags.O, "log-tag-o", "", "set object types tag of the logger: comma separated")
	logSetCmd.PersistentFlags().IntVar(&setLogTags.V, "log-tag-v", 10, "set verbosity tag of the logger: 0 - 15")
}

func logSetRun(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		cmd.Help()
		os.Exit(1)
	}

	ctx := context.Background()

	if setLogTags.L != "" {
		err := log.SetLevels(ctx, logTag.L)
		if err != nil {
			fmt.Printf("%#v\n", maskAny(err))
			os.Exit(1)
		}
	}

	if setLogTags.O != "" {
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
