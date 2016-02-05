package logcontrol

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	serverspec "github.com/xh3b4sd/anna/server/spec"
)

func resetLevelsEndpoint(lc serverspec.LogControl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := lc.ResetLevels(ctx)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return api.WithSuccess(), nil
	}
}

func resetObjectsEndpoint(lc serverspec.LogControl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := lc.ResetObjects(ctx)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return api.WithSuccess(), nil
	}
}

func resetVerbosityEndpoint(lc serverspec.LogControl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := lc.ResetVerbosity(ctx)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return api.WithSuccess(), nil
	}
}

func setLevelsEndpoint(lc serverspec.LogControl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetLevelsRequest)

		err := lc.SetLevels(ctx, req.Levels)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return api.WithSuccess(), nil
	}
}

func setObjectsEndpoint(lc serverspec.LogControl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetObjectsRequest)

		err := lc.SetObjects(ctx, req.Objects)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return api.WithSuccess(), nil
	}
}

func setVerbosityEndpoint(lc serverspec.LogControl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetVerbosityRequest)

		err := lc.SetVerbosity(ctx, req.Verbosity)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return api.WithSuccess(), nil
	}
}
