package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	interfaceTextReadPlainCmd = &cobra.Command{
		Use:   "plain [text] ...",
		Short: "make Anna read plain text",
		Long:  "make Anna read plain text",
		Run:   interfaceTextReadPlainRun,
	}
)

func interfaceTextReadPlainRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(1)
	}

	ctx := context.Background()

	ID, err := textInterface.ReadPlainWithPlain(ctx, strings.Join(args, " "))
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
