package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "show current version",
		Long:  "show current version",
		Run:   versionRun,
	}
)

func versionRun(cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n", version)
}
