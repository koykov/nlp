package nlp

type OL32 uint32

func (m *OL32) encode(lo, hi uint16) {
	*m = OL32(lo)<<16 | OL32(hi)
}

func (m *OL32) decode() (lo, hi uint16) {
	lo = uint16(*m >> 16)
	hi = uint16(*m & 0xffff)
	return
}
