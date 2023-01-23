package tokenizer

import "github.com/koykov/nlp"

type Simple[T nlp.Byteseq] struct {
	nlp.Tokenizer[T]
}

func NewSimple[T nlp.Byteseq]() Simple[T] {
	return NewSimpleWithOptions[T](nlp.DefaultTokenSeparator, true, false)
}

func NewSimpleWithOptions[T nlp.Byteseq](sep string, keepBlank bool, discardEOF bool) Simple[T] {
	tkn := nlp.NewTokenizerWithOptions[T](sep, keepBlank, discardEOF)
	return Simple[T]{tkn}
}

func (t Simple[T]) Tokenize(dst nlp.Tokens, x T) nlp.Tokens {
	dst = t.Tokenizer.Tokenize(dst, x)
	return dst
}
