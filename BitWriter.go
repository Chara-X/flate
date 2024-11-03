package flate

import "io"

type BitWriter struct {
	io.Writer
}

func (bw *BitWriter) WriteBits(b uint64, n int) {
	panic("not implemented")
}
