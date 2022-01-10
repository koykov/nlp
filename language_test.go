package nlp

import "testing"

func TestLanguage(t *testing.T) {
	assert := func(t testing.TB, x, e string) {
		if x != e {
			t.Errorf("assertion failed: need '%s', got '%s'", e, x)
		}
	}
	t.Run("name", func(t *testing.T) { assert(t, Russian.String(), "Russian") })
	t.Run("native", func(t *testing.T) { assert(t, Russian.Native(), "Русский") })
	t.Run("iso639-1", func(t *testing.T) { assert(t, Russian.Iso6391(), "ru") })
	t.Run("iso639-3", func(t *testing.T) { assert(t, Russian.Iso6393(), "rus") })
}

func BenchmarkLanguage(b *testing.B) {
	assert := func(b *testing.B, l Language, typ uint) {
		var x string
		for i := 0; i < b.N; i++ {
			switch typ {
			case 0:
				x = l.String()
			case 1:
				x = l.Native()
			case 2:
				x = l.Iso6391()
			case 3:
				x = l.Iso6393()
			}
		}
		_ = x
	}
	b.Run("name", func(b *testing.B) { assert(b, Russian, 0) })
	// b.Run("native", func(b *testing.B) { assert(b, Russian, 1) })
	// b.Run("iso639-1", func(b *testing.B) { assert(b, Russian, 2) })
	// b.Run("iso639-3", func(b *testing.B) { assert(b, Russian, 3) })
}
