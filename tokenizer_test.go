package nlp

import "testing"

const (
	tknzSrc = "Good muffins cost $3.88\tin New York.  Please buy me\ntwo of them.\n\nThanks.\n\n\n"
	tknzSep = " \n\t"
)

var tokenizerStages = []struct {
	key    string
	src    string
	sep    string
	keepBL bool
	disEOF bool
	bl     TokenizerBlankLines
	tok    []string
}{
	{
		key:    "keep blank lines",
		src:    tknzSrc,
		sep:    tknzSep,
		keepBL: true,
		bl:     TokenizerBlankLinesKeep,
		tok:    []string{"Good", "muffins", "cost", "$3.88", "in", "New", "York.", "", "Please", "buy", "me", "two", "of", "them.", "", "Thanks.", "", "", ""},
	},
	{
		key: "discard blank lines",
		src: tknzSrc,
		sep: tknzSep,
		tok: []string{"Good", "muffins", "cost", "$3.88", "in", "New", "York.", "Please", "buy", "me", "two", "of", "them.", "Thanks."},
	},
	{
		key:    "discard EOF",
		src:    tknzSrc,
		sep:    tknzSep,
		keepBL: true,
		disEOF: true,
		bl:     TokenizerBlankLinesDiscardEOF,
		tok:    []string{"Good", "muffins", "cost", "$3.88", "in", "New", "York.", "", "Please", "buy", "me", "two", "of", "them.", "", "Thanks.", "", ""},
	},
}

func TestTokenizer(t *testing.T) {
	for _, stage := range tokenizerStages {
		t.Run(stage.key, func(t *testing.T) {
			ctx := NewCtx[string]()
			ctx.SetText(stage.src).
				WithTokenizer(NewStringTokenizer[string](stage.sep, stage.bl)).
				Tokenize().
				GetTokens().
				Each(func(i int, tok Token) {
					if tok.String() != stage.tok[i] {
						t.FailNow()
					}
				})
		})
	}
}

func BenchmarkTokenizer(b *testing.B) {
	for _, stage := range tokenizerStages {
		b.Run(stage.key, func(b *testing.B) {
			b.ReportAllocs()
			var buf Tokens
			ctx := NewCtx[string]()
			// Declare tokenizer outside of loop due to redundant allocations.
			// See https://github.com/koykov/lab/tree/master/generic_value_alloc BenchmarkGVAlloc/string/* cases.
			tkn := StringTokenizer[string]{stage.sep, stage.bl}
			for i := 0; i < b.N; i++ {
				buf = ctx.Reset().
					SetText(stage.src).
					WithTokenizer(&tkn).
					Tokenize().
					GetTokens()
			}
			_ = buf
		})
	}
}
