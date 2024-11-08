package flate

import (
	"io"

	"github.com/Chara-X/binary"
)

type Deflator struct{ *binary.Writer }

func NewDeflator(w io.Writer) *Deflator { return &Deflator{&binary.Writer{Writer: w}} }
func (d *Deflator) WriteHeader()        { d.WriteBits(2, 3) }
func (d *Deflator) Write(b []byte) (n int, err error) {
	for i := 0; i < len(b); {
		var length, offset = 0, 0
		for j := 0; j < i; j++ {
			var l = 0
			for i+l < len(b) && b[i+l] == b[j+l] {
				l++
			}
			if l > length {
				length, offset = l, i-j
			}
		}
		if length > 3 {
			d.WriteBits(uint64(StdLengthEncoding[length-3].B), StdLengthEncoding[length-3].N)
			d.WriteBits(uint64(StdOffsetEncoding[offset-1].B), StdOffsetEncoding[offset-1].N)
			i += length
		} else {
			d.WriteBits(uint64(StdLiteralEncoding[b[i]].B), StdLiteralEncoding[b[i]].N)
			i++
		}
	}
	return len(b), nil
}
func (d *Deflator) WriteTailer() { d.WriteBits(0, 7) }
func (d *Deflator) Close()       { d.Writer.Close() }
