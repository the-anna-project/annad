package textinterface

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
)

func fetchURLEndpoint(ti TextInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(fetchURLRequest)

		err := ti.FetchURL(req.URL)
		if err != nil {
			return nil, maskAny(err)
		}

		return api.Success(), nil
	}
}

func readFileEndpoint(ti TextInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(readFileRequest)

		err := ti.ReadFile(req.File)
		if err != nil {
			return nil, maskAny(err)
		}

		return api.Success(), nil
	}
}

func readStreamEndpoint(ti TextInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(readStreamRequest)

		err := ti.ReadStream(req.Stream)
		if err != nil {
			return nil, maskAny(err)
		}

		return api.Success(), nil
	}
}

func readPlainEndpoint(ti TextInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(readPlainRequest)

		err := ti.ReadPlain(req.Plain)
		if err != nil {
			return nil, maskAny(err)
		}

		return api.Success(), nil
	}
}
