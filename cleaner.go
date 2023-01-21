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
)

type CleanerInterface[T Byteseq] interface {
	Clean(dst []rune, x T) []rune
}

type Cleaner[T Byteseq] struct {
	m uint32
}

func NewCleaner[T Byteseq]() Cleaner[T] {
	cln := NewCleanerWithMask[T](CleanControl | CleanMark | CleanSymbol | CleanNumber | CleanPunct)
	return cln
}

func NewCleanerWithMask[T Byteseq](m uint32) Cleaner[T] {
	return Cleaner[T]{m: m}
}

func (c Cleaner[T]) Clean(dst []rune, x T) []rune {
	s := q2s(x)
	for _, r := range s {
		if c.m > 0 {
			if c.m&CleanControl > 0 {
				if unicode.IsControl(r) {
					continue
				}
			}
			if c.m&CleanPrint > 0 {
				if unicode.IsPrint(r) {
					continue
				}
			}
			if c.m&CleanGraphic > 0 {
				if unicode.IsGraphic(r) {
					continue
				}
			}
			if c.m&CleanMark > 0 {
				if unicode.IsMark(r) {
					continue
				}
			}
			if c.m&CleanPunct > 0 {
				if unicode.IsPunct(r) {
					continue
				}
			}
			if c.m&CleanSpace > 0 {
				if unicode.IsSpace(r) {
					continue
				}
			}
			if c.m&CleanDigit > 0 {
				if unicode.IsDigit(r) {
					continue
				}
			}
			if c.m&CleanNumber > 0 {
				if unicode.IsNumber(r) {
					continue
				}
			}
			if c.m&CleanSymbol > 0 {
				if unicode.IsSymbol(r) {
					continue
				}
			}
			if c.m&CleanLetter > 0 {
				if unicode.IsLetter(r) {
					continue
				}
			}
		}
		dst = append(dst, r)
	}
	return dst
}
