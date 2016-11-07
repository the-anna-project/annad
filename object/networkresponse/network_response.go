// Package networkresponse implements structures and helpers for network
// responses. The server packages makes use of this to provide a consistent API
// response scheme.
package networkresponse

var (
	// CodeData represents the API response code of a data response.
	CodeData = "10001"

	// TextData represents the API response text of a data response.
	TextData = "data"

	// CodeSuccess represents the API response code of a success response.
	CodeSuccess = "10002"

	// TextSuccess represents the API response text of a success response.
	TextSuccess = "success"

	// CodeError represents the API response code of a error response.
	CodeError = "10003"

	// TextError represents the API response text of a error response.
	TextError = "error"
)

// Object is the response type each API call should return.
type Object struct {
	Code string      `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Text string      `json:"text,omitempty"`
}

// WithData returns a response having the given data set as Data. Text 'data'
// translates to the Code 10001.
func WithData(data interface{}) Object {
	return Object{
		Code: CodeData,
		Data: data,
		Text: TextData,
	}
}

// WithSuccess returns a response indicating the success of the requested
// action. Text 'success' translates to the Code 10002.
func WithSuccess() Object {
	return Object{
		Code: CodeSuccess,
		Data: "",
		Text: TextSuccess,
	}
}

// WithError returns a response indicating an error of the requested action.
// Text 'error' translates to the Code 10003.
func WithError(err error) Object {
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	return Object{
		Code: CodeError,
		Data: msg,
		Text: TextError,
	}
}
