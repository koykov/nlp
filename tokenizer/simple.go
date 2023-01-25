package tokenizer

import "github.com/koykov/nlp"

type Simple
[T nlp.byteseq.Byteseq] struct {
nlp.StringTokenizer[T]
}

func NewSimple[T nlp.byteseq.Byteseq]() Simple[T] {
	return NewSimpleWithOptions[T](nlp.DefaultTokenSeparator, true, false)
}

func NewSimpleWithOptions[T nlp.byteseq.Byteseq](sep string, keepBlank bool, discardEOF bool) Simple[T] {
	tkn := nlp.NewStringTokenizer[T](sep, keepBlank, discardEOF)
	return Simple[T]{tkn}
}

func (t Simple[T]) Tokenize(dst nlp.Tokens, x T) nlp.Tokens {
	dst = t.StringTokenizer.Tokenize(dst, x)
	return dst
}
