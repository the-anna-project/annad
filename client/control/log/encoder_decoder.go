package logcontrol

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/xh3b4sd/anna/server/control/log"
)

// reset levels

func resetLevelsEncoder(r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return maskAny(err)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func resetLevelsDecoder(resp *http.Response) (interface{}, error) {
	var response logcontrol.ResetLevelsResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, maskAny(err)
}

// reset object types

func resetObjectTypesEncoder(r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return maskAny(err)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func resetObjectTypesDecoder(resp *http.Response) (interface{}, error) {
	var response logcontrol.ResetObjectTypesResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, maskAny(err)
}

// reset verbosity

func resetVerbosityEncoder(r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return maskAny(err)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func resetVerbosityDecoder(resp *http.Response) (interface{}, error) {
	var response logcontrol.ResetVerbosityResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, maskAny(err)
}

// set levels

func setLevelsEncoder(r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return maskAny(err)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func setLevelsDecoder(resp *http.Response) (interface{}, error) {
	var response logcontrol.SetLevelsResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, maskAny(err)
}

// set object types

func setObjectTypesEncoder(r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return maskAny(err)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func setObjectTypesDecoder(resp *http.Response) (interface{}, error) {
	var response logcontrol.SetObjectTypesResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, maskAny(err)
}

// set verbosity

func setVerbosityEncoder(r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return maskAny(err)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func setVerbosityDecoder(resp *http.Response) (interface{}, error) {
	var response logcontrol.SetVerbosityResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, maskAny(err)
}
