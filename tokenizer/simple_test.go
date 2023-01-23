package tokenizer

import (
	"testing"
)

func TestSimple(t *testing.T) {
	testInstance[string](t, NewSimple[string]())
}

func BenchmarkSimple(b *testing.B) {
	benchInstance[string](b, NewSimple[string]())
}
