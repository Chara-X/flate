package flate

import "math/bits"

var StdLiteralEncoding = [256]PrefixCode{}
var StdLengthEncoding = [256]PrefixCode{}
var StdOffsetEncoding = [32768]PrefixCode{}

func init() {
	for c := 0; c < 256; c++ {
		var code, n = 0, 0
		switch {
		case c < 144:
			code = c + 48
			n = 8
		case c < 256:
			code = c + 256
			n = 9
		}
		StdLiteralEncoding[c] = PrefixCode{Code: reverseBits(code, n), N: n}
	}
	for c := 0; c < 256; c++ {
		var code, n = 0, 0
		switch {
		case c < 8:
			code = c + 1<<0
			n = 7
		case c < 10:
			code = c + 9<<1
			n = 8
		case c < 12:
			code = c + 10<<1
			n = 8
		case c < 14:
			code = c + 11<<1
			n = 8
		case c < 16:
			code = c + 12<<1
			n = 8
		case c < 20:
			code = c + 13<<2
			n = 9
		case c < 24:
			code = c + 14<<2
			n = 9
		case c < 28:
			code = c + 15<<2
			n = 9
		case c < 32:
			code = c + 16<<2
			n = 9
		case c < 40:
			code = c + 17<<3
			n = 10
		case c < 48:
			code = c + 18<<3
			n = 10
		case c < 56:
			code = c + 19<<3
			n = 10
		case c < 64:
			code = c + 20<<3
			n = 10
		case c < 80:
			code = c + 21<<4
			n = 11
		case c < 96:
			code = c + 22<<4
			n = 11
		case c < 112:
			code = c + 23<<4
			n = 11
		case c < 128:
			code = c + 192<<4
			n = 12
		case c < 160:
			code = c + 193<<5
			n = 13
		case c < 192:
			code = c + 194<<5
			n = 13
		case c < 224:
			code = c + 195<<5
			n = 13
		case c < 255:
			code = c + 196<<5
			n = 13
		case c < 256:
			code = c + 197<<0
			n = 8
		}
		StdLengthEncoding[c] = PrefixCode{Code: reverseBits(code, n), N: n}
	}
	for c := 0; c < 32768; c++ {
		var code, n = 0, 0
		switch {
		case c < 1:
			code = c + 0<<0
			n = 5
		case c < 2:
			code = c + 1<<0
			n = 5
		case c < 3:
			code = c + 2<<0
			n = 5
		case c < 4:
			code = c + 3<<0
			n = 5
		case c < 6:
			code = c + 4<<1
			n = 6
		case c < 8:
			code = c + 5<<1
			n = 6
		case c < 12:
			code = c + 6<<2
			n = 7
		case c < 16:
			code = c + 7<<2
			n = 7
		case c < 24:
			code = c + 8<<3
			n = 8
		case c < 32:
			code = c + 9<<3
			n = 8
		case c < 48:
			code = c + 10<<4
			n = 9
		case c < 64:
			code = c + 11<<4
			n = 9
		case c < 96:
			code = c + 12<<5
			n = 10
		case c < 128:
			code = c + 13<<5
			n = 10
		case c < 192:
			code = c + 14<<6
			n = 11
		case c < 256:
			code = c + 15<<6
			n = 11
		case c < 384:
			code = c + 16<<7
			n = 12
		case c < 512:
			code = c + 17<<7
			n = 12
		case c < 768:
			code = c + 18<<8
			n = 13
		case c < 1024:
			code = c + 19<<8
			n = 13
		case c < 1536:
			code = c + 20<<9
			n = 14
		case c < 2048:
			code = c + 21<<9
			n = 14
		case c < 3072:
			code = c + 22<<10
			n = 15
		case c < 4096:
			code = c + 23<<10
			n = 15
		case c < 6144:
			code = c + 24<<11
			n = 16
		case c < 8192:
			code = c + 25<<11
			n = 16
		case c < 12288:
			code = c + 26<<12
			n = 17
		case c < 16384:
			code = c + 27<<12
			n = 17
		case c < 24576:
			code = c + 28<<13
			n = 18
		case c < 32768:
			code = c + 29<<13
			n = 18
		}
		StdOffsetEncoding[c] = PrefixCode{Code: reverseBits(code, n), N: n}
	}
}
func reverseBits(x int, n int) int { return int(bits.Reverse64(uint64(x << (64 - n)))) }
