package text

import (
	"net/url"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func newStreamTextEndpoint(URL url.URL, path string) endpoint.Endpoint {
	URL.Path = path
	URL.RawPath = path

	newEndpoint := httptransport.NewClient(
		"POST",
		&URL,
		streamTextEncoder,
		streamTextDecoder,
	).Endpoint()

	return newEndpoint
}
