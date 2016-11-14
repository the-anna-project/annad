package main

import (
	"math/rand"
	"time"

	"github.com/xh3b4sd/anna/command"
	"github.com/xh3b4sd/anna/command/boot"
	"github.com/xh3b4sd/anna/command/version"
	"github.com/xh3b4sd/anna/object/config"
	"github.com/xh3b4sd/anna/object/config/endpoint"
	"github.com/xh3b4sd/anna/object/config/endpoint/metric"
	"github.com/xh3b4sd/anna/object/config/endpoint/text"
	"github.com/xh3b4sd/anna/object/config/space"
	connectionspace "github.com/xh3b4sd/anna/object/config/space/connection"
	"github.com/xh3b4sd/anna/object/config/space/dimension"
	"github.com/xh3b4sd/anna/object/config/space/peer"
	"github.com/xh3b4sd/anna/object/config/storage"
	connectionstorage "github.com/xh3b4sd/anna/object/config/storage/connection"
	"github.com/xh3b4sd/anna/object/config/storage/feature"
	"github.com/xh3b4sd/anna/object/config/storage/general"
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
	bootCommand := boot.New()
	versionCommand := version.New()

	configCollection := config.NewCollection()
	configCollection.SetEndpointCollection(endpoint.NewCollection())
	configCollection.SetSpaceCollection(space.NewCollection())
	configCollection.SetStorageCollection(storage.NewCollection())
	configCollection.Endpoint().SetMetric(metric.New())
	configCollection.Endpoint().SetText(text.New())
	configCollection.Space().SetConnection(connectionspace.New())
	configCollection.Space().SetDimension(dimension.New())
	configCollection.Space().SetPeer(peer.New())
	configCollection.Storage().SetConnection(connectionstorage.New())
	configCollection.Storage().SetFeature(feature.New())
	configCollection.Storage().SetGeneral(general.New())

	bootCommand.SetConfigCollection(configCollection)
	bootCommand.SetGitCommit(gitCommit)
	bootCommand.SetGoArch(goArch)
	bootCommand.SetGoOS(goOS)
	bootCommand.SetGoVersion(goVersion)
	bootCommand.SetProjectVersion(projectVersion)
	versionCommand.SetGitCommit(gitCommit)
	versionCommand.SetGoArch(goArch)
	versionCommand.SetGoOS(goOS)
	versionCommand.SetGoVersion(goVersion)
	versionCommand.SetProjectVersion(projectVersion)

	annadCommand.SetBootCommand(bootCommand)
	annadCommand.SetVersionCommand(versionCommand)

	annadCommand.New().Execute()
}
