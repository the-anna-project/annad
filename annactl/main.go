package main

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/client"
	serverspec "github.com/xh3b4sd/anna/server/spec"
)

var (
	globalFlags struct {
		Host string
	}

	textInterface serverspec.TextInterface

	mainCmd = &cobra.Command{
		Use:   "annactl",
		Short: "Interact with Anna",
		Long:  "Interact with Anna",
		Run:   mainRun,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			newTextInterfaceConfig := client.DefaultTextInterfaceConfig()
			newTextInterfaceConfig.Host = globalFlags.Host
			textInterface = client.NewTextInterface(newTextInterfaceConfig)
		},
	}
)

func init() {
	mainCmd.PersistentFlags().StringVar(&globalFlags.Host, "host", "127.0.0.1:9119", "Host name to connect to bootxe service")
}

func mainRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func main() {
	mainCmd.AddCommand(readPlainCmd)

	mainCmd.Execute()
}
