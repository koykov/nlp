package nlp

import (
	"testing"

	"github.com/koykov/fastconv"
)

const clnTestSource = "control \u009f print ´ graphic ï mark \u035f punct { space \u00a0 digit 5 number ² symbol = letter ì"
const clnTestSourceCyr = "контроль \u009f печатные ´ графические ï маркировка \u035f пунктуация { пробел \u00a0 цифра 5 число ² символ = буква Я"

var clnStages = []struct {
	key  string
	src  string
	mask uint32
	exp  string
}{
	{
		key: "no clean",
		src: clnTestSource,
		exp: clnTestSource,
	},
	{
		key:  "control",
		src:  clnTestSource,
		mask: CleanControl,
		exp:  "control  print ´ graphic ï mark \u035f punct { space \u00a0 digit 5 number ² symbol = letter ì",
	},
	{
		key:  "mark",
		src:  clnTestSource,
		mask: CleanMark,
		exp:  "control \u009F print ´ graphic ï mark  punct { space \u00a0 digit 5 number ² symbol = letter ì",
	},
	{
		key:  "punct",
		src:  clnTestSource,
		mask: CleanPunct,
		exp:  "control \u009F print ´ graphic ï mark \u035f punct  space \u00a0 digit 5 number ² symbol = letter ì",
	},
	{
		key:  "space",
		src:  clnTestSource,
		mask: CleanSpace,
		exp:  "control\u009fprint´graphicïmark\u035fpunct{spacedigit5number²symbol=letterì",
	},
	{
		key:  "digit",
		src:  clnTestSource,
		mask: CleanDigit,
		exp:  "control \u009f print ´ graphic ï mark \u035f punct { space \u00a0 digit  number ² symbol = letter ì",
	},
	{
		key:  "number",
		src:  clnTestSource,
		mask: CleanNumber,
		exp:  "control \u009f print ´ graphic ï mark \u035f punct { space \u00a0 digit  number  symbol = letter ì",
	},
	{
		key:  "symbol",
		src:  clnTestSource,
		mask: CleanSymbol,
		exp:  "control \u009f print  graphic ï mark \u035f punct { space \u00a0 digit 5 number ² symbol  letter ì",
	},
	{
		key:  "letter",
		src:  clnTestSource,
		mask: CleanLetter,
		exp:  " \u009f  ´    \u035f  {  \u00a0  5  ²  =  ",
	},
	{
		key:  "print",
		src:  clnTestSource,
		mask: CleanPrint,
		exp:  "\u009F\u00a0",
	},
	{
		key:  "graphic",
		src:  clnTestSource,
		mask: CleanGraphic,
		exp:  "\u009F",
	},
	{
		key:  "default mask",
		src:  clnTestSource,
		mask: DefaultCleanMask,
		exp:  "control  print  graphic ï mark  punct  space \u00a0 digit  number  symbol  letter ì",
	},
	// cyrillic
	{
		key: "no clean cyrillic",
		src: clnTestSourceCyr,
		exp: clnTestSourceCyr,
	},
	{
		key:  "control cyrillic",
		src:  clnTestSourceCyr,
		mask: CleanControl,
		exp:  "контроль  печатные ´ графические ï маркировка \u035f пунктуация { пробел \u00a0 цифра 5 число ² символ = буква Я",
	},
	{
		key:  "mark cyrillic",
		src:  clnTestSourceCyr,
		mask: CleanMark,
		exp:  "контроль \u009f печатные ´ графические ï маркировка  пунктуация { пробел \u00a0 цифра 5 число ² символ = буква Я",
	},
	{
		key:  "punct cyrillic",
		src:  clnTestSourceCyr,
		mask: CleanPunct,
		exp:  "контроль \u009f печатные ´ графические ï маркировка \u035f пунктуация  пробел \u00a0 цифра 5 число ² символ = буква Я",
	},
	{
		key:  "space cyrillic",
		src:  clnTestSourceCyr,
		mask: CleanSpace,
		exp:  "контроль\u009fпечатные´графическиеïмаркировка\u035fпунктуация{пробелцифра5число²символ=букваЯ",
	},
	{
		key:  "digit cyrillic",
		src:  clnTestSourceCyr,
		mask: CleanDigit,
		exp:  "контроль \u009f печатные ´ графические ï маркировка \u035f пунктуация { пробел \u00a0 цифра  число ² символ = буква Я",
	},
	{
		key:  "number cyrillic",
		src:  clnTestSourceCyr,
		mask: CleanNumber,
		exp:  "контроль \u009f печатные ´ графические ï маркировка \u035f пунктуация { пробел \u00a0 цифра  число  символ = буква Я",
	},
	{
		key:  "symbol cyrillic",
		src:  clnTestSourceCyr,
		mask: CleanSymbol,
		exp:  "контроль \u009f печатные  графические ï маркировка \u035f пунктуация { пробел \u00a0 цифра 5 число ² символ  буква Я",
	},
	{
		key:  "letter cyrillic",
		src:  clnTestSourceCyr,
		mask: CleanLetter,
		exp:  " \u009f  ´    \u035f  {  \u00a0  5  ²  =  ",
	},
	{
		key:  "print cyrillic",
		src:  clnTestSourceCyr,
		mask: CleanPrint,
		exp:  "\u009F\u00a0",
	},
	{
		key:  "graphic cyrillic",
		src:  clnTestSourceCyr,
		mask: CleanGraphic,
		exp:  "\u009F",
	},
	{
		key:  "default mask cyrillic",
		src:  clnTestSourceCyr,
		mask: DefaultCleanMask,
		exp:  "контроль  печатные  графические ï маркировка  пунктуация  пробел \u00a0 цифра  число  символ  буква Я",
	},
}

func TestCleaner(t *testing.T) {
	for _, stage := range clnStages {
		t.Run(stage.key, func(t *testing.T) {
			cln := NewUnicodeCleaner[string](stage.mask)
			r := cln.Clean(nil, stage.src)
			_, s := fastconv.AppendR2S(nil, r)
			if s != stage.exp {
				t.FailNow()
			}
		})
	}
}

func BenchmarkCleaner(b *testing.B) {
	for _, stage := range clnStages {
		b.Run(stage.key, func(b *testing.B) {
			b.ReportAllocs()
			cln := NewUnicodeCleaner[string](stage.mask)
			var buf []rune
			for i := 0; i < b.N; i++ {
				buf = cln.Clean(buf[:0], stage.src)
			}
			_ = buf
		})
	}
}
