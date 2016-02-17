package main

import (
	"github.com/spf13/cobra"
)

var (
	controlLogCmd = &cobra.Command{
		Use:   "log",
		Short: "Log control for Anna.",
		Long:  "Log control for Anna.",
		Run:   controlLogRun,
	}
)

func controlLogRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
