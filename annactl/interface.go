package main

import (
	"github.com/spf13/cobra"
)

var (
	interfaceCmd = &cobra.Command{
		Use:   "interface",
		Short: "Interface for Anna's behaviors.",
		Long:  "Interface for Anna's behaviors.",
		Run:   interfaceRun,
	}
)

func interfaceRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
