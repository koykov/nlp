package nlp

import (
	"testing"

	"github.com/koykov/fastconv"
)

const clnTestSource = "control \u009f print ´ graphic ï mark \u035f punct { space \u00a0 digit 5 number ² symbol = letter ì"

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
}

func TestCleaner(t *testing.T) {
	for _, stage := range clnStages {
		t.Run(stage.key, func(t *testing.T) {
			cln := NewCleanerWithMask(stage.mask)
			r := cln.CleanString(nil, stage.src)
			_, s := fastconv.AppendR2S(nil, r)
			if s != stage.exp {
				println(s)
				t.FailNow()
			}
		})
	}
}
