package text

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
)

// get response for ID

func getResponseForIDEncoder(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return maskAny(err)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func getResponseForIDDecoder(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response api.GetResponseForIDResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, maskAny(err)
}

// read core request

func readCoreRequestEncoder(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return maskAny(err)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func readCoreRequestDecoder(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response api.ReadCoreRequestResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, maskAny(err)
}
