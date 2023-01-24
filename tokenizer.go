package nlp

import (
	"github.com/koykov/bitset"
	"github.com/koykov/bytealg"
)

const (
	DefaultTokenSeparator = " \n\t"

	tknFlagInit           = 0
	tknFlagKeepBlankLines = 1
	tknFlagDiscardEOF     = 2
)

type Tokenizer[T Byteseq] interface {
	Tokenize(dst Tokens, x T) Tokens
}

type StringTokenizer[T Byteseq] struct {
	bs  bitset.Bitset8
	sep string
}

func NewStringTokenizer[T Byteseq](sep string, keepBlank bool, discardEOF bool) StringTokenizer[T] {
	var bs bitset.Bitset8
	bs.SetBit(tknFlagInit, true)
	bs.SetBit(tknFlagKeepBlankLines, keepBlank)
	bs.SetBit(tknFlagDiscardEOF, discardEOF)
	return StringTokenizer[T]{
		sep: sep,
		bs:  bs,
	}
}

func (t StringTokenizer[T]) Tokenize(dst Tokens, x T) Tokens {
	s := q2s(x)
	if len(s) == 0 {
		return dst
	}

	if !t.bs.CheckBit(tknFlagInit) {
		t.sep = DefaultTokenSeparator
		t.bs.SetBit(tknFlagInit, true)
		t.bs.SetBit(tknFlagKeepBlankLines, true)
		t.bs.SetBit(tknFlagDiscardEOF, true)
	}

	lo, hi := 0, 0
	for {
		p := bytealg.IndexAnyAtStr(s, t.sep, lo)
		if p == -1 {
			hi = len(s)
			if hi == lo && t.bs.CheckBit(tknFlagDiscardEOF) {
				break
			}
			dst = append(dst, ParseToken(s, lo, hi))
			break
		}
		hi = p
		if hi == lo && !t.bs.CheckBit(tknFlagKeepBlankLines) {
			lo = hi + 1
			continue
		}
		dst = append(dst, ParseToken(s, lo, hi))
		lo = hi + 1
	}
	return dst
}
