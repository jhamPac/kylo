package kylo

import "io"

// NewCancellableReader takes in rdr and returns an io.Reader
func NewCancellableReader(rdr io.Reader) io.Reader {
	return rdr
}
