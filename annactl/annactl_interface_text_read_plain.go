package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	textinputobject "github.com/the-anna-project/input/object/text"
)

func (a *annactl) InitAnnactlInterfaceTextReadPlainCmd() *cobra.Command {
	a.Service().Log().Line("func", "InitAnnactlInterfaceTextReadPlainCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "plain <input>",
		Short: "Make Anna read plain text.",
		Long:  "Make Anna read plain text.",
		Run:   a.ExecAnnactlInterfaceTextReadPlainCmd,
		PreRun: func(cmd *cobra.Command, args []string) {
			var err error
			a.sessionID, err = a.GetSessionID()
			panicOnError(err)
		},
	}

	// Define command line flags.
	newCmd.PersistentFlags().BoolVar(&a.flags.InterfaceTextReadPlain.Echo, "echo", false, "echo input by bypassing the neural network")
	//newCmd.PersistentFlags().StringVar(&a.flags.InterfaceTextReadPlain.Expectation, "expectation", "", "expectation object in JSON format")

	return newCmd
}

func (a *annactl) ExecAnnactlInterfaceTextReadPlainCmd(cmd *cobra.Command, args []string) {
	a.Service().Log().Line("func", "ExecAnnactlInterfaceTextReadPlainCmd")

	ctx := context.Background()

	go func() {
		a.Service().Log().Line("msg", "Waiting for input.\n")

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			textInputObject := textinputobject.New()
			textInputObject.SetEcho(a.flags.InterfaceTextReadPlain.Echo)
			textInputObject.SetInput(scanner.Text())
			textInputObject.SetSessionID(a.sessionID)

			a.Service().Input().Text().Channel() <- textInputObject

			err := scanner.Err()
			if err != nil {
				a.Service().Log().Line("msg", "%#v", maskAny(err))
			}
		}
	}()

	go func() {
		err := a.textInterface.StreamText(ctx)
		if err != nil {
			a.Service().Log().Line("msg", "%#v", maskAny(err))
		}
	}()

	for {
		select {
		case textResponse := <-a.Service().Output().Text().Channel():
			fmt.Printf("%s\n", textResponse.Output())
		}
	}
}
