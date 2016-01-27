package client

import (
	"net/url"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/server/interface/text"
	serverspec "github.com/xh3b4sd/anna/server/spec"
)

type TextInterfaceConfig struct {
	URL *url.URL
}

func DefaultTextInterfaceConfig() TextInterfaceConfig {
	newTextInterfaceConfig := TextInterfaceConfig{
		URL: &url.URL{
			Host:   "127.0.0.1:9119",
			Scheme: "http",
		},
	}

	return newTextInterfaceConfig
}

func NewTextInterface(config TextInterfaceConfig) serverspec.TextInterface {
	newTextInterface := &textInterface{
		TextInterfaceConfig: config,

		readPlainWithID:    newReadPlainWithIDEndpoint(config.URL, "/interface/text/action/readplain"),
		readPlainWithPlain: newReadPlainWithPlainEndpoint(config.URL, "/interface/text/action/readplain"),
	}

	return newTextInterface
}

type textInterface struct {
	TextInterfaceConfig

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
		return "", maskAnyWIthCause(err, invalidAPIResponseError)
	}

	apiResponse := response.(textinterface.ReadPlainResponse)
	if api.WithData("").Code != apiResponse.Code {
		return "", maskAny(invalidAPIResponseError)
	}

	if data, ok := apiResponse.Data.(string); ok {
		return data, nil
	} else {
		return "", maskAny(invalidAPIResponseError)
	}
}

func (ti textInterface) ReadPlainWithPlain(ctx context.Context, plain string) (string, error) {
	response, err := ti.readPlainWithPlain(ctx, textinterface.ReadPlainRequest{Plain: plain})
	if err != nil {
		return "", maskAnyWIthCause(err, invalidAPIResponseError)
	}

	apiResponse := response.(textinterface.ReadPlainResponse)
	if api.WithID("").Code != apiResponse.Code {
		return "", maskAny(invalidAPIResponseError)
	}

	if s, ok := apiResponse.Data.(string); ok {
		return s, nil
	} else {
		return "", maskAny(invalidAPIResponseError)
	}
}
