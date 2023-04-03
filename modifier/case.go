package modifier

import (
	"unicode"

	"github.com/koykov/byteseq"
	"github.com/koykov/fastconv"
)

// group: Lower

type Lower[T byteseq.Byteseq] struct{}

func (m Lower[T]) Modify(x T) T {
	r := m.AppendModify(nil, x)
	_, s := fastconv.AppendR2S(nil, r)
	return byteseq.S2Q[T](s)
}

func (m Lower[T]) AppendModify(dst []rune, x T) []rune {
	dst = case_(dst, x, unicode.LowerCase)
	return dst
}

// endgroup: Lower

// group: Upper

type Upper[T byteseq.Byteseq] struct{}

func (m Upper[T]) Modify(x T) T {
	r := m.AppendModify(nil, x)
	_, s := fastconv.AppendR2S(nil, r)
	return byteseq.S2Q[T](s)
}

func (m Upper[T]) AppendModify(dst []rune, x T) []rune {
	dst = case_(dst, x, unicode.UpperCase)
	return dst
}

// endgroup: Upper

// group: Title

type Title[T byteseq.Byteseq] struct{}

func (m Title[T]) Modify(x T) T {
	r := m.AppendModify(nil, x)
	_, s := fastconv.AppendR2S(nil, r)
	return byteseq.S2Q[T](s)
}

func (m Title[T]) AppendModify(dst []rune, x T) []rune {
	dst = case_(dst, x, unicode.TitleCase)
	return dst
}

// endgroup: Lower

func case_[T byteseq.Byteseq](dst []rune, x T, case_ int) []rune {
	s := byteseq.Q2S(x)
	if len(s) == 0 {
		return dst
	}
	var sep bool
	for i, r := range s {
		switch case_ {
		case unicode.LowerCase:
			dst = append(dst, unicode.ToLower(r))
		case unicode.UpperCase:
			dst = append(dst, unicode.ToUpper(r))
		case unicode.TitleCase:
			if sep || i == 0 {
				dst = append(dst, unicode.ToTitle(r))
				sep = false
			} else {
				dst = append(dst, r)
			}
			sep = unicode.IsSpace(r) || (unicode.IsPunct(r) && r != '_')
		}
	}
	return dst
}
