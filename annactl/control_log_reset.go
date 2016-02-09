package main

import (
	"github.com/spf13/cobra"
)

var (
	controlLogResetCmd = &cobra.Command{
		Use:   "reset",
		Short: "reset log configuration",
		Long:  "reset log configuration",
		Run:   controlLogResetRun,
	}
)

func controlLogResetRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
