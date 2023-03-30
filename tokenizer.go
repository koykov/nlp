package nlp

import (
	"github.com/koykov/bytealg"
	"github.com/koykov/byteseq"
)

type TokenizerBlankLines int

const (
	TokenizerBlankLinesDiscard TokenizerBlankLines = iota
	TokenizerBlankLinesKeep
	TokenizerBlankLinesDiscardEOF

	DefaultTokenSeparator = " \n\t"
)

type StringTokenizer[T byteseq.Byteseq] struct {
	Separator  string
	BlankLines TokenizerBlankLines
}

func NewStringTokenizer[T byteseq.Byteseq](sep string, blankLines TokenizerBlankLines) StringTokenizer[T] {
	return StringTokenizer[T]{
		Separator:  sep,
		BlankLines: blankLines,
	}
}

func (t StringTokenizer[T]) Tokenize(x T) Tokens {
	return t.AppendTokenize(nil, x)
}

func (t StringTokenizer[T]) AppendTokenize(dst Tokens, x T) Tokens {
	s := byteseq.Q2S(x)
	if len(s) == 0 {
		return dst
	}

	sep, bl := t.Separator, t.BlankLines
	lo, hi := 0, 0
	for {
		p := bytealg.IndexAnyAtStr(s, sep, lo)
		if p == -1 {
			hi = len(s)
			if hi == lo && (bl == TokenizerBlankLinesDiscardEOF || bl == TokenizerBlankLinesDiscard) {
				break
			}
			dst = append(dst, ParseToken(s, lo, hi))
			break
		}
		hi = p
		if hi == lo && bl == TokenizerBlankLinesDiscard {
			lo = hi + 1
			continue
		}
		dst = append(dst, ParseToken(s, lo, hi))
		lo = hi + 1
	}
	return dst
}
