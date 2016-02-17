package main

import (
	"github.com/spf13/cobra"
)

var (
	controlLogSetCmd = &cobra.Command{
		Use:   "set",
		Short: "Make Anna set log configuration.",
		Long:  "Make Anna set log configuration.",
		Run:   controlLogSetRun,
	}
)

func controlLogSetRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
