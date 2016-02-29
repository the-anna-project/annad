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

	// ReadPlainWithID causes Anna to read information a job associated with the
	// given ID provides.
	ReadPlainWithID(ctx context.Context, ID string) (string, error)

	// ReadPlainWithPlain causes Anna to read information as provided by plain.
	// Note that this returns a job ID. The actual response to a request created
	// with ReadPlainWithPlain can be read with ReadPlainWithID.
	ReadPlainWithPlain(ctx context.Context, plain string) (string, error)
}
