package text

import (
	"net/url"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func newReadPlainWithIDEndpoint(URL url.URL, path string) endpoint.Endpoint {
	URL.Path = path
	URL.RawPath = path

	newEndpoint := httptransport.NewClient(
		"POST",
		&URL,
		readPlainEncoder,
		readPlainDecoder,
	).Endpoint()

	return newEndpoint
}

func newReadPlainWithPlainEndpoint(URL url.URL, path string) endpoint.Endpoint {
	URL.Path = path
	URL.RawPath = path

	newEndpoint := httptransport.NewClient(
		"POST",
		&URL,
		readPlainEncoder,
		readPlainDecoder,
	).Endpoint()

	return newEndpoint
}
