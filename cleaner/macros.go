package cleaner

import (
	"github.com/koykov/bytealg"
	"github.com/koykov/byteseq"
	"github.com/koykov/fastconv"
)

type Macros[T byteseq.Byteseq] struct {
	Left, Right string
}

func (c Macros[T]) Clean(x T) []rune {
	return c.AppendClean(nil, x)
}

func (c Macros[T]) AppendClean(dst []rune, x T) []rune {
	l, r := c.Left, c.Right
	s := byteseq.Q2S(x)
	if len(l) == 0 || len(r) == 0 {
		dst = fastconv.AppendS2R(dst, s)
		return dst
	}
	off := 0
	for {
		p := bytealg.IndexAtStr(s, l, off)
		if p == -1 {
			dst = fastconv.AppendS2R(dst, s[off:])
			return dst
		}
		dst = fastconv.AppendS2R(dst, s[:p])
		off = p + len(l)
		if p = bytealg.IndexAtStr(s, r, off); p == -1 {
			dst = fastconv.AppendS2R(dst, s[off:])
			return dst
		}
		off = p + len(r)
	}
}
