package text

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

func streamTextEndpoint(ti spec.TextInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.StreamTextRequest)

		if req.CoreRequest.IsEmpty() {
			// There must be an core request given. We don't have one, thus we return
			// an error.
			return api.WithError(maskAny(invalidRequestError)), nil
		}

		in := make(chan api.NetworkRequest, 1)
		out := make(chan api.NetworkResponse, 1000)

		go func() {
			// TODO stream continously
			in <- req.NetworkRequest
		}()

		go func() {
			// TODO stream continously
			api.WithData(<-out)
		}()

		err := ti.StreamText(ctx, in, out)
		if err != nil {
			return api.WithError(maskAny(err)), nil
		}

		return nil, nil
	}
}
