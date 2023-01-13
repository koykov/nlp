package nlp

import (
	"testing"
	"unsafe"
)

func TestNgram(t *testing.T) {
	assertSize := func(t testing.TB, a, b uint64) {
		if a != b {
			t.Errorf("size mismatch: need %d, got %d", b, a)
		}
	}
	t.Run("sizeof(unigram)", func(t *testing.T) { assertSize(t, uint64(unsafe.Sizeof(Unigram(0))), 2) })
	t.Run("sizeof(bigram)", func(t *testing.T) { assertSize(t, uint64(unsafe.Sizeof(Bigram(0))), 4) })
	t.Run("sizeof(trigram)", func(t *testing.T) { assertSize(t, uint64(unsafe.Sizeof(Trigram{})), 6) })
	t.Run("sizeof(quadrigram)", func(t *testing.T) { assertSize(t, uint64(unsafe.Sizeof(Quadrigram(0))), 8) })
	t.Run("sizeof(fivegram)", func(t *testing.T) { assertSize(t, uint64(unsafe.Sizeof(Fivegram{})), 10) })
}
