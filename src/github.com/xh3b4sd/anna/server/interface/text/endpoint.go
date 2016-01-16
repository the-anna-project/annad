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

		response, err := ti.ReadPlain([]byte(req.Plain))
		if err != nil {
			return nil, maskAny(err)
		}

		return api.WithData(string(response)), nil
	}
}
