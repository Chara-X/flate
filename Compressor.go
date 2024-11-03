package flate

type Compressor struct {
	w *BitWriter
}

func (c *Compressor) Write(b []byte) (n int, err error) {
	var freqs = make([]int, 256)
	for _, v := range b {
		freqs[v]++
	}
	freqs[256] = 1
	return 0, nil
}
