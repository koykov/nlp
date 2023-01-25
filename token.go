package nlp

import (
	"reflect"
	"unsafe"

	"github.com/koykov/byteseq"
)

type Token struct {
	h   reflect.SliceHeader
	off int
}

func (t Token) Bytes() []byte {
	return *(*[]byte)(unsafe.Pointer(&t.h))
}

func (t Token) String() string {
	h := reflect.StringHeader{
		Data: t.h.Data,
		Len:  t.h.Len,
	}
	return *(*string)(unsafe.Pointer(&h))
}

func (t Token) Span() (lo, hi int) {
	return t.off, t.off + t.h.Len
}

type Tokens []Token

func (t Tokens) Each(fn func(i int, t Token)) {
	for i := 0; i < len(t); i++ {
		fn(i, t[i])
	}
}

func (t Tokens) Equal(e Tokens) bool {
	lt, le := len(t), len(e)
	if lt != le {
		return false
	}
	if lt == 0 {
		return true
	}
	for i := 0; i < lt; i++ {
		if t[i].String() != e[i].String() {
			return false
		}
	}
	return true
}

func (t *Tokens) Reset() {
	*t = (*t)[:0]
}

func ParseToken[T byteseq.Byteseq](s T, lo, hi int) Token {
	switch any(s).(type) {
	case string:
		return s2t(string(s), lo, hi)
	case []byte:
		return b2t([]byte(s), lo, hi)
	}
	return Token{}
}

func b2t(p []byte, lo, hi int) Token {
	if l := len(p); l == 0 || lo < 0 || lo >= hi || hi < 0 || hi > l || hi < lo {
		return Token{}
	}
	h := *(*reflect.SliceHeader)(unsafe.Pointer(&p))
	h.Data += uintptr(lo)
	h.Len = hi - lo
	h.Cap = h.Len
	return Token{
		h:   h,
		off: lo,
	}
}

func s2t(s string, lo, hi int) Token {
	if l := len(s); l == 0 || lo < 0 || lo >= hi || hi < 0 || hi > l || hi < lo {
		return Token{}
	}
	h := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	h.Data += uintptr(lo)
	h.Len = hi - lo
	return Token{
		h: reflect.SliceHeader{
			Data: h.Data,
			Len:  h.Len,
			Cap:  h.Len,
		},
		off: lo,
	}
}
