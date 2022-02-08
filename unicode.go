package nlp

import "unicode"

func mustSkip(r rune) bool {
	return unicode.IsSymbol(r) || unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsDigit(r)
}
