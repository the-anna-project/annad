package textinterface

import (
	"net/url"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/server/interface/text"
	serverspec "github.com/xh3b4sd/anna/server/spec"
)

type Config struct {
	URL *url.URL
}

func DefaultConfig() Config {
	newConfig := Config{
		URL: &url.URL{
			Host:   "127.0.0.1:9119",
			Scheme: "http",
		},
	}

	return newConfig
}

func NewTextInterface(config Config) serverspec.TextInterface {
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
		if s, ok := apiResponse.Data.(string); ok {
			return "", maskAnyf(invalidAPIResponseError, s)
		}
	}

	if api.WithData("").Code != apiResponse.Code {
		return "", maskAny(invalidAPIResponseError)
	}

	if s, ok := apiResponse.Data.(string); ok {
		return s, nil
	} else {
		return "", maskAny(invalidAPIResponseError)
	}
}

func (ti textInterface) ReadPlainWithPlain(ctx context.Context, plain string) (string, error) {
	response, err := ti.readPlainWithPlain(ctx, textinterface.ReadPlainRequest{Plain: plain})
	if err != nil {
		return "", maskAnyf(invalidAPIResponseError, err.Error())
	}

	apiResponse := response.(textinterface.ReadPlainResponse)

	if api.WithError(nil).Code == apiResponse.Code {
		if s, ok := apiResponse.Data.(string); ok {
			return "", maskAnyf(invalidAPIResponseError, s)
		}
	}

	if api.WithID("").Code != apiResponse.Code {
		return "", maskAny(invalidAPIResponseError)
	}

	if s, ok := apiResponse.Data.(string); ok {
		return s, nil
	} else {
		return "", maskAny(invalidAPIResponseError)
	}
}
