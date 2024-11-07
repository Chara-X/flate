package flate

var StdLiteralEncoding, StdLengthEncoding, StdOffsetEncoding = [256]PrefixCode{}, [256]PrefixCode{}, [32768]PrefixCode{}

func init() {
	for c := 0; c < 256; c++ {
		var code = []int{}
		switch {
		case c < 144:
			code = []int{c + 48, 8}
		default:
			code = []int{c + 256, 9}
		}
		StdLiteralEncoding[c] = NewPrefixCode(code[0], code[1], 0, 0)
	}
	for c := 0; c < 256; c++ {
		var code = []int{}
		switch {
		case c < 8:
			code = []int{c + 1, 7, 0, 0}
		case c < 10:
			code = []int{9, 7, c - 8, 1}
		case c < 12:
			code = []int{10, 7, c - 10, 1}
		case c < 14:
			code = []int{11, 7, c - 12, 1}
		case c < 16:
			code = []int{12, 7, c - 14, 1}
		case c < 20:
			code = []int{13, 7, c - 16, 3}
		case c < 24:
			code = []int{14, 7, c - 20, 3}
		case c < 28:
			code = []int{15, 7, c - 24, 3}
		case c < 32:
			code = []int{16, 7, c - 28, 3}
		case c < 40:
			code = []int{17, 7, c - 32, 7}
		case c < 48:
			code = []int{18, 7, c - 40, 7}
		case c < 56:
			code = []int{19, 7, c - 48, 7}
		case c < 64:
			code = []int{20, 7, c - 56, 7}
		case c < 80:
			code = []int{21, 7, c - 64, 15}
		case c < 96:
			code = []int{22, 7, c - 80, 15}
		case c < 112:
			code = []int{23, 7, c - 96, 15}
		case c < 128:
			code = []int{24, 8, c - 112, 15}
		case c < 160:
			code = []int{25, 8, c - 128, 31}
		case c < 192:
			code = []int{26, 8, c - 160, 31}
		case c < 224:
			code = []int{27, 8, c - 192, 31}
		case c < 255:
			code = []int{28, 8, c - 224, 31}
		case c < 256:
			code = []int{29, 8, 0, 0}
		}
		StdLengthEncoding[c] = NewPrefixCode(code[0], code[1], code[2], code[3])
	}
	for c := 0; c < 32768; c++ {
		var code = []int{}
		switch {
		case c < 4:
			code = []int{c, 5, 0, 0}
		case c < 6:
			code = []int{4, 5, c - 4, 1}
		case c < 8:
			code = []int{5, 5, c - 6, 1}
		case c < 12:
			code = []int{6, 5, c - 8, 2}
		case c < 16:
			code = []int{7, 5, c - 12, 2}
		case c < 24:
			code = []int{8, 5, c - 16, 3}
		case c < 32:
			code = []int{9, 5, c - 24, 3}
		case c < 48:
			code = []int{10, 5, c - 32, 4}
		case c < 64:
			code = []int{11, 5, c - 48, 4}
		case c < 96:
			code = []int{12, 5, c - 64, 5}
		case c < 128:
			code = []int{13, 5, c - 96, 5}
		case c < 192:
			code = []int{14, 5, c - 128, 6}
		case c < 256:
			code = []int{15, 5, c - 192, 6}
		case c < 384:
			code = []int{16, 5, c - 256, 7}
		case c < 512:
			code = []int{17, 5, c - 384, 7}
		case c < 768:
			code = []int{18, 5, c - 512, 8}
		case c < 1024:
			code = []int{19, 5, c - 768, 8}
		case c < 1536:
			code = []int{20, 5, c - 1024, 9}
		case c < 2048:
			code = []int{21, 5, c - 1536, 9}
		case c < 3072:
			code = []int{22, 5, c - 2048, 10}
		case c < 4096:
			code = []int{23, 5, c - 3072, 10}
		case c < 6144:
			code = []int{24, 5, c - 4096, 11}
		case c < 8192:
			code = []int{25, 5, c - 6144, 11}
		case c < 12288:
			code = []int{26, 5, c - 8192, 12}
		case c < 16384:
			code = []int{27, 5, c - 12288, 12}
		case c < 24576:
			code = []int{28, 5, c - 16384, 13}
		case c < 32768:
			code = []int{29, 5, c - 24576, 13}
		}
		StdOffsetEncoding[c] = NewPrefixCode(code[0], code[1], code[2], code[3])
	}
}
