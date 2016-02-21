package logcontrol

import (
	"net/url"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/server/control/log"
	"github.com/xh3b4sd/anna/spec"
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

func NewLogControl(config Config) spec.LogControl {
	newLogControl := &logControl{
		Config: config,

		resetLevels:    newResetLevelsEndpoint(*config.URL, "/control/log/reset/levels"),
		resetObjects:   newResetObjectsEndpoint(*config.URL, "/control/log/reset/objects"),
		resetVerbosity: newResetVerbosityEndpoint(*config.URL, "/control/log/reset/verbosity"),
		setLevels:      newSetLevelsEndpoint(*config.URL, "/control/log/set/levels"),
		setObjects:     newSetObjectsEndpoint(*config.URL, "/control/log/set/objects"),
		setVerbosity:   newSetVerbosityEndpoint(*config.URL, "/control/log/set/verbosity"),
	}

	return newLogControl
}

type logControl struct {
	Config

	resetLevels    endpoint.Endpoint
	resetObjects   endpoint.Endpoint
	resetVerbosity endpoint.Endpoint
	setLevels      endpoint.Endpoint
	setObjects     endpoint.Endpoint
	setVerbosity   endpoint.Endpoint
}

func (l logControl) ResetLevels(ctx context.Context) error {
	response, err := l.resetLevels(ctx, nil)
	if err != nil {
		return maskAnyf(invalidAPIResponseError, err.Error())
	}

	apiResponse := response.(logcontrol.ResetLevelsResponse)

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

func (l logControl) ResetObjects(ctx context.Context) error {
	response, err := l.resetObjects(ctx, nil)
	if err != nil {
		return maskAnyf(invalidAPIResponseError, err.Error())
	}

	apiResponse := response.(logcontrol.ResetObjectsResponse)

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

func (l logControl) ResetVerbosity(ctx context.Context) error {
	response, err := l.resetVerbosity(ctx, nil)
	if err != nil {
		return maskAnyf(invalidAPIResponseError, err.Error())
	}

	apiResponse := response.(logcontrol.ResetVerbosityResponse)

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

func (l logControl) SetLevels(ctx context.Context, levels string) error {
	response, err := l.setLevels(ctx, logcontrol.SetLevelsRequest{Levels: levels})
	if err != nil {
		return maskAnyf(invalidAPIResponseError, err.Error())
	}

	apiResponse := response.(logcontrol.SetLevelsResponse)

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

func (l logControl) SetObjects(ctx context.Context, objects string) error {
	response, err := l.setObjects(ctx, logcontrol.SetObjectsRequest{Objects: objects})
	if err != nil {
		return maskAnyf(invalidAPIResponseError, err.Error())
	}

	apiResponse := response.(logcontrol.SetObjectsResponse)

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

func (l logControl) SetVerbosity(ctx context.Context, verbosity int) error {
	response, err := l.setVerbosity(ctx, logcontrol.SetVerbosityRequest{Verbosity: verbosity})
	if err != nil {
		return maskAnyf(invalidAPIResponseError, err.Error())
	}

	apiResponse := response.(logcontrol.SetVerbosityResponse)

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
