package cleaner

import (
	"github.com/koykov/byteseq"
)

type Space[T byteseq.Byteseq] struct{}

func (c *Space[T]) Clean(x T) []rune {
	return c.AppendClean(nil, x)
}

func (c *Space[T]) AppendClean(dst []rune, x T) []rune {
	p := byteseq.Q2B(x)
	var ps bool
	for i := 0; i < len(p); i++ {
		if p[i] == ' ' {
			if !ps {
				dst = append(dst, rune(p[i]))
				ps = true
				continue
			}
		} else {
			ps = false
			dst = append(dst, rune(p[i]))
		}
	}

	l, r := 0, len(dst)-1
	for ; l < len(dst); l++ {
		if dst[l] != ' ' {
			break
		}
	}
	for ; r >= l; r-- {
		if dst[r] != ' ' {
			break
		}
	}

	return dst[l : r+1]
}
