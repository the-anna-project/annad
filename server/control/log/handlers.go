package logcontrol

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/spec"
)

func NewHandlers(ctx context.Context, lc spec.LogControl) map[string]*httptransport.Server {
	handlers := map[string]*httptransport.Server{}

	handlers["/control/log/reset/levels"] = httptransport.NewServer(
		ctx,
		resetLevelsEndpoint(lc),
		resetLevelsDecoder,
		resetLevelsEncoder,
	)

	handlers["/control/log/reset/objects"] = httptransport.NewServer(
		ctx,
		resetObjectsEndpoint(lc),
		resetObjectsDecoder,
		resetObjectsEncoder,
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

	handlers["/control/log/set/objects"] = httptransport.NewServer(
		ctx,
		setObjectsEndpoint(lc),
		setObjectsDecoder,
		setObjectsEncoder,
	)

	handlers["/control/log/set/verbosity"] = httptransport.NewServer(
		ctx,
		setVerbosityEndpoint(lc),
		setVerbosityDecoder,
		setVerbosityEncoder,
	)

	return handlers
}
