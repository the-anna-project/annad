package main

import (
	"github.com/spf13/cobra"
)

var (
	controlLogResetCmd = &cobra.Command{
		Use:   "reset",
		Short: "Make Anna reset log configuration.",
		Long:  "Make Anna reset log configuration.",
		Run:   controlLogResetRun,
	}
)

func controlLogResetRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
