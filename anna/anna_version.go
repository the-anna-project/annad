package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *anna) InitAnnaVersionCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call InitAnnaVersionCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "version",
		Short: "Show current version of the binary.",
		Long:  "Show current version of the binary.",
		Run:   a.ExecAnnaVersionCmd,
	}

	return newCmd
}

func (a *anna) ExecAnnaVersionCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call ExecAnnaVersionCmd")

	fmt.Printf("%s\n", a.Version)
}
