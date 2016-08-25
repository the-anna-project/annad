package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlInterfaceTextReadPlainCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitAnnactlInterfaceTextReadPlainCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "plain <input>",
		Short: "Make Anna read plain text.",
		Long:  "Make Anna read plain text.",
		Run:   a.ExecAnnactlInterfaceTextReadPlainCmd,
		PreRun: func(cmd *cobra.Command, args []string) {
			var err error
			a.SessionID, err = a.GetSessionID()
			panicOnError(err)
		},
	}

	// Define command line flags.
	newCmd.PersistentFlags().BoolVar(&a.Flags.InterfaceTextReadPlain.Echo, "echo", false, "echo input by bypassing the neural network")
	//newCmd.PersistentFlags().StringVar(&a.Flags.InterfaceTextReadPlain.Expectation, "expectation", "", "expectation object in JSON format")

	return newCmd
}

func (a *annactl) ExecAnnactlInterfaceTextReadPlainCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecAnnactlInterfaceTextReadPlainCmd")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	in := make(chan spec.TextRequest, 1000)
	out := make(chan spec.TextResponse, 1000)

	go func() {
		a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "Waiting for input.\n")

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			newTextRequestConfig := api.DefaultTextRequestConfig()
			newTextRequestConfig.Echo = a.Flags.InterfaceTextReadPlain.Echo
			newTextRequestConfig.Input = scanner.Text()
			newTextRequestConfig.SessionID = a.SessionID
			newTextRequest, err := api.NewTextRequest(newTextRequestConfig)
			if err != nil {
				a.Log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
			}

			in <- newTextRequest

			err = scanner.Err()
			if err != nil {
				a.Log.WithTags(spec.Tags{L: "E", O: a, T: nil, V: 4}, "%#v", maskAny(err))
			}
		}
	}()

	go func() {
		for {
			select {
			case textResponse := <-out:
				fmt.Printf("%s\n", textResponse.GetOutput())
			}
		}
	}()

	fail := make(chan error, 1)

	go func() {
		select {
		case <-a.Closer:
			cancel()
		case err := <-fail:
			a.Log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
		}
	}()

	err := a.TextInterface.StreamText(ctx, in, out)
	if err != nil {
		fail <- maskAny(err)
	}
}
