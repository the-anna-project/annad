package textinterface

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
)

// fetch url

func fetchURLDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var request FetchURLRequest
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
	var request ReadFileRequest
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
	var request ReadStreamRequest
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
	var request ReadPlainRequest
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
