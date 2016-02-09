package logcontrol

import (
	"encoding/json"
	"net/http"
)

// reset levels

func resetLevelsDecoder(r *http.Request) (interface{}, error) {
	return nil, nil
}

func resetLevelsEncoder(w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}

// reset object types

func resetObjectsDecoder(r *http.Request) (interface{}, error) {
	return nil, nil
}

func resetObjectsEncoder(w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}

// reset verbosity

func resetVerbosityDecoder(r *http.Request) (interface{}, error) {
	return nil, nil
}

func resetVerbosityEncoder(w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}

// set levels

func setLevelsDecoder(r *http.Request) (interface{}, error) {
	var request SetLevelsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, maskAny(err)
	}
	return request, nil
}

func setLevelsEncoder(w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}

// set object types

func setObjectsDecoder(r *http.Request) (interface{}, error) {
	var request SetObjectsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, maskAny(err)
	}
	return request, nil
}

func setObjectsEncoder(w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}

// set verbosity

func setVerbosityDecoder(r *http.Request) (interface{}, error) {
	var request SetVerbosityRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, maskAny(err)
	}
	return request, nil
}

func setVerbosityEncoder(w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}
