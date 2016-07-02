package api

// get response for ID

// GetResponseForIDRequest represents the request payload of the route used to
// fetch the response for a job by providing the job's ID. There must be an ID
// given.
type GetResponseForIDRequest struct {
	ID string `json:"id"`
}

// GetResponseForIDResponse represents the response payload of the route used
// to fetch the response for a job by providing the job's ID. This payload by
// convention follows the same schema as all other API responses.
type GetResponseForIDResponse Response

// read core request

// ReadCoreRequestRequest represents the request payload of the route used to
// read input by providing a core request. There must be a core request given.
type ReadCoreRequestRequest struct {
	CoreRequest CoreRequest `json:"core_request"`
	SessionID   string      `json:"session_id,omitempty"`
}

// ReadCoreRequestResponse represents the response payload of the route used to
// read input by providing a core request. This payload by convention follows
// the same schema as all other API responses.
type ReadCoreRequestResponse Response
