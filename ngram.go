package nlp

type Unigram uint16

func (u Unigram) String() string {
	r := [1]rune{rune(u)}
	return string(r[:])
}

type Bigram uint32

func (b Bigram) String() string {
	r := [2]rune{rune(uint16(b >> 16)), rune(uint16(b))}
	return string(r[:])
}

type Trigram struct {
	a, b, c Unigram
}

func (t Trigram) String() string {
	r := [3]rune{rune(t.a), rune(t.b), rune(t.c)}
	return string(r[:])
}

type Quadrigram uint64

func (q Quadrigram) String() string {
	r := [4]rune{rune(uint16(q >> 48)), rune(uint16(q >> 32)), rune(uint16(q >> 16)), rune(uint16(q))}
	return string(r[:])
}

type Fivegram struct {
	a, b, c, d, e Unigram
}

func (f Fivegram) String() string {
	r := [5]rune{rune(f.a), rune(f.b), rune(f.c), rune(f.d), rune(f.e)}
	return string(r[:])
}
