package textinterface

import (
	"github.com/xh3b4sd/anna/api"
)

// fetch url

type fetchURLRequest struct {
	URL string `json:"url"`
}

type fetchURLResponse api.Response

// read file

type readFileRequest struct {
	File string `json:"file"`
}

type readFileResponse api.Response

// read stream

type readStreamRequest struct {
	Stream string `json:"stream"`
}

type readStreamResponse api.Response

// read plain

type readPlainRequest struct {
	Plain string `json:"plain"`
}

type readPlainResponse api.Response
