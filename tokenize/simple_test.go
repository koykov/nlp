package tokenize

import (
	"testing"
)

var (
	stagesSpace = []stage[string]{
		{
			key: "tokens",
			src: "Good muffins cost $3.88\nin New York.  Please buy me\ntwo of them.\n\nThanks.",
			exp: []string{"Good", "muffins", "cost", "$3.88\nin", "New", "York.", "", "Please", "buy", "me\ntwo", "of", "them.\n\nThanks."},
		},
	}
	stagesTab = []stage[string]{
		{
			key: "tokens",
			src: "a\tb c\n\t d",
			exp: []string{"a", "b c\n", " d"},
		},
	}
	stagesChar = []stage[string]{
		{
			key: "tokens",
			src: "Lorem ipsum\tdolor sit amet.",
			exp: []string{"L", "o", "r", "e", "m", " ", "i", "p", "s", "u", "m", "\t", "d", "o", "l", "o", "r", " ", "s", "i", "t", " ", "a", "m", "e", "t", "."},
		},
	}
	stagesLine = []stage[string]{
		{
			key: "tokens",
			src: "Good muffins cost $3.88\nin New York.  Please buy me\ntwo of them.\n\nThanks.",
			exp: []string{"Good muffins cost $3.88", "in New York.  Please buy me", "two of them.", "Thanks."},
		},
	}
)

func TestSimple(t *testing.T) {
	testInstance[string](t, SpaceTokenizer[string]{}, stagesCommon)
	testInstance[string](t, SpaceTokenizer[string]{}, stagesSpace)

	testInstance[string](t, TabTokenizer[string]{}, stagesCommon)
	testInstance[string](t, TabTokenizer[string]{}, stagesTab)

	testInstance[string](t, CharTokenizer[string]{}, stagesChar)

	testInstance[string](t, LineTokenizer[string]{}, stagesCommon)
	testInstance[string](t, LineTokenizer[string]{}, stagesLine)
}

func BenchmarkSimple(b *testing.B) {
	benchInstance[string](b, SpaceTokenizer[string]{}, stagesCommon)
	benchInstance[string](b, SpaceTokenizer[string]{}, stagesSpace)

	benchInstance[string](b, TabTokenizer[string]{}, stagesCommon)
	benchInstance[string](b, TabTokenizer[string]{}, stagesTab)

	benchInstance[string](b, CharTokenizer[string]{}, stagesChar)

	benchInstance[string](b, LineTokenizer[string]{}, stagesCommon)
	benchInstance[string](b, LineTokenizer[string]{}, stagesLine)
}
