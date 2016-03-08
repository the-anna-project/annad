package logcontrol

import (
	"github.com/xh3b4sd/anna/api"
)

// reset levels

// ResetLevelsResponse represents the response's payload of the route used to
// reset log levels. This payload by convention follows the same schema as all
// other API responses.
type ResetLevelsResponse api.Response

// reset object types

// ResetObjectsResponse represents the response's payload of the route used to
// reset log objects. This payload by convention follows the same schema as all
// other API responses.
type ResetObjectsResponse api.Response

// reset verbosity

// ResetVerbosityResponse represents the response's payload of the route used to
// reset log verbosity. This payload by convention follows the same schema as
// all other API responses.
type ResetVerbosityResponse api.Response

// set levels

// SetLevelsRequest represents the request payload of the route used to set log
// levels.
type SetLevelsRequest struct {
	Levels string `json:"levels,omitempty"`
}

// SetLevelsResponse represents the response's payload of the route used to set
// log levels. This payload by convention follows the same schema as all other
// API responses.
type SetLevelsResponse api.Response

// set object types

// SetObjectsRequest represents the request payload of the route used to set
// log objects.
type SetObjectsRequest struct {
	Objects string `json:"objects,omitempty"`
}

// SetObjectsResponse represents the response's payload of the route used to set
// log objects. This payload by convention follows the same schema as all
// other API responses.
type SetObjectsResponse api.Response

// set verbosity types

// SetVerbosityRequest represents the request payload of the route used to set
// log verbosity.
type SetVerbosityRequest struct {
	Verbosity int `json:"verbosity,omitempty"`
}

// SetVerbosityResponse represents the response's payload of the route used to
// set log verbosity. This payload by convention follows the same schema as
// all other API responses.
type SetVerbosityResponse api.Response
