package tokenizer

import (
	"github.com/koykov/byteseq"
	"github.com/koykov/nlp"
)

type Space[T byteseq.Byteseq] struct {
	o bool
	t nlp.StringTokenizer[T]
}

func (t Space[T]) Tokenize(x T) nlp.Tokens {
	return t.t.AppendTokenize(nil, x)
}

func (t Space[T]) AppendTokenize(dst nlp.Tokens, x T) nlp.Tokens {
	if !t.o {
		t.t.Separator = " "
		t.t.BlankLines = nlp.TokenizerBlankLinesKeep
		t.o = true
	}
	return t.t.AppendTokenize(dst, x)
}
