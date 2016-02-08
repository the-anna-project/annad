package logcontrol

import (
	"github.com/xh3b4sd/anna/api"
)

// reset levels

type ResetLevelsResponse api.Response

// reset object types

type ResetObjectTypesResponse api.Response

// reset verbosity

type ResetVerbosityResponse api.Response

// set levels

type SetLevelsRequest struct {
	Levels string `json:"levels,omitempty"`
}

type SetLevelsResponse api.Response

// set object types

type SetObjectTypesRequest struct {
	ObjectTypes string `json:"object_types,omitempty"`
}

type SetObjectTypesResponse api.Response

// set verbosity types

type SetVerbosityRequest struct {
	Verbosity int `json:"verbosity,omitempty"`
}

type SetVerbosityResponse api.Response
