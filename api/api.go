// Package api implements structures and helpers for network responses. The
// server packages makes use of this to provide a consistent API response
// scheme.
package api

var (
	CodeData, TextData       = "10001", "data"
	CodeSuccess, TextSuccess = "10002", "success"
	CodeError, TextError     = "10003", "error"
)

// Response is the response type each API call should return.
type Response struct {
	Code string      `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Text string      `json:"text,omitempty"`
}

// WithData returns a response having the given data set as Data. Text 'data'
// translates to the Code 10001.
func WithData(data interface{}) Response {
	return Response{
		Code: CodeData,
		Data: data,
		Text: TextData,
	}
}

// WithSuccess returns a response indicating the success of the requested
// action. Text 'success' translates to the Code 10002.
func WithSuccess() Response {
	return Response{
		Code: CodeSuccess,
		Data: "",
		Text: TextSuccess,
	}
}

// WithError returns a response indicating an error of the requested action.
// Text 'error' translates to the Code 10003.
func WithError(err error) Response {
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	return Response{
		Code: CodeError,
		Data: msg,
		Text: TextError,
	}
}
