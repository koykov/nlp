package tokenizer

import (
	"regexp"

	"github.com/koykov/byteseq"
	"github.com/koykov/nlp"
)

var (
	reEmoticons = regexp.MustCompile(`(?:[<>]?[:;=8][\-o*']?[)\](\[dDpP/:}{@|\\]|[)\](\[dDpP/:}{@|\\][\-o*']?[:;=8][<>]?|</?3)`)
	reURL       = regexp.MustCompile(`(?:https?:(?:/{1,3}|[a-z0-9%])|[a-z0-9.\-]+[.](?:[a-z]{2,13})/)(?:[^\s()<>{}\[\]]+|\([^\s()]*?\([^\s()]+\)[^\s()]*?\)|\([^\s]+?\))+(?:\([^\s()]*?\([^\s()]+\)[^\s()]*?\)|\([^\s]+?\)|[^\s` + "`" + `!()\[\]{};:'".,<>?«»“”‘’])|(?:(?<!@)[a-z0-9]+(?:[.\-][a-z0-9]+)*[.](?:[a-z]{2,13})\b/?(?!@)`)
	reFlag      = regexp.MustCompile("(?:[\U0001F1E6-\U0001F1FF]{2}|\U0001F3F4\U000E0067\U000E0062\U000E0065\U000E006e\U000E0067\U000E007F|\U0001F3F4\U000E0067\U000E0062\U000E0073\U000E0063\U000E0074\U000E007F|\U0001F3F4\U000E0067\U000E0062\U000E0077\U000E006C\U000E0073\U000E007F)")
	rePhone     = regexp.MustCompile(`(?:(?:\+?[01][ *\-.)]*)?(?:[(]?\d{3}[ *\-.)]*)?\d{3}[ *\-.)]*\d{4})`)
)

type TweetTokenizer[T byteseq.Byteseq] struct {
	PreserveCase      bool
	ReduceLen         bool
	StripHandles      bool
	MatchPhoneNumbers bool

	pc, rl, sh, mp, o bool
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

func (t *TweetTokenizer[T]) Tokenize(x T) nlp.Tokens {
	return t.AppendTokenize(nil, x)
}

func (t *TweetTokenizer[T]) AppendTokenize(dst nlp.Tokens, x T) nlp.Tokens {
	if !t.o {
		t.pc, t.rl, t.sh, t.mp, t.o = t.PreserveCase, t.ReduceLen, t.StripHandles, t.MatchPhoneNumbers, true
	}
	_ = x
	// todo: implement me
	return dst
}
