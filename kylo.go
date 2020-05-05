package kylo

import (
	"context"
	"io"
)

type readerCtx struct {
	ctx      context.Context
	delegate io.Reader
}

func (r *readerCtx) Read(p []byte) (n int, err error) {
	if err := r.ctx.Err(); err != nil {
		return 0, err
	}
	return r.delegate.Read(p)
}

// NewCancellableReader takes in rdr and returns an io.Reader
func NewCancellableReader(ctx context.Context, rdr io.Reader) io.Reader {
	return &readerCtx{
		ctx:      ctx,
		delegate: rdr,
	}
}
