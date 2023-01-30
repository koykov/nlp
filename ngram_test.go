package nlp

import (
	"testing"
	"unsafe"

	"github.com/koykov/fastconv"
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

	t.Run("string/unigram", func(t *testing.T) {
		s := "村"
		r := fastconv.AppendS2R(nil, s)
		u := Unigram(r[0])
		if u.String() != s {
			t.FailNow()
		}
	})
	// t.Run("string/bigram", func(t *testing.T) {
	// 	s := "村変"
	// 	r := fastconv.AppendS2R(nil, s)
	// 	u := Bigram(r[0])
	// 	if u.String() != s {
	// 		t.FailNow()
	// 	}
	// })
}
