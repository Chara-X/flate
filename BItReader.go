package flate

import "io"

type BitReader struct {
	io.Reader
	bits  uint64
	nbits uint
}

func (r *BitReader) ReadBits(n int) uint64 {
	panic("not implemented")
}
func (r *BitReader) ReadByte() byte {
	panic("not implemented")
}
