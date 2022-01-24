package nlp

import "unicode"

type Script uint

func (s Script) Evaluate(p []byte) bool {
	// todo implement
	_ = p
	return false
}

func (s Script) EvaluateString(p string) bool {
	// todo implement
	_ = p
	return false
}

func (s Script) EvaluateRunes(p []rune) bool {
	// todo implement
	_ = p
	return false
}

func (s Script) EvaluateByte(b byte) bool {
	// todo implement
	_ = b
	return false
}

func (s Script) EvaluateRune(r rune) bool {
	t, ok := s.gets()
	if !ok {
		return false
	}
	return unicode.Is(t, r)
}

func (s Script) gets() (*unicode.RangeTable, bool) {
	if int(s) >= len(__st_buf) {
		return nil, false
	}
	return __st_buf[s], true
}

func (s Script) EvaluateRune2(r rune) bool {
	return __sre_buf[s](r)
}

// EvaluateRune1 is testing precompiled version of Han script todo remove me
func (s Script) EvaluateRune1(r rune) bool {
	// return __sre_buf[s].Evaluate(r)
	l16 := len(unicode.Han.R16)
	if r <= rune(0xfad9) {
		c := uint16(r)
		if l16 <= 18 || c <= unicode.MaxLatin1 {
			if c < 0x2e80 {
				return false
			}
			if c <= 0x2e99 {
				return true
			}

			if c < 0x2e9b {
				return false
			}
			if c <= 0x2ef3 {
				return true
			}

			if c < 0x2f00 {
				return false
			}
			if c <= 0x2fd5 {
				return true
			}

			if c < 0x3005 {
				return false
			}
			if c <= 0x3007 {
				return true
			}

			if c < 0x3021 {
				return false
			}
			if c <= 0x3029 {
				return true
			}

			if c < 0x3038 {
				return false
			}
			if c <= 0x303b {
				return true
			}

			if c < 0x3400 {
				return false
			}
			if c <= 0x4dbf {
				return true
			}

			if c < 0x4e00 {
				return false
			}
			if c <= 0x9ffc {
				return true
			}

			if c < 0xf900 {
				return false
			}
			if c <= 0xfa6d {
				return true
			}

			if c < 0xfa70 {
				return false
			}
			if c <= 0xfad9 {
				return true
			}
		}
	}
	return false
}
