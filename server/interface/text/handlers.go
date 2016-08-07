package text

import (
	"github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/spec"
)

// NewHandlers returns HTTP handlers configured for the text interface object.
func NewHandlers(ctx context.Context, ti spec.TextInterface) map[string]*http.Server {
	handlers := map[string]*http.Server{}

	handlers["/interface/text"] = http.NewServer(
		ctx,
		streamTextEndpoint(ti),
		streamTextDecoder,
		streamTextEncoder,
	)

	return handlers
}
