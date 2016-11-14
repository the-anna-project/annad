package main

import (
	"math/rand"
	"time"

	"github.com/xh3b4sd/anna/command"
	"github.com/xh3b4sd/anna/command/boot"
	"github.com/xh3b4sd/anna/command/version"
	"github.com/xh3b4sd/anna/config"
	"github.com/xh3b4sd/anna/config/endpoint"
	"github.com/xh3b4sd/anna/config/endpoint/metric"
	"github.com/xh3b4sd/anna/config/endpoint/text"
	"github.com/xh3b4sd/anna/config/space"
	connectionspace "github.com/xh3b4sd/anna/config/space/connection"
	"github.com/xh3b4sd/anna/config/space/dimension"
	"github.com/xh3b4sd/anna/config/space/peer"
	"github.com/xh3b4sd/anna/config/storage"
	connectionstorage "github.com/xh3b4sd/anna/config/storage/connection"
	"github.com/xh3b4sd/anna/config/storage/feature"
	"github.com/xh3b4sd/anna/config/storage/general"
)

var (
	// gitCommit is the git commit the project is build with. It is given via
	// build flags that inject the git commit hash.
	gitCommit string
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	//
	annadCommand := command.New()
	bootCommand := boot.New()
	versionCommand := version.New()

	//
	configCollection := config.NewCollection()
	//
	configCollection.SetEndpointCollection(endpoint.NewCollection())
	configCollection.SetSpaceCollection(space.NewCollection())
	configCollection.SetStorageCollection(storage.NewCollection())
	//
	configCollection.Endpoint().SetMetric(metric.New())
	configCollection.Endpoint().SetText(text.New())
	//
	configCollection.Space().SetConnection(connectionspace.New())
	configCollection.Space().SetDimension(dimension.New())
	configCollection.Space().SetPeer(peer.New())
	//
	configCollection.Storage().SetConnection(connectionstorage.New())
	configCollection.Storage().SetFeature(feature.New())
	configCollection.Storage().SetGeneral(general.New())

	//
	bootCommand.SetConfigCollection(configCollection)
	bootCommand.SetGitCommit(gitCommit)
	versionCommand.SetGitCommit(gitCommit)

	//
	annadCommand.SetBootCommand(bootCommand)
	annadCommand.SetVersionCommand(versionCommand)

	//
	annadCommand.Execute()
}
