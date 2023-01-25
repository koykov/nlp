package nlp

import (
	"github.com/koykov/bytealg"
)

type TokenizerBlankLines int

const (
	TokenizerBlankLinesDiscard TokenizerBlankLines = iota
	TokenizerBlankLinesKeep
	TokenizerBlankLinesDiscardEOF

	DefaultTokenSeparator = " \n\t"
)

type Tokenizer[T Byteseq] interface {
	Tokenize(x T) Tokens
	AppendTokenize(dst Tokens, x T) Tokens
}

type StringTokenizer[T Byteseq] struct {
	Separator  string
	BlankLines TokenizerBlankLines

	o   bool
	sep string
	bl  TokenizerBlankLines
}

func NewStringTokenizer[T Byteseq](sep string, blankLines TokenizerBlankLines) StringTokenizer[T] {
	return StringTokenizer[T]{
		Separator:  sep,
		BlankLines: blankLines,
	}
}

func (t *StringTokenizer[T]) Tokenize(x T) Tokens {
	return t.AppendTokenize(nil, x)
}

func (t *StringTokenizer[T]) AppendTokenize(dst Tokens, x T) Tokens {
	s := q2s(x)
	if len(s) == 0 {
		return dst
	}

	if !t.o {
		t.sep = t.Separator
		t.bl = t.BlankLines
		t.o = true
	}

	lo, hi := 0, 0
	for {
		p := bytealg.IndexAnyAtStr(s, t.sep, lo)
		if p == -1 {
			hi = len(s)
			if hi == lo && t.bl == TokenizerBlankLinesDiscardEOF || t.bl == TokenizerBlankLinesDiscard {
				break
			}
			dst = append(dst, ParseToken(s, lo, hi))
			break
		}
		hi = p
		if hi == lo && t.bl == TokenizerBlankLinesDiscard {
			lo = hi + 1
			continue
		}
		dst = append(dst, ParseToken(s, lo, hi))
		lo = hi + 1
	}
	return dst
}
