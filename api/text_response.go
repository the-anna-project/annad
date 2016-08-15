package api

import (
	"encoding/json"

	"github.com/xh3b4sd/anna/spec"
)

// TextResponseConfig represents the configuration used to create a new text
// response object.
type TextResponseConfig struct {
	// Settings.

	// Output represents the output being calculated by the neural network.
	Output string `json:"output"`
}

// DefaultTextResponseConfig provides a default configuration to create a new
// text response object by best effort.
func DefaultTextResponseConfig() TextResponseConfig {
	newConfig := TextResponseConfig{
		Output: "",
	}

	return newConfig
}

// NewTextResponse creates a new configured text response object.
func NewTextResponse(config TextResponseConfig) (spec.TextResponse, error) {
	newTextResponse := &textResponse{
		TextResponseConfig: config,
	}

	return newTextResponse, nil
}

// NewEmptyTextResponse simply returns an empty, maybe invalid, text response
// object. This should only be used for things like unmarshaling.
func NewEmptyTextResponse() spec.TextResponse {
	return &textResponse{}
}

type textResponse struct {
	TextResponseConfig
}

func (tr *textResponse) GetOutput() string {
	return tr.Output
}

// textResponseClone is for making use of the stdlib json implementation. The
// textResponse object implements its own marshaler and unmarshaler but only to
// provide json implementations for spec.TextResponse. Note, not redirecting
// the type will cause infinite recursion.
type textResponseClone textResponse

func (tr *textResponse) MarshalJSON() ([]byte, error) {
	newTextResponse := textResponseClone(*tr)

	raw, err := json.Marshal(newTextResponse)
	if err != nil {
		return nil, maskAny(err)
	}

	return raw, nil
}

func (tr *textResponse) UnmarshalJSON(b []byte) error {
	newTextResponse := textResponseClone{}

	err := json.Unmarshal(b, &newTextResponse)
	if err != nil {
		return maskAny(err)
	}

	tr.Output = newTextResponse.Output

	return nil
}
