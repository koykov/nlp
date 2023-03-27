package cleaner

import (
	"testing"

	"github.com/koykov/fastconv"
)

var stagesMacros = []stage{
	{key: "empty input", src: "", exp: ""},
	{
		key: "copy",
		src: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		exp: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		l:   "[", r: "]",
	},
	{
		key: "square",
		src: "Alsatian [adj]",
		exp: "Alsatian ",
		l:   "[", r: "]",
	},
	{
		key: "curly",
		src: "Arctic Circle {n}",
		exp: "Arctic Circle ",
		l:   "{", r: "}",
	},
}

func TestMacros(t *testing.T) {
	for _, stage := range stagesMacros {
		t.Run(stage.key, func(t *testing.T) {
			c := Macros[string]{Left: stage.l, Right: stage.r}
			r := c.Clean(stage.src)
			_, s := fastconv.AppendR2S(nil, r)
			if s != stage.exp {
				t.FailNow()
			}
		})
	}
}

func BenchmarkMacros(b *testing.B) {
	for _, stage := range stagesMacros {
		b.Run(stage.key, func(b *testing.B) {
			b.ReportAllocs()
			var r []rune
			c := Macros[string]{Left: "[", Right: "]"}
			for i := 0; i < b.N; i++ {
				r = c.AppendClean(r[:0], stage.src)
			}
			_ = r
		})
	}
}
