package main

import (
	"github.com/spf13/cobra"
)

var (
	controlLogCmd = &cobra.Command{
		Use:   "log",
		Short: "control Anna's log behavior",
		Long:  "control Anna's log behavior",
		Run:   controlLogRun,
	}
)

func controlLogRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
