package tokenizer

import (
	"github.com/koykov/byteseq"
	"github.com/koykov/nlp"
)

type Char[T byteseq.Byteseq] struct{}

func (t Char[T]) Tokenize(x T) nlp.Tokens {
	return t.AppendTokenize(nil, x)
}

func (t Char[T]) AppendTokenize(dst nlp.Tokens, x T) nlp.Tokens {
	for i := 0; i < len(x); i++ {
		dst = append(dst, nlp.ParseToken(x, i, i+1))
	}
	return dst
}
