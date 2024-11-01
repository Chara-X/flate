package flate

type Compressor struct{}

func (c *Compressor) Write(b []byte) (n int, err error) {
	n = len(b)
	for len(b) > 0 {
		c.step()
		// b = b[c.fill(c, b):]
		// if c.err != nil {
		// 	return 0, c.err
		// }
	}
	return n, nil
}
func (d *Compressor) step() {}
