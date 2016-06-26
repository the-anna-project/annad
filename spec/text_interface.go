package spec

import (
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
)

// TextInterface provides a way to feed neural networks with text input.
type TextInterface interface {
	// GetResponseForID asks for a job result using the given job ID. This call
	// might block so long the job is not finished.
	GetResponseForID(ctx context.Context, jobID string) (string, error)

	// ReadCoreRequest creates a new job to feed the neural network with input.
	// Only a new job ID will be returned, immediately. The actual job response
	// can be retrieved using GetResponseForID.
	ReadCoreRequest(ctx context.Context, coreRequest api.CoreRequest, sessionID string) (string, error)
}
