package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annad) InitAnnadVersionCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call InitAnnadVersionCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "version",
		Short: "Show current version of the binary.",
		Long:  "Show current version of the binary.",
		Run:   a.ExecAnnadVersionCmd,
	}

	return newCmd
}

func (a *annad) ExecAnnadVersionCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call ExecAnnadVersionCmd")

	fmt.Printf("%s\n", a.Version)
}
