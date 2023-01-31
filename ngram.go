package nlp

import (
	"encoding/binary"

	"github.com/koykov/bytealg"
)

type Unigram uint16

type Bigram uint32

type Trigram struct {
	a, b, c Unigram
}

type Quadrigram uint64

type Fivegram struct {
	a, b, c, d, e Unigram
}

// group: Unigram

func NewUnigram(r rune) Unigram {
	return Unigram(r)
}

func (u Unigram) AppendTo(dst []byte) []byte {
	off := len(dst)
	dst = bytealg.GrowDelta(dst, 2)
	binary.LittleEndian.PutUint16(dst[off:], uint16(u))
	return dst
}

func (u Unigram) String() string {
	r := [1]rune{rune(u)}
	return string(r[:])
}

// endgroup: Unigram

// group: Bigram

func NewBigram(a, b rune) (n Bigram) {
	u0, u1 := Unigram(a), Unigram(b)
	n = Bigram(u0) << 16
	n = n | Bigram(u1)
	return
}

func (b Bigram) AppendTo(dst []byte) []byte {
	off := len(dst)
	dst = bytealg.GrowDelta(dst, 4)
	binary.LittleEndian.PutUint32(dst[off:], uint32(b))
	return dst
}

func (b Bigram) String() string {
	r := [2]rune{rune(uint16(b >> 16)), rune(uint16(b))}
	return string(r[:])
}

// endgroup: Bigram

// group: Trigram

func NewTrigram(a, b, c rune) Trigram {
	return Trigram{Unigram(a), Unigram(b), Unigram(c)}
}

func (t Trigram) AppendTo(dst []byte) []byte {
	off := len(dst)
	dst = bytealg.GrowDelta(dst, 6)
	binary.LittleEndian.PutUint16(dst[off:], uint16(t.a))
	binary.LittleEndian.PutUint16(dst[off+2:], uint16(t.b))
	binary.LittleEndian.PutUint16(dst[off+4:], uint16(t.c))
	return dst
}

func (t Trigram) String() string {
	r := [3]rune{rune(t.a), rune(t.b), rune(t.c)}
	return string(r[:])
}

// endgroup: Trigram

// group: Quadrigram

func NewQuadrigram(a, b, c, d rune) (n Quadrigram) {
	u0, u1, u2, u3 := Unigram(a), Unigram(b), Unigram(c), Unigram(d)
	n = Quadrigram(u0) << 48
	n = n | Quadrigram(u1)<<32
	n = n | Quadrigram(u2)<<16
	n = n | Quadrigram(u3)
	return
}

func (q Quadrigram) AppendTo(dst []byte) []byte {
	off := len(dst)
	dst = bytealg.GrowDelta(dst, 8)
	binary.LittleEndian.PutUint64(dst[off:], uint64(q))
	return dst
}

func (q Quadrigram) String() string {
	r := [4]rune{rune(uint16(q >> 48)), rune(uint16(q >> 32)), rune(uint16(q >> 16)), rune(uint16(q))}
	return string(r[:])
}

// endgroup: Quadrigram

// group: Fivegram

func NewFivegram(a, b, c, d, e rune) Fivegram {
	return Fivegram{Unigram(a), Unigram(b), Unigram(c), Unigram(d), Unigram(e)}
}

func (f Fivegram) AppendTo(dst []byte) []byte {
	off := len(dst)
	dst = bytealg.GrowDelta(dst, 10)
	binary.LittleEndian.PutUint16(dst[off:], uint16(f.a))
	binary.LittleEndian.PutUint16(dst[off+2:], uint16(f.b))
	binary.LittleEndian.PutUint16(dst[off+4:], uint16(f.c))
	binary.LittleEndian.PutUint16(dst[off+6:], uint16(f.d))
	binary.LittleEndian.PutUint16(dst[off+8:], uint16(f.e))
	return dst
}

func (f Fivegram) String() string {
	r := [5]rune{rune(f.a), rune(f.b), rune(f.c), rune(f.d), rune(f.e)}
	return string(r[:])
}

// endgroup: Fivegram
