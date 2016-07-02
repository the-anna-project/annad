package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitInterfaceTextReadFileCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitInterfaceTextReadFileCmd")

	newCmd := &cobra.Command{
		Use:   "file <file>",
		Short: "Make Anna read plain a file.",
		Long:  "Make Anna read plain a file.",
		Run:   a.ExecInterfaceTextReadFileCmd,
		PreRun: func(cmd *cobra.Command, args []string) {
			var err error
			a.SessionID, err = a.GetSessionID()
			panicOnError(err)
		},
	}

	return newCmd
}

func (a *annactl) ExecInterfaceTextReadFileCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecInterfaceTextReadFileCmd")

	if len(args) == 0 || len(args) >= 2 {
		cmd.Help()
		os.Exit(1)
	}

	ctx := context.Background()

	b, err := a.FileSystem.ReadFile(args[0])
	if err != nil {
		a.Log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}

	var coreRequest api.CoreRequest
	err = json.Unmarshal(b, &coreRequest)
	if err != nil {
		a.Log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}

	ID, err := a.TextInterface.ReadCoreRequest(ctx, coreRequest, a.SessionID)
	if err != nil {
		a.Log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}

	response, err := a.TextInterface.GetResponseForID(ctx, ID)
	if err != nil {
		a.Log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}

	fmt.Printf("%s\n", response)
}
