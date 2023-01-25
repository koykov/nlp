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
			tkn := NewStringTokenizer[string](stage.sep, stage.bl)
			tkn.AppendTokenize(nil, stage.src).Each(func(i int, tok Token) {
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
			var buf Tokens
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				tkn := NewStringTokenizer[string](stage.sep, stage.bl)
				buf = tkn.AppendTokenize(buf[:0], stage.src)
			}
			_ = buf
		})
	}
}
