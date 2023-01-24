package nlp

import (
	"unicode"
)

const (
	CleanControl = 1 << iota
	CleanMark
	CleanPunct
	CleanSpace
	CleanDigit
	CleanNumber
	CleanSymbol
	CleanLetter
	CleanPrint
	CleanGraphic

	DefaultCleanMask = CleanControl | CleanMark | CleanSymbol | CleanNumber | CleanPunct
)

type Cleaner[T Byteseq] interface {
	Clean(dst []rune, x T) []rune
}

type UnicodeCleaner[T Byteseq] struct {
	mask uint32
	init bool
}

func NewUnicodeCleaner[T Byteseq](m uint32) UnicodeCleaner[T] {
	return UnicodeCleaner[T]{mask: m, init: true}
}

func (c UnicodeCleaner[T]) Clean(dst []rune, x T) []rune {
	s := q2s(x)
	mask := c.mask
	if !c.init && mask == 0 {
		mask = DefaultCleanMask
	}
	for _, r := range s {
		if mask > 0 {
			if mask&CleanControl > 0 {
				if unicode.IsControl(r) {
					continue
				}
			}
			if mask&CleanPrint > 0 {
				if unicode.IsPrint(r) {
					continue
				}
			}
			if mask&CleanGraphic > 0 {
				if unicode.IsGraphic(r) {
					continue
				}
			}
			if mask&CleanMark > 0 {
				if unicode.IsMark(r) {
					continue
				}
			}
			if mask&CleanPunct > 0 {
				if unicode.IsPunct(r) {
					continue
				}
			}
			if mask&CleanSpace > 0 {
				if unicode.IsSpace(r) {
					continue
				}
			}
			if mask&CleanDigit > 0 {
				if unicode.IsDigit(r) {
					continue
				}
			}
			if mask&CleanNumber > 0 {
				if unicode.IsNumber(r) {
					continue
				}
			}
			if mask&CleanSymbol > 0 {
				if unicode.IsSymbol(r) {
					continue
				}
			}
			if mask&CleanLetter > 0 {
				if unicode.IsLetter(r) {
					continue
				}
			}
		}
		dst = append(dst, r)
	}
	return dst
}
