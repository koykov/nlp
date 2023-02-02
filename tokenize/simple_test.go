package tokenize

import (
	"testing"

	"github.com/koykov/nlp"
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
	testInstance[string](t, NewSpaceTokenizer[string](), stagesCommon)
	testInstance[string](t, NewSpaceTokenizer[string](), stagesSpace)

	testInstance[string](t, NewTabTokenizer[string](), stagesCommon)
	testInstance[string](t, NewTabTokenizer[string](), stagesTab)

	testInstance[string](t, NewCharTokenizer[string](), stagesChar)

	testInstance[string](t, NewLineTokenizer[string](nlp.TokenizerBlankLinesDiscard), stagesCommon)
	testInstance[string](t, NewLineTokenizer[string](nlp.TokenizerBlankLinesDiscard), stagesLine)
}

func BenchmarkSimple(b *testing.B) {
	benchInstance[string](b, NewSpaceTokenizer[string](), stagesCommon)
	benchInstance[string](b, NewSpaceTokenizer[string](), stagesSpace)

	benchInstance[string](b, NewTabTokenizer[string](), stagesCommon)
	benchInstance[string](b, NewTabTokenizer[string](), stagesTab)

	benchInstance[string](b, NewCharTokenizer[string](), stagesChar)

	benchInstance[string](b, NewLineTokenizer[string](nlp.TokenizerBlankLinesDiscard), stagesCommon)
	benchInstance[string](b, NewLineTokenizer[string](nlp.TokenizerBlankLinesDiscard), stagesLine)
}
