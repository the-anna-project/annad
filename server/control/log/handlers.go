package logcontrol

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"

	serverspec "github.com/xh3b4sd/anna/server/spec"
)

func NewHandlers(ctx context.Context, lc serverspec.LogControl) map[string]*httptransport.Server {
	handlers := map[string]*httptransport.Server{}

	handlers["/control/log/reset/levels"] = httptransport.NewServer(
		ctx,
		resetLevelsEndpoint(lc),
		resetLevelsDecoder,
		resetLevelsEncoder,
	)

	handlers["/control/log/reset/objecttypes"] = httptransport.NewServer(
		ctx,
		resetObjectTypesEndpoint(lc),
		resetObjectTypesDecoder,
		resetObjectTypesEncoder,
	)

	handlers["/control/log/reset/verbosity"] = httptransport.NewServer(
		ctx,
		resetVerbosityEndpoint(lc),
		resetVerbosityDecoder,
		resetVerbosityEncoder,
	)

	handlers["/control/log/set/levels"] = httptransport.NewServer(
		ctx,
		setLevelsEndpoint(lc),
		setLevelsDecoder,
		setLevelsEncoder,
	)

	handlers["/control/log/set/objecttypes"] = httptransport.NewServer(
		ctx,
		setObjectTypesEndpoint(lc),
		setObjectTypesDecoder,
		setObjectTypesEncoder,
	)

	handlers["/control/log/set/verbosity"] = httptransport.NewServer(
		ctx,
		setVerbosityEndpoint(lc),
		setVerbosityDecoder,
		setVerbosityEncoder,
	)

	return handlers
}
