package nlp

import (
	"github.com/koykov/bytealg"
	"github.com/koykov/fastconv"
)

const DefaultTokenSeparator = " \n\t"

type TokenizerInterface interface {
	Tokenize(dst Tokens, p []byte) Tokens
	TokenizeString(dst Tokens, s string) Tokens
}

type Tokenizer struct {
	sep string
	bl  bool
	eof bool
}

func NewTokenizer() Tokenizer {
	return NewTokenizerWithOptions(DefaultTokenSeparator, true, false)
}

func NewTokenizerWithOptions(sep string, keepBlank bool, discardEOF bool) Tokenizer {
	return Tokenizer{
		sep: sep,
		bl:  keepBlank,
		eof: discardEOF,
	}
}

func (t Tokenizer) Tokenize(dst Tokens, p []byte) Tokens {
	return t.TokenizeString(dst, fastconv.B2S(p))
}

func (t Tokenizer) TokenizeString(dst Tokens, s string) Tokens {
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
