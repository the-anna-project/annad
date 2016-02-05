package main

import (
	"github.com/spf13/cobra"
)

var (
	controlCmd = &cobra.Command{
		Use:   "control",
		Short: "control for Anna's behavior",
		Long:  "control for Anna's behavior",
		Run:   controlRun,
	}
)

func controlRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
