package log

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

func resetLevelsEndpoint(lc spec.LogControl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := lc.ResetLevels(ctx)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return api.WithSuccess(), nil
	}
}

func resetObjectsEndpoint(lc spec.LogControl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := lc.ResetObjects(ctx)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return api.WithSuccess(), nil
	}
}

func resetVerbosityEndpoint(lc spec.LogControl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := lc.ResetVerbosity(ctx)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return api.WithSuccess(), nil
	}
}

func setLevelsEndpoint(lc spec.LogControl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.SetLevelsRequest)

		err := lc.SetLevels(ctx, req.Levels)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return api.WithSuccess(), nil
	}
}

func setObjectsEndpoint(lc spec.LogControl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.SetObjectsRequest)

		err := lc.SetObjects(ctx, req.Objects)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return api.WithSuccess(), nil
	}
}

func setVerbosityEndpoint(lc spec.LogControl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.SetVerbosityRequest)

		err := lc.SetVerbosity(ctx, req.Verbosity)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return api.WithSuccess(), nil
	}
}
