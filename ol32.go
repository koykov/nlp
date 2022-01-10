package nlp

type OL32 uint32

func (m *OL32) Encode(lo, hi uint16) {
	*m = OL32(lo)<<16 | OL32(hi)
}

func (m OL32) Decode() (lo, hi uint16) {
	lo = uint16(m >> 16)
	hi = uint16(m & 0xffff)
	return
}

func (m *OL32) Reset() {
	*m = 0
}
