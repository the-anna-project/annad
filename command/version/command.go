package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (a *annad) InitAnnadVersionCmd() *cobra.Command {
	a.Service().Log().Line("func", "InitAnnadVersionCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "version",
		Short: "Show current version of the binary.",
		Long:  "Show current version of the binary.",
		Run:   a.ExecAnnadVersionCmd,
	}

	return newCmd
}

func (a *annad) ExecAnnadVersionCmd(cmd *cobra.Command, args []string) {
	a.Service().Log().Line("func", "ExecAnnadVersionCmd")

	fmt.Printf("%s\n", a.version)
}
