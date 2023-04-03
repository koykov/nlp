package modifier

import (
	"fmt"
	"testing"

	"github.com/koykov/nlp"
)

var (
	stagesUpper = []stage[string]{
		{"no input", "", ""},
		{"copy", "ONLYUPPER", "ONLYUPPER"},
		{"ascii", "abc", "ABC"},
		{"ascii 1", "AbC123", "ABC123"},
		{"ascii 2", "azAZ09_", "AZAZ09_"},
		{"long string", "longStrinGwitHmixofsmaLLandcAps", "LONGSTRINGWITHMIXOFSMALLANDCAPS"},
		{"long words", "RENAN BASTOS 93 AOSDAJDJAIDJAIDAJIaidsjjaidijadsjiadjiOOKKO", "RENAN BASTOS 93 AOSDAJDJAIDJAIDAJIAIDSJJAIDIJADSJIADJIOOKKO"},
		{"long string no ascii", "long\u0250string\u0250with\u0250nonascii\u2C6Fchars", "LONG\u2C6FSTRING\u2C6FWITH\u2C6FNONASCII\u2C6FCHARS"},
		{"grows one byte per char", "\u0250\u0250\u0250\u0250\u0250", "\u2C6F\u2C6F\u2C6F\u2C6F\u2C6F"},
		{"self and max rune", "a\u0080\U0010FFFF", "A\u0080\U0010FFFF"},
		{"cyrillic", "привет мир!", "ПРИВЕТ МИР!"},
	}

	stagesLower = []stage[string]{
		{"no input", "", ""},
		{"copy", "abc", "abc"},
		{"ascii", "AbC123", "abc123"},
		{"ascii 1", "azAZ09_", "azaz09_"},
		{"long string", "longStrinGwitHmixofsmaLLandcAps", "longstringwithmixofsmallandcaps"},
		{"long words", "renan bastos 93 AOSDAJDJAIDJAIDAJIaidsjjaidijadsjiadjiOOKKO", "renan bastos 93 aosdajdjaidjaidajiaidsjjaidijadsjiadjiookko"},
		{"long string no ascii", "LONG\u2C6FSTRING\u2C6FWITH\u2C6FNONASCII\u2C6FCHARS", "long\u0250string\u0250with\u0250nonascii\u0250chars"},
		{"shrinks one byte per char", "\u2C6D\u2C6D\u2C6D\u2C6D\u2C6D", "\u0251\u0251\u0251\u0251\u0251"},
		{"self and max rune", "A\u0080\U0010FFFF", "a\u0080\U0010FFFF"},
		{"cyrillic", "ПРИВЕТ МИР!", "привет мир!"},
	}

	stagesTitle = []stage[string]{
		{"no input", "", ""},
		{"single char", "a", "A"},
		{"ascii", " aaa aaa aaa ", " Aaa Aaa Aaa "},
		{"ascii 1", " Aaa Aaa Aaa ", " Aaa Aaa Aaa "},
		{"ascii 2", "123a456", "123a456"},
		{"separator", "double-blind", "Double-Blind"},
		{"no ascii", "ÿøû", "Ÿøû"},
		{"underscore", "with_underscore", "With_underscore"},
		{"unicode", "unicode \xe2\x80\xa8 line separator", "Unicode \xe2\x80\xa8 Line Separator"},
		{"cyrillic", "привет мир!", "Привет Мир!"},
	}
)

func TestCase(t *testing.T) {
	fn := func(t *testing.T, s *stage[string], mod nlp.Modifier[string]) {
		ctx := nlp.NewCtx[string]()
		r := ctx.SetText(s.src).
			WithModifier(mod).
			Modify().
			GetText()
		if r != s.exp {
			t.FailNow()
		}
	}
	for _, s := range stagesUpper {
		t.Run(fmt.Sprintf("upper/%s", s.key), func(t *testing.T) { fn(t, &s, Upper[string]{}) })
	}
	for _, s := range stagesLower {
		t.Run(fmt.Sprintf("lower/%s", s.key), func(t *testing.T) { fn(t, &s, Lower[string]{}) })
	}
	for _, s := range stagesTitle {
		t.Run(fmt.Sprintf("title/%s", s.key), func(t *testing.T) { fn(t, &s, Title[string]{}) })
	}
}

func BenchmarkCase(b *testing.B) {
	fn := func(b *testing.B, s *stage[string], mod nlp.Modifier[string]) {
		b.ReportAllocs()
		ctx := nlp.NewCtx[string]()
		for i := 0; i < b.N; i++ {
			_ = ctx.Reset().
				SetText(s.src).
				WithModifier(mod).
				Modify().
				GetText()
		}
	}
	for _, s := range stagesUpper {
		b.Run(fmt.Sprintf("upper/%s", s.key), func(b *testing.B) { fn(b, &s, Upper[string]{}) })
	}
	for _, s := range stagesLower {
		b.Run(fmt.Sprintf("lower/%s", s.key), func(b *testing.B) { fn(b, &s, Lower[string]{}) })
	}
	for _, s := range stagesTitle {
		b.Run(fmt.Sprintf("title/%s", s.key), func(b *testing.B) { fn(b, &s, Title[string]{}) })
	}
}
