package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitVersionCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitVersionCmd")

	newCmd := &cobra.Command{
		Use:   "version",
		Short: "Show current version of the binary.",
		Long:  "Show current version of the binary.",
		Run:   a.ExecVersionCmd,
	}

	return newCmd
}

func (a *annactl) ExecVersionCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecVersionCmd")

	fmt.Printf("%s\n", a.Version)
}
