package main

import (
	"github.com/spf13/cobra"
)

var (
	interfaceTextCmd = &cobra.Command{
		Use:   "text",
		Short: "text interface for Anna",
		Long:  "text interface for Anna",
		Run:   interfaceTextRun,
	}
)

func interfaceTextRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
