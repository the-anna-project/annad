// package api implements structures and helpers for network responses. The
// server packages makes use of this to provide a consistent API response
// scheme.
package api

// Response is the response type each API call should return.
type Response struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Text string      `json:"text,omitempty"`
}

// WithID returns a response having the given ID set as Data. Text 'id'
// translates to the Code 10001.
func WithID(ID string) Response {
	return Response{
		Code: 10001,
		Data: ID,
		Text: "id",
	}
}

// WithData returns a response having the given data set as Data. Text 'data'
// translates to the Code 10002.
func WithData(data string) Response {
	return Response{
		Code: 10002,
		Data: data,
		Text: "data",
	}
}
