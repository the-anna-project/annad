package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlVersionCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitAnnactlVersionCmd")

	newCmd := &cobra.Command{
		Use:   "version",
		Short: "Show current version of the binary.",
		Long:  "Show current version of the binary.",
		Run:   a.ExecAnnactlVersionCmd,
	}

	return newCmd
}

func (a *annactl) ExecAnnactlVersionCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecAnnactlVersionCmd")

	fmt.Printf("%s\n", a.Version)
}
