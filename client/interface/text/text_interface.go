// Package textinterface provides functionality to interact with Anna's text
// interface network API.
package textinterface

import (
	"net/url"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/server/interface/text"
	"github.com/xh3b4sd/anna/spec"
)

// Config represents the configuration used to create a new text interface
// object.
type Config struct {
	// URL represents the API route to call.
	URL *url.URL
}

// DefaultConfig provides a default configuration to create a new text
// interface object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		URL: &url.URL{
			Host:   "127.0.0.1:9119",
			Scheme: "http",
		},
	}

	return newConfig
}

// NewTextInterface creates a new configured text interface object.
func NewTextInterface(config Config) spec.TextInterface {
	newTextInterface := &textInterface{
		Config: config,

		readPlainWithID:    newReadPlainWithIDEndpoint(*config.URL, "/interface/text/action/readplain"),
		readPlainWithPlain: newReadPlainWithPlainEndpoint(*config.URL, "/interface/text/action/readplain"),
	}

	return newTextInterface
}

type textInterface struct {
	Config

	readPlainWithID    endpoint.Endpoint
	readPlainWithPlain endpoint.Endpoint
}

func (ti textInterface) FetchURL(url string) ([]byte, error) {
	// TODO
	return nil, nil
}

func (ti textInterface) ReadFile(file string) ([]byte, error) {
	// TODO
	return nil, nil
}

func (ti textInterface) ReadStream(stream string) ([]byte, error) {
	// TODO
	return nil, nil
}

func (ti textInterface) ReadPlainWithID(ctx context.Context, ID string) (string, error) {
	response, err := ti.readPlainWithID(ctx, textinterface.ReadPlainRequest{ID: ID})
	if err != nil {
		return "", maskAnyf(invalidAPIResponseError, err.Error())
	}

	apiResponse := response.(textinterface.ReadPlainResponse)

	if api.WithError(nil).Code == apiResponse.Code {
		switch t := apiResponse.Data.(type) {
		case string:
			return "", maskAnyf(invalidAPIResponseError, t)
		}
	}

	if api.WithData("").Code == apiResponse.Code {
		switch t := apiResponse.Data.(type) {
		case string:
			// We received the expected response.
			return t, nil
		}
	}

	return "", maskAnyf(invalidAPIResponseError, "unexpected API response")
}

func (ti textInterface) ReadPlainWithInput(ctx context.Context, input, expected, sessionID string) (string, error) {
	response, err := ti.readPlainWithPlain(ctx, textinterface.ReadPlainRequest{Input: input, Expected: expected, SessionID: sessionID})
	if err != nil {
		return "", maskAnyf(invalidAPIResponseError, err.Error())
	}

	apiResponse := response.(textinterface.ReadPlainResponse)

	if api.WithError(nil).Code == apiResponse.Code {
		switch t := apiResponse.Data.(type) {
		case string:
			return "", maskAnyf(invalidAPIResponseError, t)
		}
	}

	if api.WithID("").Code == apiResponse.Code {
		switch t := apiResponse.Data.(type) {
		case string:
			// We received the expected response.
			return t, nil
		}
	}

	return "", maskAnyf(invalidAPIResponseError, "unexpected API response")
}
