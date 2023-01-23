package tokenizer

import (
	"testing"
)

func TestSimple(t *testing.T) {
	testInstance[string](t, Simple[string]{})
}
