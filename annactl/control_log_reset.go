package main

import (
	"github.com/spf13/cobra"
)

var (
	controlLogResetCmd = &cobra.Command{
		Use:   "reset",
		Short: "make Anna reset log configuration",
		Long:  "make Anna reset log configuration",
		Run:   controlLogResetRun,
	}
)

func controlLogResetRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
