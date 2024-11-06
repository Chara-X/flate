package flate

import "io"

type BitReader struct {
	io.Reader
}

func (r *BitReader) ReadBits(n int) uint64 {
	panic("not implemented")
}
func (r *BitReader) ReadByte() byte {
	panic("not implemented")
}
