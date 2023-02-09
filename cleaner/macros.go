package cleaner

import "github.com/koykov/byteseq"

type Macros[T byteseq.Byteseq] struct{}

func (c Macros[T]) Clean(x T) []rune {
	return c.AppendClean(nil, x)
}

func (Macros[T]) AppendClean(dst []rune, x T) []rune {
	_ = x
	return dst
}
