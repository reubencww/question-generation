package storage

import (
	"context"
	"io"
)

type Storage interface {
	Upload(ctx context.Context, filename, mime string, file io.Reader) error
	GetURL(ctx context.Context, filename string) (string, error)
}
