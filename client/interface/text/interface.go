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

		streamText: newStreamTextEndpoint(*config.URL, "/interface/text"),
	}

	return newInterface, nil
}

type tinterface struct {
	InterfaceConfig

	streamText endpoint.Endpoint
}

func (i tinterface) StreamText(ctx context.Context, in chan api.TextRequest, out chan api.TextResponse) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case textRequest := <-in:
			response, err := i.streamText(ctx, api.StreamTextRequest{TextRequest: textRequest})
			if err != nil {
				return maskAnyf(invalidAPIResponseError, err.Error())
			}

			apiResponse := response.(api.StreamTextResponse)

			if api.WithError(nil).Code == apiResponse.Code {
				if errMessage, ok := apiResponse.Data.(string); ok {
					return maskAnyf(invalidAPIResponseError, errMessage)
				}
			}

			if api.WithID("").Code == apiResponse.Code {
				if output, ok := apiResponse.Data.(string); ok {
					out <- api.TextResponse{Output: output}
				}
			}

			return maskAnyf(invalidAPIResponseError, "unexpected API response")
		}
	}
}
