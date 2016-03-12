package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitInterfaceTextReadPlainCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitInterfaceTextReadPlainCmd")

	newCmd := &cobra.Command{
		Use:   "plain [text] ...",
		Short: "Make Anna read plain text.",
		Long:  "Make Anna read plain text.",
		Run:   a.ExecInterfaceTextReadPlainCmd,
		PreRun: func(cmd *cobra.Command, args []string) {
			var err error
			a.SessionID, err = a.GetSessionID()
			panicOnError(err)
		},
	}

	newCmd.PersistentFlags().StringVar(&a.Flags.InterfaceTextReadPlain.Expected, "expected", "", "output expected to receive with respect to the given input")

	return newCmd
}

func (a *annactl) ExecInterfaceTextReadPlainCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecInterfaceTextReadPlainCmd")

	if len(args) == 0 {
		cmd.Help()
		os.Exit(1)
	}

	ctx := context.Background()

	ID, err := a.TextInterface.ReadPlainWithInput(ctx, strings.Join(args, " "), a.Flags.InterfaceTextReadPlain.Expected, a.SessionID)
	if err != nil {
		a.Log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}

	data, err := a.TextInterface.ReadPlainWithID(ctx, ID)
	if err != nil {
		a.Log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}

	fmt.Printf("%s\n", data)
}
