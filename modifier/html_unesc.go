package modifier

import (
	"github.com/koykov/byteseq"
	"github.com/koykov/entity/html"
	"github.com/koykov/fastconv"
)

type HTMLUnescaper[T byteseq.Byteseq] struct{}

func (m HTMLUnescaper[T]) Modify(x T) T {
	r := m.AppendModify(nil, x)
	_, s := fastconv.AppendR2S(nil, r)
	return T(s)
}

func (HTMLUnescaper[T]) AppendModify(dst []rune, x T) []rune {
	dst = html.AppendUnescapeRune(dst, x)
	return dst
}
