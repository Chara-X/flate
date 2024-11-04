package flate

var codeOrder = []int{16, 17, 18, 0, 8, 7, 9, 6, 10, 5, 11, 4, 12, 3, 13, 2, 14, 1, 15}

type Decompressor struct {
	*BitReader
}

func (d *Decompressor) ReadBlock() []byte {
	var final = d.ReadBits(1)
	_ = final
	var blockType = d.ReadBits(2)
	if blockType != 2 {
		panic("not dynamic huffman")
	}
	d.ReadBits(10)
	var nCLen = d.ReadBits(4) + 4
	var codeLens = make([]byte, 19)
	for i := 0; i < int(nCLen); i++ {
		codeLens[codeOrder[i]] = byte(d.ReadBits(3))
	}
	return nil
}
