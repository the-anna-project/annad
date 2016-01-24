package textinterface

import (
	"github.com/xh3b4sd/anna/api"
)

// fetch url

type fetchURLRequest struct {
	ID  string `json:"id,omitempty"`
	URL string `json:"url"`
}

type fetchURLResponse api.Response

// read file

type readFileRequest struct {
	File string `json:"file"`
	ID   string `json:"id,omitempty"`
}

type readFileResponse api.Response

// read stream

type readStreamRequest struct {
	ID     string `json:"id,omitempty"`
	Stream string `json:"stream"`
}

type readStreamResponse api.Response

// read plain

type readPlainRequest struct {
	ID    string `json:"id,omitempty"`
	Plain string `json:"plain,omitempty"`
}

type readPlainResponse api.Response
