package api

// reset levels

// ResetLevelsResponse represents the response's payload of the route used to
// reset log levels. This payload by convention follows the same schema as all
// other API responses.
type ResetLevelsResponse Response

// reset object types

// ResetObjectsResponse represents the response's payload of the route used to
// reset log objects. This payload by convention follows the same schema as all
// other API responses.
type ResetObjectsResponse Response

// reset verbosity

// ResetVerbosityResponse represents the response's payload of the route used to
// reset log verbosity. This payload by convention follows the same schema as
// all other API responses.
type ResetVerbosityResponse Response

// set levels

// SetLevelsRequest represents the request payload of the route used to set log
// levels.
type SetLevelsRequest struct {
	Levels string `json:"levels,omitempty"`
}

// SetLevelsResponse represents the response's payload of the route used to set
// log levels. This payload by convention follows the same schema as all other
// API responses.
type SetLevelsResponse Response

// set object types

// SetObjectsRequest represents the request payload of the route used to set
// log objects.
type SetObjectsRequest struct {
	Objects string `json:"objects,omitempty"`
}

// SetObjectsResponse represents the response's payload of the route used to set
// log objects. This payload by convention follows the same schema as all
// other API responses.
type SetObjectsResponse Response

// set verbosity types

// SetVerbosityRequest represents the request payload of the route used to set
// log verbosity.
type SetVerbosityRequest struct {
	Verbosity int `json:"verbosity,omitempty"`
}

// SetVerbosityResponse represents the response's payload of the route used to
// set log verbosity. This payload by convention follows the same schema as
// all other API responses.
type SetVerbosityResponse Response
