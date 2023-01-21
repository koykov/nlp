package nlp

import (
	"github.com/koykov/bytealg"
)

const DefaultTokenSeparator = " \n\t"

type TokenizerInterface[T Byteseq] interface {
	Tokenize(dst Tokens, x T) Tokens
}

type Tokenizer[T Byteseq] struct {
	sep string
	bl  bool
	eof bool
}

func NewTokenizer[T Byteseq]() Tokenizer[T] {
	return NewTokenizerWithOptions[T](DefaultTokenSeparator, true, false)
}

func NewTokenizerWithOptions[T Byteseq](sep string, keepBlank bool, discardEOF bool) Tokenizer[T] {
	return Tokenizer[T]{
		sep: sep,
		bl:  keepBlank,
		eof: discardEOF,
	}
}

func (t Tokenizer[T]) Tokenize(dst Tokens, x T) Tokens {
	s := q2s(x)
	lo, hi := 0, 0
	for {
		p := bytealg.IndexAnyAtStr(s, t.sep, lo)
		if p == -1 {
			hi = len(s)
			if hi == lo && t.eof {
				break
			}
			dst = append(dst, ParseToken(s, lo, hi))
			break
		}
		hi = p
		if hi == lo && !t.bl {
			lo = hi + 1
			continue
		}
		dst = append(dst, ParseToken(s, lo, hi))
		lo = hi + 1
	}
	return dst
}
