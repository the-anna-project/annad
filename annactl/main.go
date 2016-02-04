// Package annactl implements a command line client for Anna. Cobra CLI is used
// as framework. The commands are simple wrappers around the client package.
package main

import (
	"fmt"
	"net"
	"os"

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
			host, port, err := net.SplitHostPort(globalFlags.Host)
			if err != nil {
				fmt.Printf("%#v\n", maskAny(err))
				os.Exit(1)
			}

			newTextInterfaceConfig := client.DefaultTextInterfaceConfig()
			newTextInterfaceConfig.URL.Host = net.JoinHostPort(host, port)
			textInterface = client.NewTextInterface(newTextInterfaceConfig)
		},
	}
)

func init() {
	mainCmd.PersistentFlags().StringVar(&globalFlags.Host, "host", "127.0.0.1:9119", "host:port to connect to Anna server")
}

func mainRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func main() {
	mainCmd.AddCommand(readPlainCmd)

	mainCmd.Execute()
}
