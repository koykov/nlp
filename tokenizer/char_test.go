package tokenizer

import (
	"testing"
)

var stagesChar = []stage[string]{
	{
		key: "tokens",
		src: "Lorem ipsum\tdolor sit amet.",
		exp: []string{"L", "o", "r", "e", "m", " ", "i", "p", "s", "u", "m", "\t", "d", "o", "l", "o", "r", " ", "s", "i", "t", " ", "a", "m", "e", "t", "."},
	},
}

func TestChar(t *testing.T) {
	testInstance[string](t, Char[string]{}, stagesChar)
}

func BenchmarkChar(b *testing.B) {
	benchInstance[string](b, Char[string]{}, stagesChar)
}
