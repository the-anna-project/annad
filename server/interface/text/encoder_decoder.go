package textinterface

import (
	"encoding/json"
	"net/http"
)

// fetch url

func fetchURLDecoder(r *http.Request) (interface{}, error) {
	var request FetchURLRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, maskAny(err)
	}
	return request, nil
}

func fetchURLEncoder(w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}

// read file

func readFileDecoder(r *http.Request) (interface{}, error) {
	var request ReadFileRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, maskAny(err)
	}
	return request, nil
}

func readFileEncoder(w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}

// read stream

func readStreamDecoder(r *http.Request) (interface{}, error) {
	var request ReadStreamRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, maskAny(err)
	}
	return request, nil
}

func readStreamEncoder(w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}

// read plain

func readPlainDecoder(r *http.Request) (interface{}, error) {
	var request ReadPlainRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, maskAny(err)
	}
	return request, nil
}

func readPlainEncoder(w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}
