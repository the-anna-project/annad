package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/object/textinput"
	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlInterfaceTextReadPlainCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call InitAnnactlInterfaceTextReadPlainCmd")

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
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call ExecAnnactlInterfaceTextReadPlainCmd")

	ctx := context.Background()

	go func() {
		a.Log.WithTags(spec.Tags{C: nil, L: "I", O: a, V: 10}, "Waiting for input.\n")

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			newTextInputConfig := textinput.DefaultConfig()
			newTextInputConfig.Echo = a.Flags.InterfaceTextReadPlain.Echo
			newTextInputConfig.Input = scanner.Text()
			newTextInputConfig.SessionID = a.SessionID
			newTextInput, err := textinput.New(newTextInputConfig)
			if err != nil {
				a.Log.WithTags(spec.Tags{C: nil, L: "F", O: a, V: 1}, "%#v", maskAny(err))
			}

			a.Service().TextInput().GetChannel() <- newTextInput

			err = scanner.Err()
			if err != nil {
				a.Log.WithTags(spec.Tags{C: nil, L: "E", O: a, V: 4}, "%#v", maskAny(err))
			}
		}
	}()

	go func() {
		err := a.TextInterface.StreamText(ctx)
		if err != nil {
			a.Log.WithTags(spec.Tags{C: nil, L: "F", O: a, V: 1}, "%#v", maskAny(err))
		}
	}()

	for {
		select {
		case textResponse := <-a.Service().TextOutput().GetChannel():
			fmt.Printf("%s\n", textResponse.GetOutput())
		}
	}
}
