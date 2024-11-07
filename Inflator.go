package flate

import (
	"io"

	"github.com/Chara-X/binary"
)

type Inflator struct{ *binary.Reader }

func NewInflator(r io.Reader) *Inflator { return &Inflator{&binary.Reader{Reader: r}} }
