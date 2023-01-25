package tokenizer

import (
	"testing"
)

var stagesTab = []stage[string]{
	{
		key: "tokens",
		src: "a\tb c\n\t d",
		exp: []string{"a", "b c\n", " d"},
	},
}

func TestTab(t *testing.T) {
	testInstance[string](t, Tab[string]{}, stagesCommon)
	testInstance[string](t, Tab[string]{}, stagesTab)
}

func BenchmarkTab(b *testing.B) {
	benchInstance[string](b, Tab[string]{}, stagesCommon)
	benchInstance[string](b, Tab[string]{}, stagesTab)
}
