// Package annactl implements a command line client for Anna. Cobra CLI is used
// as framework. The commands are simple wrappers around the client package.
package main

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/client/control/log"
	"github.com/xh3b4sd/anna/client/interface/text"
	serverspec "github.com/xh3b4sd/anna/server/spec"
)

var (
	globalFlags struct {
		Host string
	}

	textInterface serverspec.TextInterface
	logControl    serverspec.LogControl

	mainCmd = &cobra.Command{
		Use:   "annactl",
		Short: "interact with Anna",
		Long:  "interact with Anna",
		Run:   mainRun,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			host, port, err := net.SplitHostPort(globalFlags.Host)
			if err != nil {
				fmt.Printf("%#v\n", maskAny(err))
				os.Exit(1)
			}
			hostport := net.JoinHostPort(host, port)

			newTextInterfaceConfig := textinterface.DefaultConfig()
			newTextInterfaceConfig.URL.Host = hostport
			textInterface = textinterface.NewTextInterface(newTextInterfaceConfig)

			newLogControlConfig := logcontrol.DefaultConfig()
			newLogControlConfig.URL.Host = hostport
			logControl = logcontrol.NewLogControl(newLogControlConfig)
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
	controlLogResetCmd.AddCommand(controlLogResetLevelsCmd)
	controlLogResetCmd.AddCommand(controlLogResetObjectsCmd)
	controlLogResetCmd.AddCommand(controlLogResetVerbosityCmd)
	controlLogSetCmd.AddCommand(controlLogSetLevelsCmd)
	controlLogSetCmd.AddCommand(controlLogSetObjectsCmd)
	controlLogSetCmd.AddCommand(controlLogSetVerbosityCmd)
	controlLogCmd.AddCommand(controlLogResetCmd)
	controlLogCmd.AddCommand(controlLogSetCmd)
	controlCmd.AddCommand(controlLogCmd)
	mainCmd.AddCommand(controlCmd)

	mainCmd.AddCommand(readPlainCmd)

	mainCmd.Execute()
}
