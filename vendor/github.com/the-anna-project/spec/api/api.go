// Package api provides information to communicate with the network API. Client
// and server implementations make use of this to provide consistent API
// schemes.
package api

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
