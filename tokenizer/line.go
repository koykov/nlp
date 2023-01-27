package tokenizer

import (
	"github.com/koykov/byteseq"
	"github.com/koykov/nlp"
)

type Line[T byteseq.Byteseq] struct {
	BlankLines nlp.TokenizerBlankLines

	o bool
	t nlp.StringTokenizer[T]
}

func (t Line[T]) Tokenize(x T) nlp.Tokens {
	return t.t.AppendTokenize(nil, x)
}

func (t Line[T]) AppendTokenize(dst nlp.Tokens, x T) nlp.Tokens {
	if !t.o {
		t.t.Separator = "\n"
		t.t.BlankLines = t.BlankLines
		t.o = true
	}
	return t.t.AppendTokenize(dst, x)
}
