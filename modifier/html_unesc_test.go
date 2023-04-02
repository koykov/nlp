package modifier

import (
	"testing"

	"github.com/koykov/nlp"
)

var stagesHTMLU = []stage[string]{
	{"empty input", "", ""},
	{
		"copy",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	},
	{"simple", "&amp; &gt; &lt;", "& > <"},
	{"stringEnd", "&amp &amp", "& &"},
	{"multiCodepoint", "text &gesl; blah", "text \u22db\ufe00 blah"},
	{"decimalEntity", "Delta = &#916; ", "Delta = Δ "},
	{"hexadecimalEntity", "Lambda = &#x3bb; = &#X3Bb ", "Lambda = λ = λ "},
	{"numericEnds", "&# &#x &#128;43 &copy = &#169f = &#xa9", "&# &#x €43 © = ©f = ©"},
	{"numericReplacements", "Footnote&#x87;", "Footnote‡"},
	{"copySingleAmpersand", "&", "&"},
	{"copyAmpersandNonEntity", "text &test", "text &test"},
	{"copyAmpersandHash", "text &#", "text &#"},
}

func TestHTMLUnescaper(t *testing.T) {
	f := func(t *testing.T, s *stage[string]) {
		ctx := nlp.NewCtx[string]()
		r := ctx.SetText(s.src).
			WithModifier(HTMLUnescaper[string]{}).
			Modify().
			GetText()
		if r != s.exp {
			t.FailNow()
		}
	}
	for _, s := range stagesHTMLU {
		t.Run(s.key, func(t *testing.T) { f(t, &s) })
	}
}

func BenchmarkHTMLUnescaper(b *testing.B) {
	f := func(b *testing.B, s *stage[string]) {
		b.ReportAllocs()
		ctx := nlp.NewCtx[string]()
		mod := HTMLUnescaper[string]{}
		var buf []rune
		for i := 0; i < b.N; i++ {
			buf = ctx.Reset().
				SetText(s.src).
				WithModifier(&mod).
				Modify().
				GetRunes()
		}
		_ = buf
	}
	for _, s := range stagesHTMLU {
		b.Run(s.key, func(b *testing.B) { f(b, &s) })
	}
}
