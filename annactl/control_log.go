package main

import (
	"github.com/spf13/cobra"
)

var (
	controlLogCmd = &cobra.Command{
		Use:   "log",
		Short: "log control for Anna",
		Long:  "log control for Anna",
		Run:   controlLogRun,
	}
)

func controlLogRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
