package nlp

import "testing"

func TestLanguage(t *testing.T) {
	assert := func(t testing.TB, x, e string) {
		if x != e {
			t.Errorf("assertion failed: need '%s', got '%s'", e, x)
		}
	}
	t.Run("name", func(t *testing.T) { assert(t, LanguageRussian.String(), "Russian") })
	t.Run("native", func(t *testing.T) { assert(t, LanguageRussian.Native(), "Русский") })
	t.Run("iso639-1", func(t *testing.T) { assert(t, LanguageRussian.Iso6391(), "ru") })
	t.Run("iso639-3", func(t *testing.T) { assert(t, LanguageRussian.Iso6393(), "rus") })
}

func BenchmarkLanguage(b *testing.B) {
	b.Run("name", func(b *testing.B) {
		var x string
		for i := 0; i < b.N; i++ {
			x = LanguageRussian.String()
		}
		_ = x
	})
	b.Run("native", func(b *testing.B) {
		var x string
		for i := 0; i < b.N; i++ {
			x = LanguageRussian.Native()
		}
		_ = x
	})
	b.Run("iso639-1", func(b *testing.B) {
		var x string
		for i := 0; i < b.N; i++ {
			x = LanguageRussian.Iso6391()
		}
		_ = x
	})
	b.Run("iso639-3", func(b *testing.B) {
		var x string
		for i := 0; i < b.N; i++ {
			x = LanguageRussian.Iso6393()
		}
		_ = x
	})
}
