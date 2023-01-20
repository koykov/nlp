package nlp

import (
	"unicode"

	"github.com/koykov/fastconv"
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

type CleanerInterface interface {
	CleanBytes(dst []rune, p []byte) []rune
	CleanString(dst []rune, s string) []rune
}

type Cleaner struct {
	m uint32
}

func NewCleaner() Cleaner {
	return NewCleanerWithMask(CleanControl | CleanMark | CleanSymbol | CleanNumber | CleanPunct)
}

func NewCleanerWithMask(m uint32) Cleaner {
	return Cleaner{m: m}
}

func (c Cleaner) CleanBytes(dst []rune, p []byte) []rune {
	return c.CleanString(dst, fastconv.B2S(p))
}

func (c Cleaner) CleanString(dst []rune, s string) []rune {
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
