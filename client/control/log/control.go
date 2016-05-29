package log

import (
	"net/url"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/server/control/log"
	"github.com/xh3b4sd/anna/spec"
)

// ControlConfig represents the configuration used to create a new log control
// object.
type ControlConfig struct {
	// URL represents the API route to call.
	URL *url.URL
}

// DefaultControlConfig provides a default configuration to create a new log
// control object by best effort.
func DefaultControlConfig() ControlConfig {
	newConfig := ControlConfig{
		URL: &url.URL{
			Host:   "127.0.0.1:9119",
			Scheme: "http",
		},
	}

	return newConfig
}

// NewControl creates a new configured log control object.
func NewControl(config ControlConfig) (spec.LogControl, error) {
	newControl := &control{
		ControlConfig: config,

		resetLevels:    newResetLevelsEndpoint(*config.URL, "/control/log/reset/levels"),
		resetObjects:   newResetObjectsEndpoint(*config.URL, "/control/log/reset/objects"),
		resetVerbosity: newResetVerbosityEndpoint(*config.URL, "/control/log/reset/verbosity"),
		setLevels:      newSetLevelsEndpoint(*config.URL, "/control/log/set/levels"),
		setObjects:     newSetObjectsEndpoint(*config.URL, "/control/log/set/objects"),
		setVerbosity:   newSetVerbosityEndpoint(*config.URL, "/control/log/set/verbosity"),
	}

	return newControl, nil
}

type control struct {
	ControlConfig

	resetLevels    endpoint.Endpoint
	resetObjects   endpoint.Endpoint
	resetVerbosity endpoint.Endpoint
	setLevels      endpoint.Endpoint
	setObjects     endpoint.Endpoint
	setVerbosity   endpoint.Endpoint
}

func (c *control) ResetLevels(ctx context.Context) error {
	response, err := c.resetLevels(ctx, nil)
	if err != nil {
		return maskAnyf(invalidAPIResponseError, err.Error())
	}

	apiResponse := response.(log.ResetLevelsResponse)

	if api.WithError(nil).Code == apiResponse.Code {
		switch t := apiResponse.Data.(type) {
		case string:
			return maskAnyf(invalidAPIResponseError, t)
		}
	}

	if api.WithSuccess().Code == apiResponse.Code {
		switch apiResponse.Data.(type) {
		case string:
			// We received the expected response.
			return nil
		}
	}

	// TODO proper logging
	return maskAnyf(invalidAPIResponseError, "unexpected API response")
}

func (c *control) ResetObjects(ctx context.Context) error {
	response, err := c.resetObjects(ctx, nil)
	if err != nil {
		return maskAnyf(invalidAPIResponseError, err.Error())
	}

	apiResponse := response.(log.ResetObjectsResponse)

	if api.WithError(nil).Code == apiResponse.Code {
		switch t := apiResponse.Data.(type) {
		case string:
			return maskAnyf(invalidAPIResponseError, t)
		}
	}

	if api.WithSuccess().Code == apiResponse.Code {
		switch apiResponse.Data.(type) {
		case string:
			// We received the expected response.
			return nil
		}
	}

	// TODO proper logging
	return maskAnyf(invalidAPIResponseError, "unexpected API response")
}

func (c *control) ResetVerbosity(ctx context.Context) error {
	response, err := c.resetVerbosity(ctx, nil)
	if err != nil {
		return maskAnyf(invalidAPIResponseError, err.Error())
	}

	apiResponse := response.(log.ResetVerbosityResponse)

	if api.WithError(nil).Code == apiResponse.Code {
		switch t := apiResponse.Data.(type) {
		case string:
			return maskAnyf(invalidAPIResponseError, t)
		}
	}

	if api.WithSuccess().Code == apiResponse.Code {
		switch apiResponse.Data.(type) {
		case string:
			// We received the expected response.
			return nil
		}
	}

	// TODO proper logging
	return maskAnyf(invalidAPIResponseError, "unexpected API response")
}

func (c *control) SetLevels(ctx context.Context, levels string) error {
	response, err := c.setLevels(ctx, log.SetLevelsRequest{Levels: levels})
	if err != nil {
		return maskAnyf(invalidAPIResponseError, err.Error())
	}

	apiResponse := response.(log.SetLevelsResponse)

	if api.WithError(nil).Code == apiResponse.Code {
		switch t := apiResponse.Data.(type) {
		case string:
			return maskAnyf(invalidAPIResponseError, t)
		}
	}

	if api.WithSuccess().Code == apiResponse.Code {
		switch apiResponse.Data.(type) {
		case string:
			// We received the expected response.
			return nil
		}
	}

	// TODO proper logging
	return maskAnyf(invalidAPIResponseError, "unexpected API response")
}

func (c *control) SetObjects(ctx context.Context, objects string) error {
	response, err := c.setObjects(ctx, log.SetObjectsRequest{Objects: objects})
	if err != nil {
		return maskAnyf(invalidAPIResponseError, err.Error())
	}

	apiResponse := response.(log.SetObjectsResponse)

	if api.WithError(nil).Code == apiResponse.Code {
		switch t := apiResponse.Data.(type) {
		case string:
			return maskAnyf(invalidAPIResponseError, t)
		}
	}

	if api.WithSuccess().Code == apiResponse.Code {
		switch apiResponse.Data.(type) {
		case string:
			// We received the expected response.
			return nil
		}
	}

	// TODO proper logging
	return maskAnyf(invalidAPIResponseError, "unexpected API response")
}

func (c *control) SetVerbosity(ctx context.Context, verbosity int) error {
	response, err := c.setVerbosity(ctx, log.SetVerbosityRequest{Verbosity: verbosity})
	if err != nil {
		return maskAnyf(invalidAPIResponseError, err.Error())
	}

	apiResponse := response.(log.SetVerbosityResponse)

	if api.WithError(nil).Code == apiResponse.Code {
		switch t := apiResponse.Data.(type) {
		case string:
			return maskAnyf(invalidAPIResponseError, t)
		}
	}

	if api.WithSuccess().Code == apiResponse.Code {
		switch apiResponse.Data.(type) {
		case string:
			// We received the expected response.
			return nil
		}
	}

	// TODO proper logging
	return maskAnyf(invalidAPIResponseError, "unexpected API response")
}
