package tokenizer

import (
	"testing"

	"github.com/koykov/nlp"
)

type stage[T nlp.Byteseq] struct {
	key string
	src T
	exp []string
}

var stages = []stage[string]{
	{
		key: "empty input",
		exp: nil,
	},
	{
		key: "no tokens",
		src: "no tokens",
		exp: []string{"no tokens"},
	},
}

func testInstance[T nlp.Byteseq](t *testing.T, tkn nlp.TokenizerInterface[T]) {
	for _, stg := range stages {
		t.Run(stg.key, func(t *testing.T) {
			var buf nlp.Tokens
			buf = tkn.Tokenize(buf[:0], T(stg.src))
			if !assertTokens(buf, stg.exp) {
				t.Errorf("tokens mismatch: %s", stg.key)
			}
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
