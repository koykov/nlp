package nlp

import (
	"encoding/binary"
	"io"
	"sync"
	"unicode"

	"github.com/koykov/bytealg"
	"github.com/koykov/byteseq"
	"github.com/koykov/fastconv"
)

const (
	ngmBlockSize = 4096
	ngmBufSize   = 16384
)

type NGModel[T byteseq.Byteseq] struct {
	v uint64

	o sync.Once
	u map[Unigram]struct{}
	b map[Bigram]struct{}
	t map[Trigram]struct{}
	q map[Quadrigram]struct{}
	f map[Fivegram]struct{}

	ul, bl, tl, ql, fl uint64

	buf  []byte
	bufR []rune
}

func (m *NGModel[T]) Parse(text T) *NGModel[T] {
	s := byteseq.Q2S(text)
	if len(text) == 0 {
		return m
	}
	m.o.Do(m.init)
	m.bufR = fastconv.AppendS2R(m.bufR[:0], s)
	l := len(m.bufR)
	_ = m.bufR[l-1]
	for i := 0; i < l; i++ {
		m.bufR[i] = unicode.ToLower(m.bufR[i])
		m.AddUnigram(Unigram(m.bufR[i]))
	}
	for i := 0; i < l-1; i++ {
		u0, u1 := Unigram(m.bufR[i]), Unigram(m.bufR[i+1])
		var b Bigram
		b = Bigram(u0) << 16
		b = b | Bigram(u1)
		m.AddBigram(b)
	}
	for i := 0; i < l-2; i++ {
		m.AddTrigram(Trigram{
			a: Unigram(m.bufR[i]),
			b: Unigram(m.bufR[i+1]),
			c: Unigram(m.bufR[i+2]),
		})
	}
	for i := 0; i < l-3; i++ {
		u0, u1, u2, u3 := Unigram(m.bufR[i]), Unigram(m.bufR[i+1]), Unigram(m.bufR[i+2]), Unigram(m.bufR[i+3])
		var q Quadrigram
		q = Quadrigram(u0) << 48
		q = q | Quadrigram(u1)<<32
		q = q | Quadrigram(u2)<<16
		q = q | Quadrigram(u3)
		m.AddQuadrigram(q)
	}
	for i := 0; i < l-4; i++ {
		m.AddFivegram(Fivegram{
			a: Unigram(m.bufR[i]),
			b: Unigram(m.bufR[i+1]),
			c: Unigram(m.bufR[i+2]),
			d: Unigram(m.bufR[i+3]),
			e: Unigram(m.bufR[i+4]),
		})
	}

	return m
}

func (m *NGModel[T]) AddUnigram(ng Unigram) *NGModel[T] {
	m.o.Do(m.init)
	if _, ok := m.u[ng]; ok {
		return m
	}
	m.u[ng] = struct{}{}
	return m
}

func (m *NGModel[T]) AddBigram(ng Bigram) *NGModel[T] {
	m.o.Do(m.init)
	if _, ok := m.b[ng]; ok {
		return m
	}
	m.b[ng] = struct{}{}
	return m
}

func (m *NGModel[T]) AddTrigram(ng Trigram) *NGModel[T] {
	m.o.Do(m.init)
	if _, ok := m.t[ng]; ok {
		return m
	}
	m.t[ng] = struct{}{}
	return m
}

func (m *NGModel[T]) AddQuadrigram(ng Quadrigram) *NGModel[T] {
	m.o.Do(m.init)
	if _, ok := m.q[ng]; ok {
		return m
	}
	m.q[ng] = struct{}{}
	return m
}

func (m *NGModel[T]) AddFivegram(ng Fivegram) *NGModel[T] {
	m.o.Do(m.init)
	if _, ok := m.f[ng]; ok {
		return m
	}
	m.f[ng] = struct{}{}
	return m
}

func (m *NGModel[T]) LoadFile(path string) error {
	_ = path
	return nil
}

func (m *NGModel[T]) Write(w io.Writer) (int, error) {
	w64 := func(dst []byte, v uint64) []byte {
		off := len(dst)
		dst = bytealg.GrowDelta(dst, 8)
		binary.LittleEndian.PutUint64(dst[off:], v)
		return dst
	}

	m.ul, m.bl, m.tl, m.ql, m.fl = uint64(len(m.u)), uint64(len(m.b)), uint64(len(m.t)), uint64(len(m.q)), uint64(len(m.f))

	m.buf = m.buf[:0]
	m.buf = w64(m.buf, m.v)
	m.buf = w64(m.buf, m.ul)
	m.buf = w64(m.buf, m.bl)
	m.buf = w64(m.buf, m.tl)
	m.buf = w64(m.buf, m.ql)
	m.buf = w64(m.buf, m.fl)

	_ = w
	return 0, nil
}

func (m *NGModel[T]) Flush() error {
	return nil
}

func (m *NGModel[T]) writeNG(w io.Writer, ng Unigram) (err error) {
	off := len(m.buf)
	m.buf = bytealg.GrowDelta(m.buf, 2)
	binary.LittleEndian.PutUint16(m.buf[off:], uint16(ng))
	if len(m.buf) > ngmBufSize {
		err = m.flushBuf(w)
	}
	return
}

func (m *NGModel[T]) flushBuf(w io.Writer) (err error) {
	p := m.buf
	if len(p) == 0 {
		return
	}
	for len(p) > ngmBlockSize {
		if _, err = w.Write(p[:ngmBlockSize]); err != nil {
			return
		}
		p = p[ngmBlockSize:]
	}
	if len(p) > 0 {
		if _, err = w.Write(p[:ngmBlockSize]); err != nil {
			return
		}
	}
	m.buf = m.buf[:0]
	return
}

func (m *NGModel[T]) init() {
	m.u = make(map[Unigram]struct{}, m.ul)
	m.b = make(map[Bigram]struct{}, m.bl)
	m.t = make(map[Trigram]struct{}, m.tl)
	m.q = make(map[Quadrigram]struct{}, m.ql)
	m.f = make(map[Fivegram]struct{}, m.fl)
}
