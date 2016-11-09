package main

import "github.com/spf13/cobra"

func (a *annactl) InitAnnactlCmd() *cobra.Command {
	a.Service().Log().Line("func", "InitAnnactlCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "annactl",
		Short: "Interact with Anna's network API. For more information see https://github.com/xh3b4sd/anna.",
		Long:  "Interact with Anna's network API. For more information see https://github.com/xh3b4sd/anna.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var err error

			// Service collection.
			a.ServiceCollection, err = newServiceCollection()
			panicOnError(err)

			// Text interface.
			a.TextInterface, err = newTextInterface(a.ServiceCollection, a.Flags.GRPCAddr)
			panicOnError(err)
		},
		Run: a.ExecAnnactlCmd,
	}

	// Add sub commands.
	newCmd.AddCommand(a.InitAnnactlInterfaceCmd())
	newCmd.AddCommand(a.InitAnnactlVersionCmd())

	// Define command line flags.
	newCmd.PersistentFlags().StringVar(&a.Flags.GRPCAddr, "grpc-addr", "127.0.0.1:9119", "host:port to bind Anna's gRPC server to")
	newCmd.PersistentFlags().StringVar(&a.Flags.HTTPAddr, "http-addr", "127.0.0.1:9120", "host:port to bind Anna's HTTP server to")

	return newCmd
}

func (a *annactl) ExecAnnactlCmd(cmd *cobra.Command, args []string) {
	a.Service().Log().Line("func", "ExecAnnactlCmd")

	cmd.HelpFunc()(cmd, nil)
}
