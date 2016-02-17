package main

import (
	"github.com/spf13/cobra"
)

var (
	controlCmd = &cobra.Command{
		Use:   "control",
		Short: "Control for Anna's behaviors.",
		Long:  "Control for Anna's behaviors.",
		Run:   controlRun,
	}
)

func controlRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
