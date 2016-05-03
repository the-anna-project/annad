package logcontrol

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/server/control/log"
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
	var response logcontrol.ResetLevelsResponse
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
	var response logcontrol.ResetObjectsResponse
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
	var response logcontrol.ResetVerbosityResponse
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
	var response logcontrol.SetLevelsResponse
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
	var response logcontrol.SetObjectsResponse
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
	var response logcontrol.SetVerbosityResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, maskAny(err)
}
