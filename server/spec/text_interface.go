package spec

import (
	"golang.org/x/net/context"
)

type TextInterface interface {
	FetchURL(url string) ([]byte, error)
	ReadFile(file string) ([]byte, error)
	ReadStream(stream string) ([]byte, error)
	ReadPlainWithID(ctx context.Context, ID string) (string, error)
	ReadPlainWithPlain(ctx context.Context, plain string) (string, error)
}
