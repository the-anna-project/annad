package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func New() *Command {
	return &Command{}
}

type Command struct {
	// Settings.

	gitCommit string
}

func (c *Command) Execute(cmd *cobra.Command, args []string) {
	fmt.Printf("Git Commit: %s\n", a.gitCommit)
}

func (a *annad) New() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "version",
		Short: "Show version information of the project.",
		Long:  "Show version information of the project.",
		Run:   a.Execute,
	}

	return newCmd
}

func (c *Command) SetGitCommit(gitCommit string) {
	c.gitCommit = gitCommit
}
