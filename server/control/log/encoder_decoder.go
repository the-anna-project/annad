package log

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
)

// reset levels

func resetLevelsDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func resetLevelsEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}

// reset object types

func resetObjectsDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func resetObjectsEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}

// reset verbosity

func resetVerbosityDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func resetVerbosityEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}

// set levels

func setLevelsDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var request api.SetLevelsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, maskAny(err)
	}
	return request, nil
}

func setLevelsEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}

// set object types

func setObjectsDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var request api.SetObjectsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, maskAny(err)
	}
	return request, nil
}

func setObjectsEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}

// set verbosity

func setVerbosityDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var request api.SetVerbosityRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, maskAny(err)
	}
	return request, nil
}

func setVerbosityEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}
