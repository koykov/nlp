package nlp

import "testing"

func TestScript(t *testing.T) {
	t.Run("eval-rune", func(t *testing.T) {
		r := '汉'
		if !ScriptHan.Evaluate(r) {
			t.FailNow()
		}
		r = 'Ю'
		if !ScriptCyrillic.Evaluate(r) {
			t.FailNow()
		}
	})
}

func BenchmarkScript(b *testing.B) {
	b.Run("is-rune", func(b *testing.B) {
		r := '汉'
		for i := 0; i < b.N; i++ {
			if !ScriptHan.Is(r) {
				b.FailNow()
			}
		}
	})
	b.Run("eval-rune", func(b *testing.B) {
		r := '汉'
		for i := 0; i < b.N; i++ {
			if !ScriptHan.Evaluate(r) {
				b.FailNow()
			}
		}
	})
}
