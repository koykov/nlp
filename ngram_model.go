package nlp

import (
	"io"
	"sync"
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

func (m *NgramModel) Load(path string) error {
	_ = path
	return nil
}

func (m *NgramModel) Write(w io.Writer) (int, error) {
	_ = w
	return 0, nil
}

func (m *NgramModel) Flush() error {
	return nil
}

func (m *NgramModel) init() {
	m.u = make(map[Unigram]struct{}, m.ul)
	m.b = make(map[Bigram]struct{}, m.bl)
	m.t = make(map[Trigram]struct{}, m.tl)
	m.q = make(map[Quadrigram]struct{}, m.ql)
	m.f = make(map[Fivegram]struct{}, m.fl)
}
