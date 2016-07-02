package text

import (
	"net/url"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func newGetResponseForIDEndpoint(URL url.URL, path string) endpoint.Endpoint {
	URL.Path = path
	URL.RawPath = path

	newEndpoint := httptransport.NewClient(
		"POST",
		&URL,
		getResponseForIDEncoder,
		getResponseForIDDecoder,
	).Endpoint()

	return newEndpoint
}

func newReadCoreRequestEndpoint(URL url.URL, path string) endpoint.Endpoint {
	URL.Path = path
	URL.RawPath = path

	newEndpoint := httptransport.NewClient(
		"POST",
		&URL,
		readCoreRequestEncoder,
		readCoreRequestDecoder,
	).Endpoint()

	return newEndpoint
}
