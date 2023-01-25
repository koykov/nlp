package tokenizer

import (
	"github.com/koykov/byteseq"
	"github.com/koykov/nlp"
)

type Tab[T byteseq.Byteseq] struct {
	o bool
	t nlp.StringTokenizer[T]
}

func (t Tab[T]) Tokenize(x T) nlp.Tokens {
	return t.t.AppendTokenize(nil, x)
}

func (t Tab[T]) AppendTokenize(dst nlp.Tokens, x T) nlp.Tokens {
	if !t.o {
		t.t.Separator = "\t"
		t.t.BlankLines = nlp.TokenizerBlankLinesKeep
		t.o = true
	}
	return t.t.AppendTokenize(dst, x)
}
