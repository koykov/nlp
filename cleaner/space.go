package cleaner

import (
	"github.com/koykov/byteseq"
)

type Space[T byteseq.Byteseq] struct{}

func (c Space[T]) Clean(x T) []rune {
	return c.AppendClean(nil, x)
}

func (c Space[T]) AppendClean(dst []rune, x T) []rune {
	s := byteseq.Q2S(x)
	var ps bool
	for _, c1 := range s {
		if c1 == ' ' {
			if !ps {
				dst = append(dst, c1)
				ps = true
				continue
			}
		} else {
			ps = false
			dst = append(dst, c1)
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
