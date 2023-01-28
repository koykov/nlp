package nlp

import (
	"unicode"

	"github.com/koykov/byteseq"
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

type UnicodeCleaner[T byteseq.Byteseq] struct {
	Mask uint32

	o bool
	m uint32
}

func NewUnicodeCleaner[T byteseq.Byteseq](mask uint32) *UnicodeCleaner[T] {
	return &UnicodeCleaner[T]{Mask: mask}
}

func (c *UnicodeCleaner[T]) Clean(x T) []rune {
	return c.AppendClean(nil, x)
}

func (c *UnicodeCleaner[T]) AppendClean(dst []rune, x T) []rune {
	s := byteseq.Q2S(x)
	if !c.o {
		c.m = c.Mask
		c.o = true
	}
	m := c.m
	for _, r := range s {
		if m > 0 {
			if m&CleanControl > 0 {
				if unicode.IsControl(r) {
					continue
				}
			}
			if m&CleanPrint > 0 {
				if unicode.IsPrint(r) {
					continue
				}
			}
			if m&CleanGraphic > 0 {
				if unicode.IsGraphic(r) {
					continue
				}
			}
			if m&CleanMark > 0 {
				if unicode.IsMark(r) {
					continue
				}
			}
			if m&CleanPunct > 0 {
				if unicode.IsPunct(r) {
					continue
				}
			}
			if m&CleanSpace > 0 {
				if unicode.IsSpace(r) {
					continue
				}
			}
			if m&CleanDigit > 0 {
				if unicode.IsDigit(r) {
					continue
				}
			}
			if m&CleanNumber > 0 {
				if unicode.IsNumber(r) {
					continue
				}
			}
			if m&CleanSymbol > 0 {
				if unicode.IsSymbol(r) {
					continue
				}
			}
			if m&CleanLetter > 0 {
				if unicode.IsLetter(r) {
					continue
				}
			}
		}
		dst = append(dst, r)
	}
	return dst
}
