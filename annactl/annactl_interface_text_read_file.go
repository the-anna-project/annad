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

func (a *annactl) InitAnnactlInterfaceTextReadFileCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call InitAnnactlInterfaceTextReadFileCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "file <file>",
		Short: "Make Anna read plain a file.",
		Long:  "Make Anna read plain a file.",
		Run:   a.ExecAnnactlInterfaceTextReadFileCmd,
		PreRun: func(cmd *cobra.Command, args []string) {
			var err error
			a.SessionID, err = a.GetSessionID()
			panicOnError(err)
		},
	}

	return newCmd
}

func (a *annactl) ExecAnnactlInterfaceTextReadFileCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call ExecAnnactlInterfaceTextReadFileCmd")

	if len(args) == 0 || len(args) >= 2 {
		cmd.HelpFunc()(cmd, nil)
		os.Exit(1)
	}

	ctx := context.Background()

	b, err := a.FileSystem.ReadFile(args[0])
	if err != nil {
		a.Log.WithTags(spec.Tags{C: nil, L: "F", O: a, V: 1}, "%#v", maskAny(err))
	}

	textRequest := api.NewEmptyTextRequest()
	err = json.Unmarshal(b, &textRequest)
	if err != nil {
		a.Log.WithTags(spec.Tags{C: nil, L: "F", O: a, V: 1}, "%#v", maskAny(err))
	}

	in := make(chan spec.TextRequest, 1)
	out := make(chan spec.TextResponse, 1000)

	go func() {
		// TODO stream continuously
		in <- textRequest
	}()
	err = a.TextInterface.StreamText(ctx, in, out)
	if err != nil {
		a.Log.WithTags(spec.Tags{C: nil, L: "F", O: a, V: 1}, "%#v", maskAny(err))
	}

	for {
		select {
		case textResponse := <-out:
			fmt.Printf("%s\n", textResponse.GetOutput())
		}
	}
}
