package log

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/object/networkresponse"
)

// reset levels

func resetLevelsEncoder(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return maskAny(err)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func resetLevelsDecoder(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response networkresponse.ResetLevels
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, maskAny(err)
}

// reset object types

func resetObjectsEncoder(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return maskAny(err)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func resetObjectsDecoder(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response networkresponse.ResetObjects
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, maskAny(err)
}

// reset verbosity

func resetVerbosityEncoder(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return maskAny(err)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func resetVerbosityDecoder(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response networkresponse.ResetVerbosity
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, maskAny(err)
}

// set levels

func setLevelsEncoder(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return maskAny(err)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func setLevelsDecoder(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response networkresponse.SetLevels
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, maskAny(err)
}

// set object types

func setObjectsEncoder(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return maskAny(err)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func setObjectsDecoder(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response networkresponse.SetObjects
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, maskAny(err)
}

// set verbosity

func setVerbosityEncoder(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return maskAny(err)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func setVerbosityDecoder(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response networkresponse.SetVerbosity
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, maskAny(err)
}
