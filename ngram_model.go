package nlp

import (
	"encoding/binary"
	"io"
	"sync"

	"github.com/koykov/bytealg"
)

const (
	ngmBlockSize = 4096
	ngmBufSize   = 16384
)

type NgramModel struct {
	v uint64

	o sync.Once
	u map[Unigram]struct{}
	b map[Bigram]struct{}
	t map[Trigram]struct{}
	q map[Quadrigram]struct{}
	f map[Fivegram]struct{}

	ul, bl, tl, ql, fl uint64

	buf []byte
}

func (m *NgramModel) AddUnigram(ng Unigram) *NgramModel {
	m.o.Do(m.init)
	if _, ok := m.u[ng]; ok {
		return m
	}
	m.u[ng] = struct{}{}
	return m
}

func (m *NgramModel) AddBigram(ng Bigram) *NgramModel {
	m.o.Do(m.init)
	if _, ok := m.b[ng]; ok {
		return m
	}
	m.b[ng] = struct{}{}
	return m
}

func (m *NgramModel) AddTrigram(ng Trigram) *NgramModel {
	m.o.Do(m.init)
	if _, ok := m.t[ng]; ok {
		return m
	}
	m.t[ng] = struct{}{}
	return m
}

func (m *NgramModel) AddQuadrigram(ng Quadrigram) *NgramModel {
	m.o.Do(m.init)
	if _, ok := m.q[ng]; ok {
		return m
	}
	m.q[ng] = struct{}{}
	return m
}

func (m *NgramModel) AddFivegram(ng Fivegram) *NgramModel {
	m.o.Do(m.init)
	if _, ok := m.f[ng]; ok {
		return m
	}
	m.f[ng] = struct{}{}
	return m
}

func (m *NgramModel) LoadFile(path string) error {
	_ = path
	return nil
}

func (m *NgramModel) Write(w io.Writer) (int, error) {
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

func (m *NgramModel) Flush() error {
	return nil
}

func (m *NgramModel) writeNG(w io.Writer, ng Unigram) (err error) {
	off := len(m.buf)
	m.buf = bytealg.GrowDelta(m.buf, 2)
	binary.LittleEndian.PutUint16(m.buf[off:], uint16(ng))
	if len(m.buf) > ngmBufSize {
		err = m.flushBuf(w)
	}
	return
}

func (m *NgramModel) flushBuf(w io.Writer) (err error) {
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

func (m *NgramModel) init() {
	m.u = make(map[Unigram]struct{}, m.ul)
	m.b = make(map[Bigram]struct{}, m.bl)
	m.t = make(map[Trigram]struct{}, m.tl)
	m.q = make(map[Quadrigram]struct{}, m.ql)
	m.f = make(map[Fivegram]struct{}, m.fl)
}
