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
		n := NewUnigram(r[0])
		if n.String() != s {
			t.FailNow()
		}
	})
	t.Run("string/bigram", func(t *testing.T) {
		s := "村変"
		r := fastconv.AppendS2R(nil, s)
		n := NewBigram(r[0], r[1])
		if n.String() != s {
			t.FailNow()
		}
	})
	t.Run("string/trigram", func(t *testing.T) {
		s := "村変界"
		r := fastconv.AppendS2R(nil, s)
		n := NewTrigram(r[0], r[1], r[2])
		if n.String() != s {
			t.FailNow()
		}
	})
	t.Run("string/quadrigram", func(t *testing.T) {
		s := "村変界広"
		r := fastconv.AppendS2R(nil, s)
		n := NewQuadrigram(r[0], r[1], r[2], r[3])
		if n.String() != s {
			t.FailNow()
		}
	})
	t.Run("string/fivegram", func(t *testing.T) {
		s := "村変界広共"
		r := fastconv.AppendS2R(nil, s)
		n := NewFivegram(r[0], r[1], r[2], r[3], r[4])
		if n.String() != s {
			t.FailNow()
		}
	})
}
