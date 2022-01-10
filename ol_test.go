package nlp

import "testing"

func TestOL(t *testing.T) {
	t.Run("OL32", func(t *testing.T) {
		var (
			x      OL32
			lo, hi uint16
		)
		lo, hi = 12, 16
		x.Encode(lo, hi)
		if x != 786448 {
			t.Error("Encode fail")
		}
		if l, h := x.Decode(); l != lo || h != hi {
			t.Error("Decode fail")
		}
	})
}

func BenchmarkOL(b *testing.B) {
	b.Run("OL32", func(b *testing.B) {
		var (
			x      OL32
			lo, hi uint16
		)
		lo, hi = 12, 16
		for i := 0; i < b.N; i++ {
			x.Encode(lo, hi)
			x.Decode()
		}
	})
}
