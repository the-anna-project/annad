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

	gitCommit      string
	goArch         string
	goOS           string
	goVersion      string
	projectVersion string
}

// Execute represents the cobra run method.
func (c *Command) Execute(cmd *cobra.Command, args []string) {
	fmt.Printf("Git Commit:         %s\n", c.gitCommit)
	fmt.Printf("Go Arch:            %s\n", c.goArch)
	fmt.Printf("Go OS:              %s\n", c.goOS)
	fmt.Printf("Go Version:         %s\n", c.goVersion)
	fmt.Printf("Project Version:    %s\n", c.projectVersion)
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

// SetGoArch sets the go architecture for the version command to be displayed.
func (c *Command) SetGoArch(goArch string) {
	c.goArch = goArch
}

// SetGoOS sets the go OS for the version command to be displayed.
func (c *Command) SetGoOS(goOS string) {
	c.goOS = goOS
}

// SetGoVersion sets the go version for the version command to be displayed.
func (c *Command) SetGoVersion(goVersion string) {
	c.goVersion = goVersion
}

// SetProjectVersion sets the project version for the version command to be displayed.
func (c *Command) SetProjectVersion(projectVersion string) {
	c.projectVersion = projectVersion
}
