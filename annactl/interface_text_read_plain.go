package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	readPlainFlags struct {
		Expected string
	}

	interfaceTextReadPlainCmd = &cobra.Command{
		Use:   "plain [text] ...",
		Short: "Make Anna read plain text.",
		Long:  "Make Anna read plain text.",
		Run:   interfaceTextReadPlainRun,
	}
)

func init() {
	interfaceTextReadPlainCmd.PersistentFlags().StringVar(&readPlainFlags.Expected, "expected", "", "output expected to receive with respect to the given input")
}

func interfaceTextReadPlainRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(1)
	}

	ctx := context.Background()

	ID, err := textInterface.ReadPlainWithInput(ctx, strings.Join(args, " "), readPlainFlags.Expected)
	if err != nil {
		fmt.Printf("%#v\n", maskAny(err))
		os.Exit(1)
	}

	data, err := textInterface.ReadPlainWithID(ctx, ID)
	if err != nil {
		fmt.Printf("%#v\n", maskAny(err))
		os.Exit(1)
	}

	fmt.Printf("%s\n", data)
}
