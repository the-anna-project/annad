package main

import (
	"math/rand"
	"time"

	"github.com/xh3b4sd/anna/annactl/command"
)

var (
	gitCommit      string
	goArch         string
	goOS           string
	goVersion      string
	projectVersion string
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	annadCommand := command.New()

	annadCommand.VersionCommand().SetGitCommit(gitCommit)
	annadCommand.VersionCommand().SetGoArch(goArch)
	annadCommand.VersionCommand().SetGoOS(goOS)
	annadCommand.VersionCommand().SetGoVersion(goVersion)
	annadCommand.VersionCommand().SetProjectVersion(projectVersion)

	annadCommand.New().Execute()
}
