package tokenizer

import (
	"github.com/koykov/byteseq"
	"github.com/koykov/nlp"
)

type TweetTokenizer[T byteseq.Byteseq] struct {
	PreserveCase      bool
	ReduceLen         bool
	StripHandles      bool
	MatchPhoneNumbers bool

	o bool
}

func NewTweetTokenizer[T byteseq.Byteseq](preserveCase, reduceLen, stripHandles, matchPhoneNumbers bool) TweetTokenizer[T] {
	return TweetTokenizer[T]{
		PreserveCase:      preserveCase,
		ReduceLen:         reduceLen,
		StripHandles:      stripHandles,
		MatchPhoneNumbers: matchPhoneNumbers,
		o:                 true,
	}
}

func (t TweetTokenizer[T]) Tokenize(x T) nlp.Tokens {
	return t.AppendTokenize(nil, x)
}

func (t TweetTokenizer[T]) AppendTokenize(dst nlp.Tokens, x T) nlp.Tokens {
	_ = x
	// todo: implement me
	return dst
}
