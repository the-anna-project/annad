package spec

import (
	"golang.org/x/net/context"
)

// TextInterface provides a way to feed neural networks with text input.
type TextInterface interface {
	// FetchURL causes Anna to fetch information from a given URL.
	FetchURL(url string) ([]byte, error)

	// ReadFile causes Anna to read information from the given file.
	ReadFile(file string) ([]byte, error)

	// ReadStream causes Anna to read information from the given stream. So far
	// the idea. The interface is wrong. The use case unclear.
	ReadStream(stream string) ([]byte, error)

	// ReadPlainWithID asks for a job result using the given job ID. This call
	// might block so long the job is not finished.
	ReadPlainWithID(ctx context.Context, jobID string) (string, error)

	// ReadPlainWithInput creates a new job to feed neural networks with input. A
	// new job ID will be returned immediately.
	//
	// Optionally there can be an expected result be given. In this case the
	// created job blocks as long as the neural networks generate output that
	// matches the provided expected one.
	//
	// Note that this only returns a job ID. The actual output can be read with
	// ReadPlainWithID.
	ReadPlainWithInput(ctx context.Context, input, expected string) (string, error)
}
