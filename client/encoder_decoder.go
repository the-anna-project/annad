package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/xh3b4sd/anna/server/interface/text"
)

func readPlainEncoder(r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return maskAny(err)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func readPlainDecoder(resp *http.Response) (interface{}, error) {
	var response textinterface.ReadPlainResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, maskAny(err)
}
