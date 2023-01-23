package tokenizer

import "github.com/koykov/nlp"

type Simple[T nlp.Byteseq] struct {
	nlp.Tokenizer[T]
}

func (t Simple[T]) Tokenize(dst nlp.Tokens, x T) nlp.Tokens {
	dst = t.Tokenizer.Tokenize(dst, x)
	return dst
}
