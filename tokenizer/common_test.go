package tokenizer

import (
	"testing"

	"github.com/koykov/byteseq"
	"github.com/koykov/nlp"
)

type stage[T byteseq.Byteseq] struct {
	key string
	src T
	exp []string
}

var stagesCommon = []stage[string]{
	{
		key: "empty input",
		exp: nil,
	},
	{
		key: "no tokens",
		src: "no-tokens",
		exp: []string{"no-tokens"},
	},
}

func testInstance[T byteseq.Byteseq](t *testing.T, tkn nlp.Tokenizer[T], stages []stage[T]) {
	for _, stg := range stages {
		t.Run(stg.key, func(t *testing.T) {
			var buf nlp.Tokens
			buf = tkn.AppendTokenize(buf[:0], stg.src)
			if !assertTokens(buf, stg.exp) {
				t.Errorf("tokens mismatch: %s", stg.key)
			}
		})
	}
}

func benchInstance[T byteseq.Byteseq](b *testing.B, tkn nlp.Tokenizer[T], stages []stage[T]) {
	for _, stg := range stages {
		b.Run(stg.key, func(b *testing.B) {
			b.ReportAllocs()
			var buf nlp.Tokens
			for i := 0; i < b.N; i++ {
				buf = tkn.AppendTokenize(buf[:0], stg.src)
			}
			_ = buf
		})
	}
}

func assertTokens(tok nlp.Tokens, exp []string) (ok bool) {
	ok = true
	if ok = len(tok) == len(exp); !ok {
		return
	}
	tok.Each(func(i int, t1 nlp.Token) {
		if ok = t1.String() == exp[i]; !ok {
			return
		}
	})
	return
}
