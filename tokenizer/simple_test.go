package tokenizer

import (
	"testing"
)

func TestSimple(t *testing.T) {
	testInstance[string](t, Simple[string]{})
}

func BenchmarkSimple(b *testing.B) {
	benchInstance[string](b, Simple[string]{})
}
