package nlp

type Unigram uint16

type Bigram uint32

type Trigram struct {
	a, b, c Unigram
}

type Quadrigram uint64

type Fivegram struct {
	a, b, c, d, e Unigram
}

func NewUnigram(r rune) Unigram {
	return Unigram(r)
}

func NewBigram(a, b rune) (n Bigram) {
	u0, u1 := Unigram(a), Unigram(b)
	n = Bigram(u0) << 16
	n = n | Bigram(u1)
	return
}

func NewTrigram(a, b, c rune) Trigram {
	return Trigram{Unigram(a), Unigram(b), Unigram(c)}
}

func NewQuadrigram(a, b, c, d rune) (n Quadrigram) {
	u0, u1, u2, u3 := Unigram(a), Unigram(b), Unigram(c), Unigram(d)
	n = Quadrigram(u0) << 48
	n = n | Quadrigram(u1)<<32
	n = n | Quadrigram(u2)<<16
	n = n | Quadrigram(u3)
	return
}

func NewFivegram(a, b, c, d, e rune) Fivegram {
	return Fivegram{Unigram(a), Unigram(b), Unigram(c), Unigram(d), Unigram(e)}
}

func (u Unigram) String() string {
	r := [1]rune{rune(u)}
	return string(r[:])
}

func (b Bigram) String() string {
	r := [2]rune{rune(uint16(b >> 16)), rune(uint16(b))}
	return string(r[:])
}

func (t Trigram) String() string {
	r := [3]rune{rune(t.a), rune(t.b), rune(t.c)}
	return string(r[:])
}

func (q Quadrigram) String() string {
	r := [4]rune{rune(uint16(q >> 48)), rune(uint16(q >> 32)), rune(uint16(q >> 16)), rune(uint16(q))}
	return string(r[:])
}

func (f Fivegram) String() string {
	r := [5]rune{rune(f.a), rune(f.b), rune(f.c), rune(f.d), rune(f.e)}
	return string(r[:])
}
