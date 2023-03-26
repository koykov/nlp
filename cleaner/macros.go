package cleaner

import (
	"github.com/koykov/bytealg"
	"github.com/koykov/byteseq"
	"github.com/koykov/fastconv"
)

type Macros[T byteseq.Byteseq] struct {
	Left, Right string

	o    bool
	l, r string
}

func (c *Macros[T]) Clean(x T) []rune {
	return c.AppendClean(nil, x)
}

func (c *Macros[T]) AppendClean(dst []rune, x T) []rune {
	if !c.o {
		c.o = true
		c.l, c.r = c.Left, c.Right
	}
	s := byteseq.Q2S(x)
	off := 0
	for {
		p := bytealg.IndexAtStr(s, c.l, off)
		if p == -1 {
			dst = fastconv.AppendS2R(dst, s[off:])
			return dst
		}
		dst = fastconv.AppendS2R(dst, s[:p-1])
		off = p + len(c.l)
		if p = bytealg.IndexAtStr(s, c.r, off); p == -1 {
			dst = fastconv.AppendS2R(dst, s[off:])
			return dst
		}
		off = p + len(c.r)
	}
}
