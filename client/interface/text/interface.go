package text

import (
	"net/url"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

// InterfaceConfig represents the configuration used to create a new text interface
// object.
type InterfaceConfig struct {
	// URL represents the API route to call.
	URL *url.URL
}

// DefaultInterfaceConfig provides a default configuration to create a new text
// interface object by best effort.
func DefaultInterfaceConfig() InterfaceConfig {
	newConfig := InterfaceConfig{
		URL: &url.URL{
			Host:   "127.0.0.1:9119",
			Scheme: "http",
		},
	}

	return newConfig
}

// NewInterface creates a new configured text interface object.
func NewInterface(config InterfaceConfig) (spec.TextInterface, error) {
	newInterface := &tinterface{
		InterfaceConfig: config,

		readPlainWithID:    newReadPlainWithIDEndpoint(*config.URL, "/interface/text/action/readplain"),
		readPlainWithPlain: newReadPlainWithPlainEndpoint(*config.URL, "/interface/text/action/readplain"),
	}

	return newInterface, nil
}

type tinterface struct {
	InterfaceConfig

	readPlainWithID    endpoint.Endpoint
	readPlainWithPlain endpoint.Endpoint
}

func (i tinterface) FetchURL(url string) ([]byte, error) {
	// TODO
	return nil, nil
}

func (i tinterface) ReadFile(file string) ([]byte, error) {
	// TODO
	return nil, nil
}

func (i tinterface) ReadStream(stream string) ([]byte, error) {
	// TODO
	return nil, nil
}

func (i tinterface) ReadPlainWithID(ctx context.Context, ID string) (string, error) {
	response, err := i.readPlainWithID(ctx, api.ReadPlainRequest{ID: ID})
	if err != nil {
		return "", maskAnyf(invalidAPIResponseError, err.Error())
	}

	apiResponse := response.(api.ReadPlainResponse)

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

func (i tinterface) ReadPlainWithInput(ctx context.Context, input string, expectation api.ExpectationRequest, sessionID string) (string, error) {
	response, err := i.readPlainWithPlain(ctx, api.ReadPlainRequest{Input: input, Expectation: expectation, SessionID: sessionID})
	if err != nil {
		return "", maskAnyf(invalidAPIResponseError, err.Error())
	}

	apiResponse := response.(api.ReadPlainResponse)

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
