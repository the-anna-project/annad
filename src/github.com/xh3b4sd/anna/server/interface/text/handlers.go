package textinterface

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func NewHandlers(ctx context.Context, ti TextInterface) map[string]*httptransport.Server {
	handlers := map[string]*httptransport.Server{}

	handlers["/interface/text/action/fetchurl"] = httptransport.NewServer(
		ctx,
		fetchURLEndpoint(ti),
		fetchURLDecoder,
		fetchURLEncoder,
	)

	handlers["/interface/text/action/readfile"] = httptransport.NewServer(
		ctx,
		readFileEndpoint(ti),
		readFileDecoder,
		readFileEncoder,
	)

	handlers["/interface/text/action/readstream"] = httptransport.NewServer(
		ctx,
		readStreamEndpoint(ti),
		readStreamDecoder,
		readStreamEncoder,
	)

	handlers["/interface/text/action/readplain"] = httptransport.NewServer(
		ctx,
		readPlainEndpoint(ti),
		readPlainDecoder,
		readPlainEncoder,
	)

	return handlers
}
