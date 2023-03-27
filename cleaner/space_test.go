package cleaner

import (
	"testing"

	"github.com/koykov/fastconv"
)

var stagesSpace = []stage{
	{key: "empty input", src: "", exp: ""},
	{
		key: "copy",
		src: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		exp: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	},
	{
		key: "many spaces",
		src: "Arkansas                     - Little Rock",
		exp: "Arkansas - Little Rock",
	},
	{
		key: "rtrim",
		src: "foobar    ",
		exp: "foobar",
	},
	{
		key: "ltrim",
		src: "    foobar",
		exp: "foobar",
	},
	{
		key: "trim",
		src: "   foobar   ",
		exp: "foobar",
	},
	{
		key: "mixed",
		src: "    Arkansas                     - Little Rock      ",
		exp: "Arkansas - Little Rock",
	},
}

func TestSpace(t *testing.T) {
	for _, stage := range stagesSpace {
		t.Run(stage.key, func(t *testing.T) {
			c := Space[string]{}
			r := c.Clean(stage.src)
			_, s := fastconv.AppendR2S(nil, r)
			if s != stage.exp {
				t.FailNow()
			}
		})
	}
}

func BenchmarkSpace(b *testing.B) {
	for _, stage := range stagesSpace {
		b.Run(stage.key, func(b *testing.B) {
			b.ReportAllocs()
			var r []rune
			c := Space[string]{}
			for i := 0; i < b.N; i++ {
				r = c.AppendClean(r[:0], stage.src)
			}
			_ = r
		})
	}
}
