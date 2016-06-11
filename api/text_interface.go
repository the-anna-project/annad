package api

// fetch url

// FetchURLRequest represents the request payload of the route used to fetch
// URLs.
type FetchURLRequest struct {
	ID  string `json:"id,omitempty"`
	URL string `json:"url"`
}

// FetchURLResponse represents the response's payload of the route used to fetch
// URLs. This payload by convention follows the same schema as all other API
// responses.
type FetchURLResponse Response

// read file

// ReadFileRequest represents the request payload of the route used to read
// files.
type ReadFileRequest struct {
	File string `json:"file"`
	ID   string `json:"id,omitempty"`
}

// ReadFileResponse represents the response's payload of the route used to read
// files. This payload by convention follows the same schema as all other API
// responses.
type ReadFileResponse Response

// read stream

// ReadStreamRequest represents the request payload of the route used to read
// streams.
type ReadStreamRequest struct {
	ID     string `json:"id,omitempty"`
	Stream string `json:"stream"`
}

// ReadStreamResponse represents the response's payload of the route used to
// read streams. This payload by convention follows the same schema as all
// other API responses.
type ReadStreamResponse Response

// read plain

// ReadPlainRequest represents the request payload of the route used to read
// plain input.
type ReadPlainRequest struct {
	ID        string `json:"id,omitempty"`
	Input     string `json:"input,omitempty"`
	Expected  string `json:"expected,omitempty"`
	SessionID string `json:"session_id,omitempty"`
}

// ReadPlainResponse represents the response's payload of the route used to read
// plain input. This payload by convention follows the same schema as all other
// API responses.
type ReadPlainResponse Response
