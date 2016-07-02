package text

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
)

// get response for ID

func getResponseForIDDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var request api.GetResponseForIDRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, maskAny(err)
	}
	return request, nil
}

func getResponseForIDEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}

// read core request

func readCoreRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var request api.ReadCoreRequestRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, maskAny(err)
	}
	return request, nil
}

func readCoreRequestEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}
