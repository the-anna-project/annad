package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitInterfaceTextReadPlainCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitInterfaceTextReadPlainCmd")

	newCmd := &cobra.Command{
		Use:   "plain <input>",
		Short: "Make Anna read plain text.",
		Long:  "Make Anna read plain text.",
		Run:   a.ExecInterfaceTextReadPlainCmd,
		PreRun: func(cmd *cobra.Command, args []string) {
			var err error
			a.SessionID, err = a.GetSessionID()
			panicOnError(err)
		},
	}

	newCmd.PersistentFlags().StringVar(&a.Flags.InterfaceTextReadPlain.Expectation, "expectation", "", "expectation object in JSON format")

	return newCmd
}

func (a *annactl) ExecInterfaceTextReadPlainCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecInterfaceTextReadPlainCmd")

	if len(args) == 0 {
		cmd.HelpFunc()(cmd, nil)
		os.Exit(1)
	}

	ctx := context.Background()

	var expectationRequest api.ExpectationRequest
	err := json.Unmarshal([]byte(a.Flags.InterfaceTextReadPlain.Expectation), &expectationRequest)
	if err != nil {
		a.Log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}

	newTextRequestConfig := api.DefaultTextRequestConfig()
	newTextRequestConfig.ExpectationRequest = expectationRequest
	newTextRequestConfig.Input = strings.Join(args, " ")
	newTextRequestConfig.SessionID = a.SessionID
	newTextRequest, err := api.NewTextRequest(newTextRequestConfig)
	if err != nil {
		a.Log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}

	in := make(chan spec.TextRequest, 1)
	out := make(chan spec.TextResponse, 1000)

	go func() {
		// TODO stream continuously
		in <- newTextRequest
	}()

	err = a.TextInterface.StreamText(ctx, in, out)
	if err != nil {
		a.Log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}

	for {
		select {
		case textResponse := <-out:
			fmt.Printf("%s\n", textResponse.GetOutput())
		}
	}
}
