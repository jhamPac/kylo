package kylo

import (
	"context"
	"io"
)

// NewCancellableReader takes in rdr and returns an io.Reader
func NewCancellableReader(ctx context.Context, rdr io.Reader) io.Reader {
	return rdr
}
