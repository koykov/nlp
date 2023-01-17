package nlp

import (
	"bytes"
	"testing"
)

var tokenStages = []struct {
	key  string
	p, e []byte
	s, r string
	l, h int
}{
	{key: "empty"},
	{key: "lorem bytes", p: []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit."), e: []byte("ipsum"), l: 6, h: 11},
	{key: "lorem string", s: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", r: "ipsum", l: 6, h: 11},
}

func TestToken(t *testing.T) {
	for _, stage := range tokenStages {
		t.Run(stage.key, func(t *testing.T) {
			if len(stage.s) > 0 {
				tok := ParseToken(stage.s, stage.l, stage.h)
				if tok.String() != stage.r {
					t.FailNow()
				}
			} else {
				tok := ParseToken(stage.p, stage.l, stage.h)
				if !bytes.Equal(tok.Bytes(), stage.e) {
					t.FailNow()
				}
			}
		})
	}
}

func BenchmarkToken(b *testing.B) {
	for _, stage := range tokenStages {
		b.Run(stage.key, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				if len(stage.s) > 0 {
					tok := ParseToken(stage.s, stage.l, stage.h)
					if tok.String() != stage.r {
						b.FailNow()
					}
				} else {
					tok := ParseToken(stage.p, stage.l, stage.h)
					if !bytes.Equal(tok.Bytes(), stage.e) {
						b.FailNow()
					}
				}
			}
		})
	}
}
