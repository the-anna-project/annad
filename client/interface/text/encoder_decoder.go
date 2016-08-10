package text

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
)

// stream text

func streamTextEncoder(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return maskAny(err)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func streamTextDecoder(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response api.StreamTextResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, maskAny(err)
}
