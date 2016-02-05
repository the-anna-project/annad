package main

import (
	"github.com/spf13/cobra"
)

var (
	interfaceCmd = &cobra.Command{
		Use:   "interface",
		Short: "interface for Anna's behavior",
		Long:  "interface for Anna's behavior",
		Run:   interfaceRun,
	}
)

func interfaceRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
