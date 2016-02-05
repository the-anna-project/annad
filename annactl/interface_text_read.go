package main

import (
	"github.com/spf13/cobra"
)

var (
	interfaceTextReadCmd = &cobra.Command{
		Use:   "read",
		Short: "make Anna read text",
		Long:  "make Anna read text",
		Run:   interfaceTextReadRun,
	}
)

func interfaceTextReadRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
