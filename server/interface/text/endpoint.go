package textinterface

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

func fetchURLEndpoint(ti spec.TextInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FetchURLRequest)

		response, err := ti.FetchURL(req.URL)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return api.WithData(string(response)), nil
	}
}

func readFileEndpoint(ti spec.TextInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadFileRequest)

		response, err := ti.ReadFile(req.File)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return api.WithData(string(response)), nil
	}
}

func readStreamEndpoint(ti spec.TextInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadStreamRequest)

		response, err := ti.ReadStream(req.Stream)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return api.WithData(string(response)), nil
	}
}

func readPlainEndpoint(ti spec.TextInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadPlainRequest)

		var err error
		var ID string
		var response string

		if req.ID == "" && req.Plain == "" {
			// All empty means error.
			return api.WithError(maskAny(invalidRequestError)), nil
		}

		if req.ID != "" && req.Plain == "" {
			// Only ID given means there is something we want to fetch by ID.
			response, err = ti.ReadPlainWithID(ctx, req.ID)
			if err != nil {
				return api.WithError(maskAny(err)), nil
			}
			return api.WithData(response), nil
		}

		if req.ID == "" && req.Plain != "" {
			// Only Plain given means we want to do something, but only return an ID
			// in the first place.
			ID, err = ti.ReadPlainWithPlain(ctx, req.Plain)
			if err != nil {
				return api.WithError(maskAny(err)), nil
			}
			return api.WithID(ID), nil
		}

		if req.ID != "" && req.Plain != "" {
			// All NOT empty means error.
			return api.WithError(maskAny(invalidRequestError)), nil
		}

		return api.WithError(maskAny(invalidRequestError)), nil
	}
}
