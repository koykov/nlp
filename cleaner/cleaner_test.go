package cleaner

import (
	"testing"

	"github.com/koykov/fastconv"
)

var stagesMacros = []stage{
	{"empty input", "", ""},
	{
		"copy",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	},
}

func TestMacros(t *testing.T) {
	for _, stage := range stagesMacros {
		t.Run(stage.key, func(t *testing.T) {
			c := Macros[string]{Left: "[", Right: "]"}
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
		var r []rune
		c := Macros[string]{Left: "[", Right: "]"}
		b.Run(stage.key, func(b *testing.B) {
			r = c.Clean(stage.src)
		})
		_ = r
	}
}
