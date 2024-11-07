package flate

import "math/bits"

type PrefixCode struct{ B, N int }

func NewPrefixCode(highB, highN, lowB, lowN int) PrefixCode {
	return PrefixCode{B: int(bits.Reverse64(uint64(highB<<(64-highN)))) | lowB&((1<<lowN)-1)<<highN, N: highN + lowN}
}
