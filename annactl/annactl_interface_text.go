package main

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlInterfaceTextCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call InitAnnactlInterfaceTextCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "text",
		Short: "Text interface for Anna.",
		Long:  "Text interface for Anna.",
		Run:   a.ExecAnnactlInterfaceTextCmd,
	}

	// Add sub commands.
	newCmd.AddCommand(a.InitAnnactlInterfaceTextReadCmd())

	return newCmd
}

func (a *annactl) ExecAnnactlInterfaceTextCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call ExecAnnactlInterfaceTextCmd")

	cmd.HelpFunc()(cmd, nil)
}
