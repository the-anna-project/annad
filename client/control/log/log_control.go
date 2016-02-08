package logcontrol

import (
	"net/url"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/server/interface/text"
	serverspec "github.com/xh3b4sd/anna/server/spec"
)

type LogControlConfig struct {
	URL *url.URL
}

func DefaultLogConfig() LogControlConfig {
	newLogControlConfig := LogControlConfig{
		URL: &url.URL{
			Host:   "127.0.0.1:9119",
			Scheme: "http",
		},
	}

	return newLogControlConfig
}

func NewLogControl(config LogControlConfig) serverspec.LogControl {
	newLogControl := &logControl{
		LogControlConfig: config,

		resetLevels:      newResetLevelsEndpoint(config.URL, "/control/log/set/levels"),
		resetObjectTypes: newResetObjectTypesEndpoint(config.URL, "/control/log/set/objecttypes"),
		resetVerbosity:   newResetVerbosityEndpoint(config.URL, "/control/log/set/verbosity"),
		setLevels:        newSetLevelsEndpoint(config.URL, "/control/log/reset/levels"),
		setObjectTypes:   newSetObjectTypesEndpoint(config.URL, "/control/log/reset/objecttypes"),
		setVerbosity:     newSetVerbosityEndpoint(config.URL, "/control/log/reset/verbosity"),
	}

	return newLogControl
}

type logControl struct {
	LogControlConfig

	resetLevels      endpoint.Endpoint
	resetObjectTypes endpoint.Endpoint
	resetVerbosity   endpoint.Endpoint
	setLevels        endpoint.Endpoint
	setObjectTypes   endpoint.Endpoint
	setVerbosity     endpoint.Endpoint
}

func (l logControl) ResetLevels(ctx context.Context) error {
	response, err := l.resetLevels(ctx)
	if err != nil {
		return maskAnyWithCause(err, invalidAPIResponseError)
	}

	apiResponse := response.(logcontrol.ResetLevelsResponse)
	if api.WithSuccess().Code != apiResponse.Code {
		return maskAny(invalidAPIResponseError)
	}

	return nil
}

func (l logControl) ResetObjectTypes(ctx context.Context) error {
	response, err := l.resetObjectTypes(ctx)
	if err != nil {
		return maskAnyWithCause(err, invalidAPIResponseError)
	}

	apiResponse := response.(logcontrol.ResetObjectTypesResponse)
	if api.WithSuccess().Code != apiResponse.Code {
		return maskAny(invalidAPIResponseError)
	}

	return nil
}

func (l logControl) ResetVerbosity(ctx context.Context) error {
	response, err := l.resetVerbosity(ctx)
	if err != nil {
		return maskAnyWithCause(err, invalidAPIResponseError)
	}

	apiResponse := response.(logcontrol.ResetVerbosityResponse)
	if api.WithSuccess().Code != apiResponse.Code {
		return maskAny(invalidAPIResponseError)
	}

	return nil
}

func (l logControl) SetLevels(ctx context.Context, levels string) error {
	response, err := l.setLevels(ctx, logcontrol.SetLevelsRequest{Levels: levels})
	if err != nil {
		return maskAnyWithCause(err, invalidAPIResponseError)
	}

	apiResponse := response.(logcontrol.SetLevelsResponse)
	if api.WithSuccess().Code != apiResponse.Code {
		return maskAny(invalidAPIResponseError)
	}

	return nil
}

func (l logControl) SetObjectTypes(ctx context.Context, objectTypes string) error {
	response, err := l.setObjectTypes(ctx, logcontrol.SetObjectTypesRequest{ObjectTypes: objectTypes})
	if err != nil {
		return maskAnyWithCause(err, invalidAPIResponseError)
	}

	apiResponse := response.(logcontrol.SetObjectTypesResponse)
	if api.WithSuccess().Code != apiResponse.Code {
		return maskAny(invalidAPIResponseError)
	}

	return nil
}

func (l logControl) SetVerbosity(ctx context.Context, verbosity int) error {
	response, err := l.setVerbosity(ctx, logcontrol.SetVerbosityRequest{Verbosity: verbosity})
	if err != nil {
		return maskAnyWithCause(err, invalidAPIResponseError)
	}

	apiResponse := response.(logcontrol.SetVerbosityResponse)
	if api.WithSuccess().Code != apiResponse.Code {
		return maskAny(invalidAPIResponseError)
	}

	return nil
}
