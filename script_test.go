package nlp

import "testing"

func TestScript(t *testing.T) {
	t.Run("eval-rune", func(t *testing.T) {
		r := '汉'
		if !ScriptHan.EvaluateRune(r) {
			t.FailNow()
		}
		r = 'Ю'
		if !ScriptCyrillic.EvaluateRune(r) {
			t.FailNow()
		}
	})
}

func BenchmarkScript(b *testing.B) {
	b.Run("eval-rune", func(b *testing.B) {
		r := '汉'
		for i := 0; i < b.N; i++ {
			if !ScriptHan.EvaluateRune(r) {
				b.FailNow()
			}
		}
	})
	b.Run("eval-rune1", func(b *testing.B) {
		r := '汉'
		for i := 0; i < b.N; i++ {
			if !ScriptHan.EvaluateRune1(r) {
				b.FailNow()
			}
		}
	})
	b.Run("eval-rune2", func(b *testing.B) {
		r := '汉'
		for i := 0; i < b.N; i++ {
			if !ScriptHan.EvaluateRune2(r) {
				b.FailNow()
			}
		}
	})
}
