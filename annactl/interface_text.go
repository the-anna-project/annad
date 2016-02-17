package main

import (
	"github.com/spf13/cobra"
)

var (
	interfaceTextCmd = &cobra.Command{
		Use:   "text",
		Short: "Text interface for Anna.",
		Long:  "Text interface for Anna.",
		Run:   interfaceTextRun,
	}
)

func interfaceTextRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
