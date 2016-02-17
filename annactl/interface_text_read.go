package main

import (
	"github.com/spf13/cobra"
)

var (
	interfaceTextReadCmd = &cobra.Command{
		Use:   "read",
		Short: "Make Anna read text.",
		Long:  "Make Anna read text.",
		Run:   interfaceTextReadRun,
	}
)

func interfaceTextReadRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
