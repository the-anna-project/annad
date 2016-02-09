package main

import (
	"github.com/spf13/cobra"
)

var (
	controlCmd = &cobra.Command{
		Use:   "control",
		Short: "control Anna's behavior",
		Long:  "control Anna's behavior",
		Run:   controlRun,
	}
)

func controlRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
