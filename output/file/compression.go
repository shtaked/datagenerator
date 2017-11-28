package file

import (
	"compress/gzip"
	"io"
)

//Compressor interface returns io.WriteCloser object
type Compressor interface {
	NewWriter(w io.Writer) io.WriteCloser
}

type GzipCompressor struct{}

func (fn GzipCompressor) NewWriter(w io.Writer) io.WriteCloser {
	return gzip.NewWriter(w)
}
