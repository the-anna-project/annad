package logcontrol

import (
	"net/url"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func newResetLevelsEndpoint(URL *url.URL, path string) endpoint.Endpoint {
	URL.Path = path
	URL.RawPath = path

	newEndpoint := httptransport.NewClient(
		"POST",
		URL,
		resetLevelsEncoder,
		resetLevelsDecoder,
	).Endpoint()

	return newEndpoint
}

func newResetObjectTypesEndpoint(URL *url.URL, path string) endpoint.Endpoint {
	URL.Path = path
	URL.RawPath = path

	newEndpoint := httptransport.NewClient(
		"POST",
		URL,
		resetObjectTypesEncoder,
		resetObjectTypesDecoder,
	).Endpoint()

	return newEndpoint
}

func newResetVerbosityEndpoint(URL *url.URL, path string) endpoint.Endpoint {
	URL.Path = path
	URL.RawPath = path

	newEndpoint := httptransport.NewClient(
		"POST",
		URL,
		resetVerbosityEncoder,
		resetVerbosityDecoder,
	).Endpoint()

	return newEndpoint
}

func newSetLevelsEndpoint(URL *url.URL, path string) endpoint.Endpoint {
	URL.Path = path
	URL.RawPath = path

	newEndpoint := httptransport.NewClient(
		"POST",
		URL,
		setLevelsEncoder,
		setLevelsDecoder,
	).Endpoint()

	return newEndpoint
}

func newSetObjectTypesEndpoint(URL *url.URL, path string) endpoint.Endpoint {
	URL.Path = path
	URL.RawPath = path

	newEndpoint := httptransport.NewClient(
		"POST",
		URL,
		setObjectTypesEncoder,
		setObjectTypesDecoder,
	).Endpoint()

	return newEndpoint
}

func newSetVerbosityEndpoint(URL *url.URL, path string) endpoint.Endpoint {
	URL.Path = path
	URL.RawPath = path

	newEndpoint := httptransport.NewClient(
		"POST",
		URL,
		setVerbosityEncoder,
		setVerbosityDecoder,
	).Endpoint()

	return newEndpoint
}
