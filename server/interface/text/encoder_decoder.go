package text

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
)

// fetch url

func fetchURLDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var request api.FetchURLRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, maskAny(err)
	}
	return request, nil
}

func fetchURLEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}

// read file

func readFileDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var request api.ReadFileRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, maskAny(err)
	}
	return request, nil
}

func readFileEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}

// read stream

func readStreamDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var request api.ReadStreamRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, maskAny(err)
	}
	return request, nil
}

func readStreamEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}

// read plain

func readPlainDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var request api.ReadPlainRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, maskAny(err)
	}
	return request, nil
}

func readPlainEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}
