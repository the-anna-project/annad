package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	textinputobject "github.com/the-anna-project/input/object/text"
)

func (a *annactl) InitAnnactlInterfaceTextReadFileCmd() *cobra.Command {
	a.Service().Log().Line("func", "InitAnnactlInterfaceTextReadFileCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "file <file>",
		Short: "Make Anna read plain a file.",
		Long:  "Make Anna read plain a file.",
		Run:   a.ExecAnnactlInterfaceTextReadFileCmd,
		PreRun: func(cmd *cobra.Command, args []string) {
			var err error
			a.sessionID, err = a.GetSessionID()
			panicOnError(err)
		},
	}

	return newCmd
}

func (a *annactl) ExecAnnactlInterfaceTextReadFileCmd(cmd *cobra.Command, args []string) {
	a.Service().Log().Line("func", "ExecAnnactlInterfaceTextReadFileCmd")

	if len(args) == 0 || len(args) >= 2 {
		cmd.HelpFunc()(cmd, nil)
		os.Exit(1)
	}

	ctx := context.Background()

	b, err := a.Service().FS().ReadFile(args[0])
	if err != nil {
		a.Service().Log().Line("msg", "%#v", maskAny(err))
	}

	textInputObject := textinputobject.New()
	err = json.Unmarshal(b, &textInputObject)
	if err != nil {
		a.Service().Log().Line("msg", "%#v", maskAny(err))
	}

	a.Service().Input().Text().Channel() <- textInputObject

	go func() {
		err = a.textInterface.StreamText(ctx)
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
