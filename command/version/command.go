package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New creates a new version command.
func New() *Command {
	return &Command{}
}

// Command represents the version command.
type Command struct {
	// Settings.

	gitCommit string
}

// Execute represents the cobra run method.
func (c *Command) Execute(cmd *cobra.Command, args []string) {
	fmt.Printf("Git Commit: %s\n", c.gitCommit)
}

// New creates a new cobra command for the version command.
func (c *Command) New() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "version",
		Short: "Show version information of the project.",
		Long:  "Show version information of the project.",
		Run:   c.Execute,
	}

	return newCmd
}

// SetGitCommit sets the git commit for the version command to be displayed.
func (c *Command) SetGitCommit(gitCommit string) {
	c.gitCommit = gitCommit
}
