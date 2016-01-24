package textinterface

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
)

func fetchURLEndpoint(ti TextInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(fetchURLRequest)

		response, err := ti.FetchURL(req.URL)
		if err != nil {
			return nil, maskAny(err)
		}

		return api.WithData(response), nil
	}
}

func readFileEndpoint(ti TextInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(readFileRequest)

		response, err := ti.ReadFile(req.File)
		if err != nil {
			return nil, maskAny(err)
		}

		return api.WithData(response), nil
	}
}

func readStreamEndpoint(ti TextInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(readStreamRequest)

		response, err := ti.ReadStream(req.Stream)
		if err != nil {
			return nil, maskAny(err)
		}

		return api.WithData(response), nil
	}
}

func readPlainEndpoint(ti TextInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(readPlainRequest)

		var err error
		var ID string
		var response []byte

		if req.ID == "" && req.Plain == "" {
			// All empty means error.
			return nil, maskAny(invalidRequestError)
		}

		if req.ID != "" && req.Plain == "" {
			// Only ID given means there is something we want to fetch by ID.
			response, err = ti.ReadPlainWithID(ctx, req.ID)
			if err != nil {
				return nil, maskAny(err)
			}
			return api.WithData(string(response)), nil
		}

		if req.ID == "" && req.Plain != "" {
			// Only Plain given means we want to do something, but only return an ID
			// in the first place.
			ID, err = ti.ReadPlainWithPlain(ctx, []byte(req.Plain))
			if err != nil {
				return nil, maskAny(err)
			}
			return api.WithID(ID), nil
		}

		if req.ID != "" && req.Plain != "" {
			// All NOT empty means error.
			return nil, maskAny(invalidRequestError)
		}

		return nil, maskAny(invalidRequestError)
	}
}
