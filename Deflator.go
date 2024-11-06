package flate

import "fmt"

type Deflator struct{ *BitWriter }

func (d *Deflator) WriteHeader() { d.WriteBits(2, 3) }
func (d *Deflator) Write(b []byte) (n int, err error) {
	// for _, v := range b {
	// 	d.WriteBits(uint64(literalEncoding[v].code), literalEncoding[v].len)
	// }
	// return len(b), nil
	for i := 0; i < len(b); {
		var length, offset = 0, 0
		for j := 0; j < i; j++ {
			var l = 0
			for i+l < len(b) && b[i+l] == b[j+l] {
				l++
			}
			if l > length {
				offset = i - j
				length = l
			}
		}
		if length > 3 {
			fmt.Println("i:", i, "length:", length, "offset:", offset)
			fmt.Println("StdLengthEncoding[length-3].Code:", StdLengthEncoding[length-3].Code)
			fmt.Println("StdLengthEncoding[length-3].N:", StdLengthEncoding[length-3].N)
			fmt.Println("StdOffsetEncoding[offset-1].Code:", StdOffsetEncoding[offset-1].Code)
			fmt.Println("StdOffsetEncoding[offset-1].N:", StdOffsetEncoding[offset-1].N)
			d.WriteBits(uint64(StdLengthEncoding[length-3].Code), StdLengthEncoding[length-3].N)
			d.WriteBits(uint64(StdOffsetEncoding[offset-1].Code), StdOffsetEncoding[offset-1].N)
			i += length
		} else {
			fmt.Println("i:", i)
			d.WriteBits(uint64(StdLiteralEncoding[b[i]].Code), StdLiteralEncoding[b[i]].N)
			i++
		}
	}
	return len(b), nil
}
func (d *Deflator) WriteTailer() { d.WriteBits(0, 7) }
func (d *Deflator) Close()       { d.BitWriter.Close() }
