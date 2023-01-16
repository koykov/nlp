package nlp

import "github.com/koykov/fastconv"

type Tokens struct {
	buf  []byte
	span []int
}

func (t *Tokens) SetSource(p []byte) {
	t.buf = append(t.buf, p...)
}

func (t *Tokens) SetSourceString(s string) {
	t.buf = append(t.buf, s...)
}

func (t *Tokens) AddSpan(lo, hi int) {
	if l := len(t.buf); lo < 0 || lo >= l || hi < 0 || hi >= l {
		return
	}
	t.span = append(t.span, lo, hi)
}

func (t Tokens) Each(fn func(i int, p []byte)) {
	var c int
	for i := 0; i < len(t.span); i += 2 {
		lo, hi := t.span[i], t.span[i+1]
		fn(c, t.buf[lo:hi])
		c++
	}
}

func (t Tokens) EachString(fn func(i int, p string)) {
	var c int
	for i := 0; i < len(t.span); i += 2 {
		lo, hi := t.span[i], t.span[i+1]
		fn(c, fastconv.B2S(t.buf[lo:hi]))
		c++
	}
}

func (t Tokens) EachSpan(fn func(i int, lo, hi int)) {
	var c int
	for i := 0; i < len(t.span); i += 2 {
		lo, hi := t.span[i], t.span[i+1]
		fn(c, lo, hi)
		c++
	}
}

func (t *Tokens) Reset() {
	t.buf = t.buf[:0]
	t.span = t.span[:0]
}
