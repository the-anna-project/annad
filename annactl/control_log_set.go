package main

import (
	"github.com/spf13/cobra"
)

var (
	controlLogSetCmd = &cobra.Command{
		Use:   "set",
		Short: "set log configuration",
		Long:  "set log configuration",
		Run:   controlLogSetRun,
	}
)

func controlLogSetRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
