package networkresponse

// reset levels

// ResetLevels represents the response's payload of the route used to
// reset log levels. This payload by convention follows the same schema as all
// other API responses.
type ResetLevels Object

// reset object types

// ResetObjects represents the response's payload of the route used to
// reset log objects. This payload by convention follows the same schema as all
// other API responses.
type ResetObjects Object

// reset verbosity

// ResetVerbosity represents the response's payload of the route used to
// reset log verbosity. This payload by convention follows the same schema as
// all other API responses.
type ResetVerbosity Object

// set levels

// SetLevelsRequest represents the request payload of the route used to set log
// levels.
type SetLevelsRequest struct {
	Levels string `json:"levels,omitempty"`
}

// SetLevels represents the response's payload of the route used to set
// log levels. This payload by convention follows the same schema as all other
// API responses.
type SetLevels Object

// set object types

// SetObjectsRequest represents the request payload of the route used to set
// log objects.
type SetObjectsRequest struct {
	Objects string `json:"objects,omitempty"`
}

// SetObjects represents the response's payload of the route used to set
// log objects. This payload by convention follows the same schema as all
// other API responses.
type SetObjects Object

// set verbosity types

// SetVerbosityRequest represents the request payload of the route used to set
// log verbosity.
type SetVerbosityRequest struct {
	Verbosity int `json:"verbosity,omitempty"`
}

// SetVerbosity represents the response's payload of the route used to
// set log verbosity. This payload by convention follows the same schema as
// all other API responses.
type SetVerbosity Object
