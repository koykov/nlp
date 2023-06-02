package nlp

import (
	"encoding/binary"
	"io"
	"sort"
	"sync"
	"unicode"

	"github.com/koykov/bytealg"
	"github.com/koykov/byteseq"
	"github.com/koykov/fastconv"
)

const (
	ngmWordSep   = " \n\t"
	ngmBlockSize = 4096
	ngmBufSize   = 16384
)

type NGModel[T byteseq.Byteseq] struct {
	Version   uint64
	Tokenizer Tokenizer[T]

	o   sync.Once
	tkn Tokenizer[T]
	u   map[Unigram]struct{}
	b   map[Bigram]struct{}
	t   map[Trigram]struct{}
	q   map[Quadrigram]struct{}
	f   map[Fivegram]struct{}

	ul, bl, tl, ql, fl uint64

	buf  []byte
	bufR []rune
	bufT Tokens
}

func (m *NGModel[T]) Parse(text T) *NGModel[T] {
	if len(text) == 0 {
		return m
	}
	m.o.Do(m.init)
	m.bufT.Reset()
	m.tkn.AppendTokenize(m.bufT, text).Each(func(_ int, t Token) {
		m.parseWord(t.String())
	})

	return m
}

func (m *NGModel[T]) parseWord(s string) {
	m.bufR = fastconv.AppendS2R(m.bufR[:0], s)
	l := len(m.bufR)
	_ = m.bufR[l-1]
	for i := 0; i < l; i++ {
		m.bufR[i] = unicode.ToLower(m.bufR[i])
		m.AddUnigram(NewUnigram(m.bufR[i]))
	}
	for i := 0; i < l-1; i++ {
		m.AddBigram(NewBigram(m.bufR[i], m.bufR[i+1]))
	}
	for i := 0; i < l-2; i++ {
		m.AddTrigram(NewTrigram(m.bufR[i], m.bufR[i+1], m.bufR[i+2]))
	}
	for i := 0; i < l-3; i++ {
		m.AddQuadrigram(NewQuadrigram(m.bufR[i], m.bufR[i+1], m.bufR[i+2], m.bufR[i+3]))
	}
	for i := 0; i < l-4; i++ {
		m.AddFivegram(NewFivegram(m.bufR[i], m.bufR[i+1], m.bufR[i+2], m.bufR[i+3], m.bufR[i+4]))
	}
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

func (m *NGModel[T]) Stat() (int, int, int, int, int) {
	m.o.Do(m.init)
	return len(m.u), len(m.b), len(m.t), len(m.q), len(m.f)
}

func (m *NGModel[T]) LoadFile(path string) error {
	_ = path
	return nil
}

func (m *NGModel[T]) Read(r io.Reader) (n int, err error) {
	var n1 int
	m.buf = bytealg.GrowDelta(m.buf[:0], 48)
	if n1, err = io.ReadAtLeast(r, m.buf, 48); err != nil {
		return
	}
	n += n1
	header := m.buf[:48]
	ver := binary.LittleEndian.Uint64(header[:8])
	if ver != m.Version {
		err = ErrBadVersion
		return
	}
	m.ul = binary.LittleEndian.Uint64(header[8:16])
	m.bl = binary.LittleEndian.Uint64(header[16:24])
	m.tl = binary.LittleEndian.Uint64(header[24:32])
	m.ql = binary.LittleEndian.Uint64(header[32:40])
	m.fl = binary.LittleEndian.Uint64(header[40:48])

	if n1, err = m.readU(r); err != nil {
		return
	}
	n += n1
	if n1, err = m.readB(r); err != nil {
		return
	}
	n += n1
	if n1, err = m.readT(r); err != nil {
		return
	}
	n += n1
	if n1, err = m.readQ(r); err != nil {
		return
	}
	n += n1
	if n1, err = m.readF(r); err != nil {
		return
	}
	n += n1

	return
}

func (m *NGModel[T]) Write(w io.Writer) (n int, err error) {
	m.o.Do(m.init)
	w64 := func(dst []byte, v uint64) []byte {
		off := len(dst)
		dst = bytealg.GrowDelta(dst, 8)
		binary.LittleEndian.PutUint64(dst[off:], v)
		return dst
	}

	m.ul, m.bl, m.tl, m.ql, m.fl = uint64(len(m.u)), uint64(len(m.b)), uint64(len(m.t)), uint64(len(m.q)), uint64(len(m.f))

	m.buf = m.buf[:0]
	m.buf = w64(m.buf, m.Version)
	m.buf = w64(m.buf, m.ul)
	m.buf = w64(m.buf, m.bl)
	m.buf = w64(m.buf, m.tl)
	m.buf = w64(m.buf, m.ql)
	m.buf = w64(m.buf, m.fl)

	writeNG := func(w io.Writer, buf ngsort) (n int, err error) {
		sort.Sort(buf)
		var n1 int
		buf.Each(func(i int, ng appenderTo) {
			m.buf = ng.AppendTo(m.buf)
			if len(m.buf) > ngmBufSize {
				n1, err = m.flushBuf(w)
				n += n1
			}
		})
		return
	}

	var n1 int
	bufU := appendUnisort(nil, m.u)
	n1, err = writeNG(w, &bufU)
	n += n1

	bufB := appendBisort(nil, m.b)
	n1, err = writeNG(w, &bufB)
	n += n1

	bufT := appendTrisort(nil, m.t)
	n1, err = writeNG(w, &bufT)
	n += n1

	bufQ := appendQuadsort(nil, m.q)
	n1, err = writeNG(w, &bufQ)
	n += n1

	bufF := appendFivesort(nil, m.f)
	n1, err = writeNG(w, &bufF)
	n += n1

	n1, _ = m.flushBuf(w)
	n += n1

	return
}

func (m *NGModel[T]) flushBuf(w io.Writer) (n int, err error) {
	p := m.buf
	if len(p) == 0 {
		return
	}
	var n1 int
	for len(p) > ngmBlockSize {
		if n1, err = w.Write(p[:ngmBlockSize]); err != nil {
			return
		}
		n += n1
		p = p[ngmBlockSize:]
	}
	if len(p) > 0 {
		if n1, err = w.Write(p); err != nil {
			return
		}
		n += n1
	}
	m.buf = m.buf[:0]
	return
}

func (m *NGModel[T]) init() {
	if m.Tokenizer == nil {
		m.Tokenizer = StringTokenizer[T]{Separator: ngmWordSep}
	}
	m.tkn = m.Tokenizer

	m.u = make(map[Unigram]struct{}, m.ul)
	m.b = make(map[Bigram]struct{}, m.bl)
	m.t = make(map[Trigram]struct{}, m.tl)
	m.q = make(map[Quadrigram]struct{}, m.ql)
	m.f = make(map[Fivegram]struct{}, m.fl)
}

func (m *NGModel[T]) readU(r io.Reader) (n int, err error) {
	const c = 2
	var (
		off uint64
		n1  int
	)
	if m.ul*c > ngmBlockSize {
		for i := 0; i < int(m.ul/ngmBlockSize); i++ {
			m.buf = bytealg.GrowDelta(m.buf[:0], ngmBlockSize)
			if n1, err = io.ReadAtLeast(r, m.buf, ngmBlockSize); err != nil {
				return
			}
			n += n1
			off += ngmBlockSize
			for j := 0; j < len(m.buf); j += c {
				raw := binary.LittleEndian.Uint16(m.buf[j : j+1])
				m.AddUnigram(NewUnigram(rune(raw)))
			}
		}
	}
	m.buf = bytealg.GrowDelta(m.buf[:0], int(m.ul*c-off))
	if n1, err = io.ReadAtLeast(r, m.buf, int(m.ul*c-off)); err != nil {
		return
	}
	n += n1
	for j := 0; j < len(m.buf); j += c {
		raw := binary.LittleEndian.Uint16(m.buf[j : j+c])
		m.AddUnigram(NewUnigram(rune(raw)))
	}
	return
}

func (m *NGModel[T]) readB(r io.Reader) (n int, err error) {
	const c = 4
	var (
		off uint64
		n1  int
	)
	if m.bl*c > ngmBlockSize {
		for i := 0; i < int(m.ul/ngmBlockSize); i++ {
			m.buf = bytealg.GrowDelta(m.buf[:0], ngmBlockSize)
			if n1, err = io.ReadAtLeast(r, m.buf, ngmBlockSize); err != nil {
				return
			}
			n += n1
			off += ngmBlockSize
			for j := 0; j < len(m.buf); j += c {
				raw := binary.LittleEndian.Uint32(m.buf[j : j+c])
				m.AddBigram(Bigram(raw))
			}
		}
	}
	m.buf = bytealg.GrowDelta(m.buf[:0], int(m.bl*c-off))
	if n1, err = io.ReadAtLeast(r, m.buf, int(m.bl*c-off)); err != nil {
		return
	}
	n += n1
	for j := 0; j < len(m.buf); j += c {
		raw := binary.LittleEndian.Uint32(m.buf[j : j+c])
		m.AddBigram(Bigram(raw))
	}
	return
}

func (m *NGModel[T]) readT(r io.Reader) (n int, err error) {
	const c = 6
	var (
		off uint64
		n1  int
	)
	if m.tl*c > ngmBlockSize {
		size := uint64(ngmBlockSize - ngmBufSize%c)
		for i := 0; i < int(m.ul/size); i++ {
			m.buf = bytealg.GrowDelta(m.buf[:0], int(size))
			if n1, err = io.ReadAtLeast(r, m.buf, int(size)); err != nil {
				return
			}
			n += n1
			off += size
			for j := 0; j < len(m.buf); j += c {
				raw0 := binary.LittleEndian.Uint16(m.buf[j : j+2])
				raw1 := binary.LittleEndian.Uint16(m.buf[j+2 : j+4])
				raw2 := binary.LittleEndian.Uint16(m.buf[j+4 : j+6])
				m.AddTrigram(NewTrigram(rune(raw0), rune(raw1), rune(raw2)))
			}
		}
	}
	m.buf = bytealg.GrowDelta(m.buf[:0], int(m.tl*c-off))
	if n1, err = io.ReadAtLeast(r, m.buf, int(m.tl*c-off)); err != nil {
		return
	}
	n += n1
	for j := 0; j < len(m.buf); j += c {
		raw0 := binary.LittleEndian.Uint16(m.buf[j : j+2])
		raw1 := binary.LittleEndian.Uint16(m.buf[j+2 : j+4])
		raw2 := binary.LittleEndian.Uint16(m.buf[j+4 : j+6])
		m.AddTrigram(NewTrigram(rune(raw0), rune(raw1), rune(raw2)))
	}
	return
}

func (m *NGModel[T]) readQ(r io.Reader) (n int, err error) {
	const c = 8
	var (
		off uint64
		n1  int
	)
	if m.tl*c > ngmBlockSize {
		size := uint64(ngmBlockSize - ngmBufSize%c)
		for i := 0; i < int(m.ul/size); i++ {
			m.buf = bytealg.GrowDelta(m.buf[:0], int(size))
			if n1, err = io.ReadAtLeast(r, m.buf, int(size)); err != nil {
				return
			}
			n += n1
			off += size
			for j := 0; j < len(m.buf); j += c {
				raw := binary.LittleEndian.Uint64(m.buf[j : j+c])
				m.AddQuadrigram(Quadrigram(raw))
			}
		}
	}
	m.buf = bytealg.GrowDelta(m.buf[:0], int(m.ql*c-off))
	if n1, err = io.ReadAtLeast(r, m.buf, int(m.ql*c-off)); err != nil {
		return
	}
	n += n1
	for j := 0; j < len(m.buf); j += c {
		raw := binary.LittleEndian.Uint64(m.buf[j : j+c])
		m.AddQuadrigram(Quadrigram(raw))
	}
	return
}

func (m *NGModel[T]) readF(r io.Reader) (n int, err error) {
	const c = 10
	var (
		off uint64
		n1  int
	)
	if m.tl*c > ngmBlockSize {
		size := uint64(ngmBlockSize - ngmBufSize%c)
		for i := 0; i < int(m.ul/size); i++ {
			m.buf = bytealg.GrowDelta(m.buf[:0], int(size))
			if n1, err = io.ReadAtLeast(r, m.buf, int(size)); err != nil {
				return
			}
			n += n1
			off += size
			for j := 0; j < len(m.buf); j += c {
				raw0 := binary.LittleEndian.Uint16(m.buf[j : j+2])
				raw1 := binary.LittleEndian.Uint16(m.buf[j+2 : j+4])
				raw2 := binary.LittleEndian.Uint16(m.buf[j+4 : j+6])
				raw3 := binary.LittleEndian.Uint16(m.buf[j+6 : j+8])
				raw4 := binary.LittleEndian.Uint16(m.buf[j+8 : j+10])
				m.AddFivegram(NewFivegram(rune(raw0), rune(raw1), rune(raw2), rune(raw3), rune(raw4)))
			}
		}
	}
	m.buf = bytealg.GrowDelta(m.buf[:0], int(m.fl*c-off))
	if n1, err = io.ReadAtLeast(r, m.buf, int(m.fl*c-off)); err != nil {
		return
	}
	n += n1
	for j := 0; j < len(m.buf); j += c {
		raw0 := binary.LittleEndian.Uint16(m.buf[j : j+2])
		raw1 := binary.LittleEndian.Uint16(m.buf[j+2 : j+4])
		raw2 := binary.LittleEndian.Uint16(m.buf[j+4 : j+6])
		raw3 := binary.LittleEndian.Uint16(m.buf[j+6 : j+8])
		raw4 := binary.LittleEndian.Uint16(m.buf[j+8 : j+10])
		m.AddFivegram(NewFivegram(rune(raw0), rune(raw1), rune(raw2), rune(raw3), rune(raw4)))
	}
	return
}
