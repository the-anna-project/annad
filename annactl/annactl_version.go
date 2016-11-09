package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (a *annactl) InitAnnactlVersionCmd() *cobra.Command {
	a.Service().Log().Line("func", "InitAnnactlVersionCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "version",
		Short: "Show current version of the binary.",
		Long:  "Show current version of the binary.",
		Run:   a.ExecAnnactlVersionCmd,
	}

	return newCmd
}

func (a *annactl) ExecAnnactlVersionCmd(cmd *cobra.Command, args []string) {
	a.Service().Log().Line("func", "ExecAnnactlVersionCmd")

	fmt.Printf("%s\n", a.Version)
}
