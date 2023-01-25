package tokenizer

import (
	"testing"
)

var stagesLine = []stage[string]{
	{
		key: "tokens",
		src: "Good muffins cost $3.88\nin New York.  Please buy me\ntwo of them.\n\nThanks.",
		exp: []string{"Good muffins cost $3.88", "in New York.  Please buy me", "two of them.", "Thanks."},
	},
}

func TestLine(t *testing.T) {
	testInstance[string](t, Line[string]{}, stagesCommon)
	testInstance[string](t, Line[string]{}, stagesLine)
}

func BenchmarkLine(b *testing.B) {
	benchInstance[string](b, Line[string]{}, stagesCommon)
	benchInstance[string](b, Line[string]{}, stagesLine)
}
