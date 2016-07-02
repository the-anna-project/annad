package text

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

func getResponseForIDEndpoint(ti spec.TextInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.GetResponseForIDRequest)

		if req.ID == "" {
			// There must be an ID given. We don't have one, thus we return an error.
			return api.WithError(maskAny(invalidRequestError)), nil
		}

		response, err := ti.GetResponseForID(ctx, req.ID)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return api.WithData(response), nil
	}
}

func readCoreRequestEndpoint(ti spec.TextInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.ReadCoreRequestRequest)

		if req.CoreRequest.IsEmpty() {
			// There must be an core request given. We don't have one, thus we return
			// an error.
			return api.WithError(maskAny(invalidRequestError)), nil
		}

		response, err := ti.ReadCoreRequest(ctx, req.CoreRequest, req.SessionID)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return api.WithData(response), nil
	}
}
