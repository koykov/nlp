package tokenizer

import (
	"testing"
)

var stagesSpace = []stage[string]{
	{
		key: "tokens",
		src: "Good muffins cost $3.88\nin New York.  Please buy me\ntwo of them.\n\nThanks.",
		exp: []string{"Good", "muffins", "cost", "$3.88\nin", "New", "York.", "", "Please", "buy", "me\ntwo", "of", "them.\n\nThanks."},
	},
}

func TestSpace(t *testing.T) {
	testInstance[string](t, Space[string]{}, stagesCommon)
	testInstance[string](t, Space[string]{}, stagesSpace)
}

func BenchmarkSpace(b *testing.B) {
	benchInstance[string](b, Space[string]{}, stagesCommon)
	benchInstance[string](b, Space[string]{}, stagesSpace)
}
