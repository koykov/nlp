package nlp

import "github.com/koykov/byteseq"

type DummyModifier[T byteseq.Byteseq] struct{}

func (DummyModifier[T]) Modify(x T) T { return x }
func (DummyModifier[T]) AppendModify(dst []rune, _ T) []rune {
	return dst
}
