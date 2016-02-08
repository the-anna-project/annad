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
			return nil, maskAny(err)
		}

		return api.WithSuccess(), nil
	}
}

func resetObjectTypesEndpoint(lc serverspec.LogControl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := lc.ResetObjectTypes(ctx)
		if err != nil {
			return nil, maskAny(err)
		}

		return api.WithSuccess(), nil
	}
}

func resetVerbosityEndpoint(lc serverspec.LogControl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := lc.ResetVerbosity(ctx)
		if err != nil {
			return nil, maskAny(err)
		}

		return api.WithSuccess(), nil
	}
}

func setLevelsEndpoint(lc serverspec.LogControl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetLevelsRequest)

		err := lc.SetLevels(ctx, req.Levels)
		if err != nil {
			return nil, maskAny(err)
		}

		return api.WithSuccess(), nil
	}
}

func setObjectTypesEndpoint(lc serverspec.LogControl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetObjectTypesRequest)

		err := lc.SetObjectTypes(ctx, req.ObjectTypes)
		if err != nil {
			return nil, maskAny(err)
		}

		return api.WithSuccess(), nil
	}
}

func setVerbosityEndpoint(lc serverspec.LogControl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetVerbosityRequest)

		err := lc.SetVerbosity(ctx, req.Verbosity)
		if err != nil {
			return nil, maskAny(err)
		}

		return api.WithSuccess(), nil
	}
}
