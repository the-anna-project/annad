package textinterface

import (
	"github.com/xh3b4sd/anna/api"
)

// fetch url

type FetchURLRequest struct {
	ID  string `json:"id,omitempty"`
	URL string `json:"url"`
}

type FetchURLResponse api.Response

// read file

type ReadFileRequest struct {
	File string `json:"file"`
	ID   string `json:"id,omitempty"`
}

type ReadFileResponse api.Response

// read stream

type ReadStreamRequest struct {
	ID     string `json:"id,omitempty"`
	Stream string `json:"stream"`
}

type ReadStreamResponse api.Response

// read plain

type ReadPlainRequest struct {
	ID    string `json:"id,omitempty"`
	Plain string `json:"plain,omitempty"`
}

type ReadPlainResponse api.Response
